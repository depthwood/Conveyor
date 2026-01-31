package core

import (
	"testing"
)

func TestMessage(t *testing.T) {
	// 创建消息
	msg := NewMessage("test body")

	// 设置和获取消息头
	msg.SetHeader("test-key", "test-value")
	value, ok := msg.GetHeader("test-key")
	if !ok || value != "test-value" {
		t.Errorf("Expected header 'test-key' to be 'test-value', got %v", value)
	}

	// 检查消息体
	if msg.Body != "test body" {
		t.Errorf("Expected body to be 'test body', got %v", msg.Body)
	}
}

func TestExchange(t *testing.T) {
	// 创建消息
	msg := NewMessage("test body")

	// 创建交换
	exchange := NewExchange(InOut, msg)

	// 设置和获取属性
	exchange.SetProperty("test-prop", "test-prop-value")
	value, ok := exchange.GetProperty("test-prop")
	if !ok || value != "test-prop-value" {
		t.Errorf("Expected property 'test-prop' to be 'test-prop-value', got %v", value)
	}

	// 检查输出消息
	if exchange.HasOutMessage() {
		t.Errorf("Expected no out message initially")
	}

	// 设置输出消息
	outMsg := NewMessage("out body")
	exchange.SetOutMessage(outMsg)

	if !exchange.HasOutMessage() {
		t.Errorf("Expected out message after setting")
	}
}

func TestProcessorFunc(t *testing.T) {
	// 创建消息和交换
	msg := NewMessage("test body")
	exchange := NewExchange(InOut, msg)

	// 创建处理器函数
	processor := ProcessorFunc(func(exchange *Exchange) error {
		// 简单处理：将消息体转为大写
		exchange.InMessage.Body = "TEST BODY"
		return nil
	})

	// 处理消息
	if err := processor.Process(exchange); err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// 检查处理结果
	if exchange.InMessage.Body != "TEST BODY" {
		t.Errorf("Expected body to be 'TEST BODY', got %v", exchange.InMessage.Body)
	}
}

func TestPipeline(t *testing.T) {
	// 创建消息和交换
	msg := NewMessage("test")
	exchange := NewExchange(InOut, msg)

	// 创建处理器链
	pipeline := NewPipeline(
		ProcessorFunc(func(exchange *Exchange) error {
			exchange.InMessage.Body = exchange.InMessage.Body.(string) + "-step1"
			return nil
		}),
		ProcessorFunc(func(exchange *Exchange) error {
			exchange.InMessage.Body = exchange.InMessage.Body.(string) + "-step2"
			return nil
		}),
	)

	// 处理消息
	if err := pipeline.Process(exchange); err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// 检查处理结果
	expected := "test-step1-step2"
	if exchange.InMessage.Body != expected {
		t.Errorf("Expected body to be '%s', got %v", expected, exchange.InMessage.Body)
	}
}
