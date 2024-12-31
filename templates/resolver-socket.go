package templates

var ResolverSockets = `package gen

import (
	"sync"
)

type SocketManager struct {
	mu      sync.RWMutex
	clients map[string]chan interface{} // 用户ID -> 消息通道
}

func NewSocketManager() *SocketManager {
	return &SocketManager{
		clients: make(map[string]chan interface{}),
	}
}

// Register 注册用户的连接
func (sm *SocketManager) Register(userID string, ch chan interface{}) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.clients[userID] = ch
}

// Unregister 移除用户的连接
func (sm *SocketManager) Unregister(userID string) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	if ch, ok := sm.clients[userID]; ok {
		close(ch)
		delete(sm.clients, userID)
	}
}

// Broadcast 推送消息给所有用户
func (sm *SocketManager) Broadcast(message interface{}) {
	sm.mu.RLock()
	defer sm.mu.RUnlock()
	for _, ch := range sm.clients {
		select {
		case ch <- message:
		default: // 避免阻塞
		}
	}
}

// SendToUser 推送消息给指定用户
func (sm *SocketManager) SendToUser(userID string, message interface{}) {
	sm.mu.RLock()
	defer sm.mu.RUnlock()
	if ch, ok := sm.clients[userID]; ok {
		select {
		case ch <- message:
		default: // 避免阻塞
		}
	}
}
`
