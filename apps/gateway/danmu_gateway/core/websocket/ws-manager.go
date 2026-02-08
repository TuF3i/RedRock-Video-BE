package websocket

import (
	"LiveDanmu/apps/gateway/danmu_gateway/core/models"
	"context"
	"sync/atomic"
	"time"

	"github.com/hertz-contrib/websocket"
)

func NewPool() *Pool {
	return &Pool{
		clients:    make(map[string]*Client),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		broadcast:  make(chan models.WebsocketMsg, 256),
	}
}

// Run 启动ws池调度
func (p *Pool) Run(ctx context.Context) {
	// 带取消的context
	ctx, cancel := context.WithCancel(ctx)
	p.Shutdown = cancel
	for {
		select {
		case <-ctx.Done():
			p.mu.Lock()
			for _, c := range p.clients {
				p.unregister <- c
			}
			close(p.register)
			close(p.unregister)
			close(p.broadcast)
			p.mu.Unlock()
			return
		// 注册连接
		case client := <-p.register:
			p.mu.Lock()
			// 注册一个连接
			p.clients[client.ID] = client
			//启动心跳
			client.startHeartbeat()
			// 原子计时器加一
			atomic.AddInt32(&p.count, 1)
			p.mu.Unlock()
		// 注销连接
		case client := <-p.unregister:
			p.mu.Lock()
			if _, ok := p.clients[client.ID]; ok {
				client.StopHeartbeat()
				delete(p.clients, client.ID)
				atomic.AddInt32(&p.count, -1)
			}
			p.mu.Unlock()
		// 广播纪元
		case message := <-p.broadcast:
			// 读锁完全复制客户端切片
			p.mu.RLock()
			clients := make([]*Client, 0, len(p.clients))
			for _, c := range p.clients {
				clients = append(clients, c)
			}
			p.mu.RUnlock()

			// 遍历广播
			for _, client := range clients {
				if !client.isAlive.Load() {
					continue
				}
				go func(c *Client) {
					// 发送失败则注销
					for i := 0; i < maxSendRetry; i++ {
						if c.sendMessage(message) {
							return
						}
						// 指数退避,ai教的
						time.Sleep(time.Duration(10*(1<<i)) * time.Millisecond)
					}
					c.Pool.unregister <- c
				}(client)
			}
		}
	}
}

// 启动心跳
func (c *Client) startHeartbeat() {
	ctx, cancel := context.WithCancel(context.Background())
	c.cancelHeart = cancel
	c.isAlive.Store(true)

	// 初始化pong响应通道
	c.pongReceived = make(chan struct{}, 1)
	// 设置pong处理器
	c.Conn.SetPongHandler(func(string) error {
		// 重置未响应计数
		c.missedPongs.Store(0)
		select {
		case c.pongReceived <- struct{}{}:
		default:
		}
		return nil
	})

	go c.heartbeatChecker(ctx)
}

// 心跳检测器
func (c *Client) heartbeatChecker(ctx context.Context) {
	ticker := time.NewTicker(pingDelay)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			if !c.checkAliveSync() {
				c.Pool.unregister <- c
				return
			}
		}
	}
}

func (c *Client) checkAliveSync() bool {
	// 设置写超时
	_ = c.Conn.SetWriteDeadline(time.Now().Add(writeWait))

	// 发送 Ping 消息
	if err := c.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
		c.isAlive.Store(false)
		return false
	}
	// 等pong
	select {
	case <-c.pongReceived:
		return true
	case <-time.After(pongWait):
		failCount := c.missedPongs.Add(1)
		if int(failCount) >= maxFailNum {
			c.isAlive.Store(false)
			return false
		}

		return true
	}
}

func (c *Client) StopHeartbeat() {
	if c.cancelHeart != nil {
		c.cancelHeart()
	}
	c.isAlive.Store(false)
}

// 发消息函数
func (c *Client) sendMessage(message models.WebsocketMsg) bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	if !c.isAlive.Load() {
		return false
	}

	if err := c.Conn.SetWriteDeadline(time.Now().Add(writeWait)); err != nil {
		c.isAlive.Store(false)
		return false
	}

	if err := c.Conn.WriteJSON(message); err != nil {
		c.isAlive.Store(false)
		return false
	}

	return true
}
