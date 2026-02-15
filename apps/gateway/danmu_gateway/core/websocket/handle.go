package websocket

import (
	"LiveDanmu/apps/gateway/danmu_gateway/core/models"
	"context"
	"errors"
	"sync"

	"github.com/hertz-contrib/websocket"
)

// NewPoolGroup 创建连接池组
func NewPoolGroup() *PoolGroup {
	return &PoolGroup{Pools: make(map[int64]*Pool), mu: sync.RWMutex{}}
}

// NewPool 创建新连接池
func (r *PoolGroup) NewPool(ctx context.Context, rvid int64) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.Pools[rvid]; ok {
		return
	}
	pool := NewPool()
	// 启动连接池调度器
	pool.Run(ctx)
	// 将连接池加入连接池组
	r.Pools[rvid] = pool
}

// CancelPool 关闭一个连接池
func (r *PoolGroup) CancelPool(rvid int64) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.Pools[rvid]; !ok {
		return
	}
	// 关闭连接池调度器
	r.Pools[rvid].Shutdown()
	// 删除连接池
	delete(r.Pools, rvid)
}

// AddConnToGroup 向一个Poor里添加一个连接
func (r *PoolGroup) AddConnToGroup(rvid int64, conn *websocket.Conn) error {
	r.mu.Lock()
	if _, ok := r.Pools[rvid]; !ok {
		r.mu.Unlock()
		return errors.New("pool not exists")
	}
	r.mu.Unlock()

	err := r.Pools[rvid].RegisterNewClient(conn)
	if err != nil {
		return err
	}

	return nil
}

// BoardCastMsg 广播消息
func (r *PoolGroup) BoardCastMsg(rvid int64, msg models.WebsocketMsg) {
	defer func() {
		if r := recover(); r != nil {
			return
		}
	}()

	r.mu.RLock()
	if _, ok := r.Pools[rvid]; !ok {
		r.mu.RUnlock()
		return
	}
	r.mu.RUnlock()

	r.Pools[rvid].broadcast <- msg
}

func (r *PoolGroup) IfPoolExist(rvid int64) bool {
	_, ok := r.Pools[rvid]
	return ok
}
