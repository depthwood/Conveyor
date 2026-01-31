package main

import (
	"fmt"
	"github.com/conveyor/conveyor"
	"github.com/conveyor/conveyor/adapter"
	"github.com/conveyor/conveyor/components"
	"github.com/conveyor/conveyor/core"
	"github.com/gin-gonic/gin"
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

	// 创建Gin引擎
	r := gin.Default()

	// 创建Gin适配器
	ginAdapter := adapter.NewGinAdapter(ctx)

	// 设置路由
	r.POST("/api/test", ginAdapter.HandleHTTPRequest("direct:test"))

	// 启动Gin服务器
	r.Run(":8080")
}
