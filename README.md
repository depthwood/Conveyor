# Conveyor

Conveyor是一个基于Apache Camel启发的开源集成框架，使用Golang开发，旨在提供简洁、高效的方式来集成各种系统和服务。

## 特性

- 基于Golang 1.18+开发，充分利用泛型和并发特性
- 提供流畅的DSL API，用于定义路由和消息处理规则
- 支持多种组件，用于不同的传输协议和消息模型
- 与Gin、GoFrame等主流Golang框架无缝集成
- 轻量级设计，易于嵌入到任何Golang应用中
- 强大的消息处理能力，支持各种企业集成模式

## 技术栈

### 核心技术
- **Golang 1.18+**：使用最新的Go语言特性，包括泛型、并发模型等
- **Go Modules**：依赖管理系统

### 外部依赖
- **github.com/gin-gonic/gin v1.11.0**：轻量级Web框架，用于HTTP请求处理
- **github.com/gogf/gf/v2 v2.9.8**：GoFrame框架，提供全功能的Web开发支持

### 项目结构
- **core**：核心包，包含Message、Exchange、Processor等核心概念
- **components**：组件包，实现各种传输协议的生产者和消费者
- **dsl**：领域特定语言，提供流畅的路由定义API
- **adapter**：适配器包，实现与外部框架的集成
- **examples**：示例代码，展示各种使用场景

## 核心概念

- **Message**：在路由中传递的消息，包含消息体和消息头
- **Exchange**：消息交换上下文，包含输入消息和输出消息
- **Processor**：消息处理器，负责处理消息交换
- **Route**：定义消息的流转路径，从源到目标
- **Component**：组件，用于创建生产者和消费者，支持不同的传输协议
- **ConveyorContext**：核心上下文，管理路由和组件

## 快速开始

### 安装

```bash
go get github.com/conveyor/conveyor
```

### 基本使用

```go
package main

import (
	"fmt"
	"github.com/conveyor/conveyor"
	"github.com/conveyor/conveyor/components"
	"github.com/conveyor/conveyor/core"
)

func main() {
	// 创建ConveyorContext
	ctx := conveyor.NewConveyorContext()
	
	// 添加Direct组件
	directComponent := components.NewDirectComponent()
	ctx.AddComponent("direct", directComponent)
	
	// 创建路由构建器
	rb := conveyor.NewRouteBuilder(ctx)
	
	// 定义路由
	rb.From("direct:test").ProcessFunc(func(exchange *core.Exchange) error {
			// 处理消息
			body := exchange.InMessage.Body.([]byte)
			fmt.Printf("Received message: %s\n", body)
			
			// 创建响应
			response := fmt.Sprintf("Processed: %s", body)
			exchange.SetOutMessage(conveyor.NewMessage([]byte(response)))
			
			return nil
		}).End()
	
	// 启动上下文
	if err := ctx.Start(); err != nil {
		panic(err)
	}
	defer ctx.Stop()
	
	// 使用生产者发送消息
	producer, _ := directComponent.CreateProducer("direct:test")
	exchange := conveyor.NewExchange(core.InOut, conveyor.NewMessage([]byte("Hello Conveyor!")))
	producer.Send(exchange)
	
	// 输出响应
	if exchange.HasOutMessage() {
		fmt.Printf("Response: %s\n", exchange.OutMessage.Body)
	}
}
```

### 与Gin框架集成

```go
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/conveyor/conveyor"
	"github.com/conveyor/conveyor/adapter"
	"github.com/conveyor/conveyor/components"
)

func main() {
	// 创建ConveyorContext
	ctx := conveyor.NewConveyorContext()
	
	// 添加Direct组件
	ctx.AddComponent("direct", components.NewDirectComponent())
	
	// 定义路由
	rb := conveyor.NewRouteBuilder(ctx)
	rb.From("direct:api").ProcessFunc(func(exchange *core.Exchange) error {
		// 处理API请求
		return nil
	}).End()
	
	// 启动上下文
	ctx.Start()
	defer ctx.Stop()
	
	// 创建Gin引擎
	r := gin.Default()
	
	// 创建Gin适配器
	ginAdapter := adapter.NewGinAdapter(ctx)
	
	// 设置路由
	r.POST("/api", ginAdapter.HandleHTTPRequest("direct:api"))
	
	// 启动Gin服务器
	r.Run(":8080")
}
```

### 与GoFrame框架集成

```go
package main

import (
	"github.com/conveyor/conveyor"
	"github.com/conveyor/conveyor/adapter"
	"github.com/conveyor/conveyor/components"
	"github.com/gogf/gf/v2/frame/g"
)

func main() {
	// 创建ConveyorContext
	ctx := conveyor.NewConveyorContext()
	
	// 添加Direct组件
	ctx.AddComponent("direct", components.NewDirectComponent())
	
	// 定义路由
	rb := conveyor.NewRouteBuilder(ctx)
	rb.From("direct:api").ProcessFunc(func(exchange *core.Exchange) error {
		// 处理API请求
		return nil
	}).End()
	
	// 启动上下文
	ctx.Start()
	defer ctx.Stop()
	
	// 创建GoFrame适配器
	gfAdapter := adapter.NewGoFrameAdapter(ctx)
	
	// 设置路由
	s := g.Server()
	s.BindHandler("POST:/api", gfAdapter.HandleHTTPRequest("direct:api"))
	
	// 启动HTTP服务器
	s.Run()
}
```

## 组件

### Direct组件

Direct组件用于在同一个应用中直接传递消息，适用于测试和简单的内部消息传递。

```go
// 添加Direct组件
ctx.AddComponent("direct", components.NewDirectComponent())

// 定义路由
rb.From("direct:source").To("direct:target").End()
```

## 未来规划

- 支持更多组件，如HTTP、JMS、Kafka等
- 实现更多企业集成模式
- 提供更强大的DSL功能
- 支持监控和管理
- 提供更多示例和文档

## 贡献

欢迎提交Issue和Pull Request，共同完善Conveyor项目。

## 许可证

Conveyor使用Apache License 2.0许可证。
