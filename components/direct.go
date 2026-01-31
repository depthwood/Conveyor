package components

import (
	"errors"
	"github.com/conveyor/conveyor/core"
)

// DirectComponent 是一个直接组件，用于在同一个应用中传递消息
type DirectComponent struct {
	consumers map[string][]core.Consumer
}

// NewDirectComponent 创建一个新的DirectComponent实例
func NewDirectComponent() *DirectComponent {
	return &DirectComponent{
		consumers: make(map[string][]core.Consumer),
	}
}

// CreateProducer 创建一个DirectProducer实例
func (dc *DirectComponent) CreateProducer(uri string) (core.Producer, error) {
	return &DirectProducer{
		component: dc,
		target:    uri,
	}, nil
}

// CreateConsumer 创建一个DirectConsumer实例
func (dc *DirectComponent) CreateConsumer(uri string, processor core.Processor) (core.Consumer, error) {
	consumer := &DirectConsumer{
		processor: processor,
	}

	// 将消费者添加到列表中
	dc.consumers[uri] = append(dc.consumers[uri], consumer)

	return consumer, nil
}

// DirectProducer 是DirectComponent的生产者
type DirectProducer struct {
	component *DirectComponent
	target    string
}

// Send 发送消息到目标URI
func (dp *DirectProducer) Send(exchange *core.Exchange) error {
	consumers, ok := dp.component.consumers[dp.target]
	if !ok || len(consumers) == 0 {
		return errors.New("no consumers for uri: " + dp.target)
	}

	// 将消息发送给所有消费者
	for _, consumer := range consumers {
		if dc, ok := consumer.(*DirectConsumer); ok {
			if err := dc.processor.Process(exchange); err != nil {
				return err
			}
		}
	}

	return nil
}

// DirectConsumer 是DirectComponent的消费者
type DirectConsumer struct {
	processor core.Processor
}

// Start 启动消费者
func (dc *DirectConsumer) Start() error {
	// DirectConsumer不需要启动，因为它是直接调用的
	return nil
}

// Stop 停止消费者
func (dc *DirectConsumer) Stop() error {
	// DirectConsumer不需要停止
	return nil
}
