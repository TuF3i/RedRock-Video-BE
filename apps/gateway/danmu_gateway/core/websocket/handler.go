package websocket

import (
	"context"
	"fmt"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/hertz-contrib/websocket"
)

var upgrader = websocket.HertzUpgrader{
	CheckOrigin: func(ctx *app.RequestContext) bool {
		return true
	},
}

// WebSocket连接入口 - 修正版
func HandleWebSocket(ctx context.Context, c *app.RequestContext) {
	roomID := c.Query("rvid")
	if roomID == "" {
		c.String(400, "rvid required")
		return
	}

	var rid int64
	fmt.Sscanf(roomID, "%d", &rid)

	// 新版本API：使用HertzHandler回调
	err := upgrader.Upgrade(c, func(conn *websocket.Conn) {
		client := &Client{
			ID:      fmt.Sprintf("%d_%d", rid, time.Now().UnixNano()),
			RoomID:  rid,
			Conn:    conn,
			Send:    make(chan []byte, 256),
			Manager: GetManager(),
		}

		client.Manager.register(client)

		// 启动写入协程，然后阻塞在读取循环
		go client.writePump()
		client.readPump() // 阻塞直到连接断开

		// 连接断开后的清理
		client.Manager.unregister(client)
	})

	if err != nil {
		hlog.Errorf("WebSocket upgrade failed: %v", err)
	}
}
