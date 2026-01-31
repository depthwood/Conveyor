package core

// Processor 定义消息处理接口
type Processor interface {
	// Process 处理消息交换
	Process(exchange *Exchange) error
}

// ProcessorFunc 是一个函数类型，实现了Processor接口
type ProcessorFunc func(exchange *Exchange) error

// Process 实现Processor接口的Process方法
func (pf ProcessorFunc) Process(exchange *Exchange) error {
	return pf(exchange)
}

// Pipeline 是一个处理器链，按顺序执行多个处理器
type Pipeline struct {
	Processors []Processor
}

// NewPipeline 创建一个新的处理器链
func NewPipeline(processors ...Processor) *Pipeline {
	return &Pipeline{
		Processors: processors,
	}
}

// Process 实现Processor接口，按顺序执行所有处理器
func (p *Pipeline) Process(exchange *Exchange) error {
	for _, processor := range p.Processors {
		if err := processor.Process(exchange); err != nil {
			return err
		}
	}
	return nil
}
