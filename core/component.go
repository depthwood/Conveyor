package core

// Component 定义组件接口，用于创建生产者和消费者
type Component interface {
	// CreateProducer 创建一个生产者
	CreateProducer(uri string) (Producer, error)

	// CreateConsumer 创建一个消费者
	CreateConsumer(uri string, processor Processor) (Consumer, error)
}

// Producer 定义消息生产者接口
type Producer interface {
	// Send 发送消息
	Send(exchange *Exchange) error
}

// Consumer 定义消息消费者接口
type Consumer interface {
	// Start 启动消费者
	Start() error

	// Stop 停止消费者
	Stop() error
}
