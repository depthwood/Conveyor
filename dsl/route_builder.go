package dsl

import (
	"github.com/conveyor/conveyor/core"
)

// RouteBuilder 是一个流畅的路由构建器
type RouteBuilder struct {
	ctx     *core.ConveyorContext
	current *routeDefinition
}

// routeDefinition 定义了一条路由的构建过程
type routeDefinition struct {
	id           string
	fromURI      string
	toURI        string
	processor    core.Processor
	routeBuilder *RouteBuilder
}

// NewRouteBuilder 创建一个新的路由构建器
func NewRouteBuilder(ctx *core.ConveyorContext) *RouteBuilder {
	return &RouteBuilder{
		ctx: ctx,
	}
}

// From 开始定义一条路由，指定从哪里接收消息
func (rb *RouteBuilder) From(uri string) *routeDefinition {
	rd := &routeDefinition{
		fromURI:      uri,
		routeBuilder: rb,
	}
	rb.current = rd
	return rd
}

// To 指定路由的目标URI
func (rd *routeDefinition) To(uri string) *routeDefinition {
	rd.toURI = uri
	return rd
}

// Process 指定路由的处理器
func (rd *routeDefinition) Process(processor core.Processor) *routeDefinition {
	rd.processor = processor
	return rd
}

// ProcessFunc 指定路由的处理器函数
func (rd *routeDefinition) ProcessFunc(fn func(exchange *core.Exchange) error) *routeDefinition {
	rd.processor = core.ProcessorFunc(fn)
	return rd
}

// ID 指定路由的ID
func (rd *routeDefinition) ID(id string) *routeDefinition {
	rd.id = id
	return rd
}

// End 结束路由定义并添加到上下文
func (rd *routeDefinition) End() {
	route := core.NewRoute(rd.id, rd.fromURI, rd.toURI, rd.processor)
	rd.routeBuilder.ctx.AddRoute(route)
}
