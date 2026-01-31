package main

import (
	"fmt"
	"github.com/conveyor/conveyor"
	"github.com/conveyor/conveyor/adapter"
	"github.com/conveyor/conveyor/components"
	"github.com/conveyor/conveyor/core"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
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
	rb.From("direct:goframe-test").ProcessFunc(func(exchange *core.Exchange) error {
		// 处理消息
		body := exchange.InMessage.Body.([]byte)
		fmt.Printf("GoFrame received message: %s\n", body)

		// 创建响应
		response := fmt.Sprintf("GoFrame processed: %s", body)
		exchange.SetOutMessage(conveyor.NewMessage([]byte(response)))

		return nil
	}).End()

	// 启动上下文
	if err := ctx.Start(); err != nil {
		panic(err)
	}
	defer ctx.Stop()

	// 创建GoFrame适配器
	gfAdapter := adapter.NewGoFrameAdapter(ctx)

	// 获取GoFrame的HTTP服务器
	s := g.Server()

	// 设置路由
	s.BindHandler("POST:/api/goframe-test", gfAdapter.HandleHTTPRequest("direct:goframe-test"))

	// 启动HTTP服务器
	s.Run()
}
