package websocket

import (
	"encoding/json"
	"sync"
	"time"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/hertz-contrib/websocket"
)

// 消息类型常量
const (
	MsgTypeDanmaku   = "danmaku"   // 弹幕消息
	MsgTypeHeartbeat = "heartbeat" // 心跳
	MsgTypePong      = "pong"      // 心跳响应
)

// 消息结构
type Message struct {
	Type string          `json:"type"`
	Data json.RawMessage `json:"data"`
}

// 客户端连接
type Client struct {
	ID      string
	RoomID  int64
	Conn    *websocket.Conn
	Send    chan []byte
	Manager *WebSocketManager
}

// 直播间
type Room struct {
	ID      int64
	Clients map[string]*Client
	mu      sync.RWMutex
}

// WebSocket管理器
type WebSocketManager struct {
	rooms map[int64]*Room
	mu    sync.RWMutex

	// 配置
	HeartbeatInterval time.Duration
	WriteTimeout      time.Duration
}

// 全局实例
var defaultManager *WebSocketManager

// 初始化管理器
func InitManager() *WebSocketManager {
	defaultManager = &WebSocketManager{
		rooms:             make(map[int64]*Room),
		HeartbeatInterval: 30 * time.Second,
		WriteTimeout:      10 * time.Second,
	}
	return defaultManager
}

// 获取管理器实例
func GetManager() *WebSocketManager {
	if defaultManager == nil {
		return InitManager()
	}
	return defaultManager
}

// 获取或创建直播间
func (m *WebSocketManager) getOrCreateRoom(roomID int64) *Room {
	m.mu.Lock()
	defer m.mu.Unlock()

	if room, ok := m.rooms[roomID]; ok {
		return room
	}

	room := &Room{
		ID:      roomID,
		Clients: make(map[string]*Client),
	}
	m.rooms[roomID] = room
	return room
}

// 注册客户端
func (m *WebSocketManager) register(client *Client) {
	room := m.getOrCreateRoom(client.RoomID)
	room.mu.Lock()
	room.Clients[client.ID] = client
	room.mu.Unlock()
	hlog.Infof("Client %s joined room %d", client.ID, client.RoomID)
}

// 注销客户端
func (m *WebSocketManager) unregister(client *Client) {
	room := m.getOrCreateRoom(client.RoomID)
	room.mu.Lock()
	if _, ok := room.Clients[client.ID]; ok {
		delete(room.Clients, client.ID)
		close(client.Send)
	}
	room.mu.Unlock()

	// 如果房间空了，清理房间
	room.mu.RLock()
	clientCount := len(room.Clients)
	room.mu.RUnlock()

	if clientCount == 0 {
		m.mu.Lock()
		delete(m.rooms, client.RoomID)
		m.mu.Unlock()
	}

	hlog.Infof("Client %s left room %d", client.ID, client.RoomID)
}

// 发送弹幕（核心暴露函数）
func SendDanmaku(roomID int64, data []byte) error {
	return GetManager().sendToRoom(roomID, data)
}

// 向直播间广播
func (m *WebSocketManager) sendToRoom(roomID int64, data []byte) error {
	m.mu.RLock()
	room, ok := m.rooms[roomID]
	m.mu.RUnlock()

	if !ok {
		return nil // 房间不存在，静默处理
	}

	// 包装成消息格式
	msg := Message{
		Type: MsgTypeDanmaku,
		Data: data,
	}
	msgBytes, _ := json.Marshal(msg)

	room.mu.RLock()
	defer room.mu.RUnlock()

	for _, client := range room.Clients {
		select {
		case client.Send <- msgBytes:
		default:
			// 发送缓冲满，关闭连接
			close(client.Send)
		}
	}
	return nil
}

// 客户端读取循环
func (c *Client) readPump() {
	defer c.Conn.Close()

	c.Conn.SetReadDeadline(time.Now().Add(60 * time.Second))
	c.Conn.SetPongHandler(func(string) error {
		c.Conn.SetReadDeadline(time.Now().Add(60 * time.Second))
		return nil
	})

	for {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				hlog.Errorf("WebSocket error: %v", err)
			}
			break
		}

		// 处理心跳
		var msg Message
		if err := json.Unmarshal(message, &msg); err == nil && msg.Type == MsgTypeHeartbeat {
			c.Conn.SetReadDeadline(time.Now().Add(60 * time.Second))
		}
	}
}

// 客户端写入循环（发送消息+心跳）
func (c *Client) writePump() {
	ticker := time.NewTicker(GetManager().HeartbeatInterval)
	defer func() {
		ticker.Stop()
		c.Conn.Close()
	}()

	for {
		select {
		case message, ok := <-c.Send:
			c.Conn.SetWriteDeadline(time.Now().Add(GetManager().WriteTimeout))
			if !ok {
				c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			c.Conn.WriteMessage(websocket.TextMessage, message)

		case <-ticker.C:
			// 发送心跳
			c.Conn.SetWriteDeadline(time.Now().Add(GetManager().WriteTimeout))
			heartbeat, _ := json.Marshal(Message{Type: MsgTypeHeartbeat})
			if err := c.Conn.WriteMessage(websocket.TextMessage, heartbeat); err != nil {
				return
			}
		}
	}
}
