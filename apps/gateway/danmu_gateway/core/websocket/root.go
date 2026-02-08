package websocket

import (
	"LiveDanmu/apps/gateway/danmu_gateway/core/models"
	"context"
	"sync"
	"sync/atomic"
	"time"

	"github.com/hertz-contrib/websocket"
)

// ws设置
const (
	pingDelay    = 30 * time.Second
	pongWait     = 60 * time.Second
	writeWait    = 10 * time.Second
	maxFailNum   = 2
	maxSendRetry = 3
)

// Pool 连接池结构体
type Pool struct {
	Shutdown   context.CancelFunc
	clients    map[string]*Client       // 活跃连接映射
	register   chan *Client             // 注册通道
	unregister chan *Client             // 注销通道
	broadcast  chan models.WebsocketMsg // 广播消息通道
	mu         sync.RWMutex             // 读写锁保护 clients 映射
	count      int32                    // 原子计数器
}

// Client 每个连接的结构体
type Client struct {
	ID       string
	Conn     *websocket.Conn
	Pool     *Pool                  // 所属连接池
	Metadata map[string]interface{} // 元数据

	pongReceived chan struct{}
	missedPongs  atomic.Int32
	isAlive      atomic.Bool
	cancelHeart  context.CancelFunc

	mu sync.RWMutex
}
