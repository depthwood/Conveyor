package conveyor

import (
	"github.com/conveyor/conveyor/core"
	"github.com/conveyor/conveyor/dsl"
)

// ConveyorContext 是Conveyor的核心上下文
type ConveyorContext = core.ConveyorContext

// Message 表示在路由中传递的消息
type Message = core.Message

// Exchange 表示消息交换上下文
type Exchange = core.Exchange

// Processor 定义消息处理接口
type Processor = core.Processor

// ProcessorFunc 是一个函数类型，实现了Processor接口
type ProcessorFunc = core.ProcessorFunc

// NewConveyorContext 创建一个新的ConveyorContext实例
func NewConveyorContext() *ConveyorContext {
	return core.NewConveyorContext()
}

// NewMessage 创建一个新的消息实例
func NewMessage(body interface{}) *Message {
	return core.NewMessage(body)
}

// NewExchange 创建一个新的交换实例
func NewExchange(pattern core.ExchangePattern, inMessage *Message) *Exchange {
	return core.NewExchange(pattern, inMessage)
}

// NewRouteBuilder 创建一个新的路由构建器
func NewRouteBuilder(ctx *ConveyorContext) *dsl.RouteBuilder {
	return dsl.NewRouteBuilder(ctx)
}
