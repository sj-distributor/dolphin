package templates

var ResolverSockets = `package gen

import (
	"fmt"
	"sync"
)

type SocketManager struct {
	mu             sync.RWMutex
	clients        map[string]chan interface{}    // 用户ID -> 消息通道
	companyClients map[string]map[string]struct{} // 公司ID -> 用户ID集合
}

func NewSocketManager() *SocketManager {
	return &SocketManager{
		clients:        make(map[string]chan interface{}),
		companyClients: make(map[string]map[string]struct{}),
	}
}

// Register 注册用户的连接，并关联到公司
func (sm *SocketManager) Register(userID, companyID string, ch chan interface{}) {
	fmt.Println("Register", userID)

	sm.mu.Lock()
	defer sm.mu.Unlock()

	// 注册用户连接
	sm.clients[userID] = ch

	// 将用户加入公司
	if sm.companyClients[companyID] == nil {
		sm.companyClients[companyID] = make(map[string]struct{})
	}
	sm.companyClients[companyID][userID] = struct{}{}
}

// Unregister 移除用户的连接，并从公司中移除
func (sm *SocketManager) Unregister(userID string) {
	fmt.Println("Unregister", userID)
	sm.mu.Lock()
	defer sm.mu.Unlock()

	// 移除用户连接
	if ch, ok := sm.clients[userID]; ok {
		close(ch)
		delete(sm.clients, userID)
	}

	// 从所有公司中移除用户
	for companyID, users := range sm.companyClients {
		if _, ok := users[userID]; ok {
			delete(users, userID)
			if len(users) == 0 {
				delete(sm.companyClients, companyID) // 如果公司下无用户，移除公司记录
			}
			break
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

// SendToCompany 推送消息给指定公司下的所有用户
func (sm *SocketManager) SendToCompany(companyID string, message interface{}) {
	sm.mu.RLock()
	defer sm.mu.RUnlock()
	if users, ok := sm.companyClients[companyID]; ok {
		for userID := range users {
			if ch, exists := sm.clients[userID]; exists {
				select {
				case ch <- message:
				default: // 避免阻塞
				}
			}
		}
	}
}

// SendToUserInCompany 推送消息给指定公司下的某个用户
func (sm *SocketManager) SendToUserInCompany(companyID, userID string, message interface{}) {
	sm.mu.RLock()
	defer sm.mu.RUnlock()

	// 检查公司是否存在
	if users, ok := sm.companyClients[companyID]; ok {
		// 检查用户是否属于公司
		if _, userExists := users[userID]; userExists {
			if ch, exists := sm.clients[userID]; exists {
				select {
				case ch <- message:
				default: // 避免阻塞
				}
			}
		}
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
`
