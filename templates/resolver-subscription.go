package templates

var ResolverSubscriptions = `package gen

import (
	"context"
	"fmt"
	"time"

	"github.com/gofrs/uuid"
)

var Socket = NewSocketManager()

type GeneratedSubscriptionResolver struct{ *GeneratedResolver }

func (r *GeneratedSubscriptionResolver) WebSocket(ctx context.Context) (<-chan interface{}, error) {
	return r.Handlers.WebSocket(ctx, r.GeneratedResolver)
}

func WebSocketHandler(ctx context.Context, r *GeneratedResolver) (<-chan interface{}, error) {
	ch := make(chan interface{})

	userID := uuid.Must(uuid.NewV4()).String()

	// 注册用户连接
	Socket.Register(userID, "*", ch)

	go func() {
		defer Socket.Unregister(userID) // 连接断开时移除
		for {
			select {
			case <-ctx.Done():
				return
			default:
				// 如果有需要可以在这里处理用户的心跳检测等逻辑
				time.Sleep(1 * time.Second)
			}
		}
	}()
	return ch, nil
}
`
