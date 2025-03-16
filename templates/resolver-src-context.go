package templates

var ResolverSrcContext = `package src
import (
	"context"
	"net/http"

	"{{.Config.Package}}/config"
)

// 结构体
type ContextStruct struct {
	ctx context.Context
}

// 设置 key-value 到 context
func (c *ContextStruct) SetValue(key any, value any) {
	c.ctx = context.WithValue(c.ctx, key, value)
}

// 获取 context 中的值
func (c *ContextStruct) GetValue(key any) string {
	if value, ok := c.ctx.Value(key).(string); ok {
		return value
	}
	return ""
}

// 获取 Header
func (c *ContextStruct) GetHeader(key string) string {
	if header, ok := c.ctx.Value(config.KeyHeader).(*http.Header); ok {
		return header.Get(key)
	}
	return ""
}

// 设置 Header
func (c *ContextStruct) SetHeader(header http.Header) {
	c.ctx = context.WithValue(c.ctx, config.KeyHeader, &header)
}

// 获取 Authorization
func (c *ContextStruct) GetAuthorization() string {
	return c.GetValue(config.KeyAuthorization)
}

// 设置 Authorization
func (c *ContextStruct) SetAuthorization(value string) {
	if len(value) > 0 {
		c.SetValue(config.KeyAuthorization, value)
	}
}

// 获取 SecretKey
func (c *ContextStruct) GetSecretKey() string {
	return c.GetValue(config.KeySecretKey)
}

// 设置 SecretKey
func (c *ContextStruct) SetSecretKey(value string) {
	if len(value) > 0 {
		c.SetValue(config.KeySecretKey, value)
	}
}

// Context
func (c *ContextStruct) Context() context.Context {
	return c.ctx
}

// Context
// 创建 ContextStruct 实例
func NewContext(ctx context.Context) *ContextStruct {
	return &ContextStruct{ctx: ctx}
}
`
