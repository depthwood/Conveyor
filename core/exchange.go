package core

// ExchangePattern 定义消息交换模式
type ExchangePattern int

const (
	// InOnly 单向消息模式，只有输入消息
	InOnly ExchangePattern = iota
	// InOut 双向消息模式，有输入和输出消息
	InOut
)

// Exchange 表示消息交换上下文
type Exchange struct {
	Pattern    ExchangePattern
	InMessage  *Message
	OutMessage *Message
	Properties map[string]interface{}
}

// NewExchange 创建一个新的交换实例
func NewExchange(pattern ExchangePattern, inMessage *Message) *Exchange {
	return &Exchange{
		Pattern:    pattern,
		InMessage:  inMessage,
		OutMessage: nil,
		Properties: make(map[string]interface{}),
	}
}

// SetProperty 设置交换属性
func (e *Exchange) SetProperty(key string, value interface{}) {
	e.Properties[key] = value
}

// GetProperty 获取交换属性
func (e *Exchange) GetProperty(key string) (interface{}, bool) {
	value, ok := e.Properties[key]
	return value, ok
}

// HasOutMessage 检查是否有输出消息
func (e *Exchange) HasOutMessage() bool {
	return e.OutMessage != nil
}

// SetOutMessage 设置输出消息
func (e *Exchange) SetOutMessage(msg *Message) {
	e.OutMessage = msg
}
