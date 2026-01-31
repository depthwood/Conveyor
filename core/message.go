package core

// Message 表示在路由中传递的消息
type Message struct {
	Headers map[string]interface{}
	Body    interface{}
}

// NewMessage 创建一个新的消息实例
func NewMessage(body interface{}) *Message {
	return &Message{
		Headers: make(map[string]interface{}),
		Body:    body,
	}
}

// SetHeader 设置消息头
func (m *Message) SetHeader(key string, value interface{}) {
	m.Headers[key] = value
}

// GetHeader 获取消息头
func (m *Message) GetHeader(key string) (interface{}, bool) {
	value, ok := m.Headers[key]
	return value, ok
}
