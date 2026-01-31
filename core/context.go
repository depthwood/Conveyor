package core

// ConveyorContext 是Conveyor的核心上下文，管理路由和组件
type ConveyorContext struct {
	Routes       []*Route
	ComponentMap map[string]Component
}

// NewConveyorContext 创建一个新的ConveyorContext实例
func NewConveyorContext() *ConveyorContext {
	return &ConveyorContext{
		Routes:       make([]*Route, 0),
		ComponentMap: make(map[string]Component),
	}
}

// AddRoute 添加一条路由
func (ctx *ConveyorContext) AddRoute(route *Route) {
	ctx.Routes = append(ctx.Routes, route)
}

// GetRoutes 获取所有路由
func (ctx *ConveyorContext) GetRoutes() []*Route {
	return ctx.Routes
}

// AddComponent 添加一个组件
func (ctx *ConveyorContext) AddComponent(name string, component Component) {
	ctx.ComponentMap[name] = component
}

// GetComponent 获取一个组件
func (ctx *ConveyorContext) GetComponent(name string) (Component, bool) {
	component, ok := ctx.ComponentMap[name]
	return component, ok
}

// Start 启动上下文
func (ctx *ConveyorContext) Start() error {
	// 启动所有路由和组件
	return nil
}

// Stop 停止上下文
func (ctx *ConveyorContext) Stop() error {
	// 停止所有路由和组件
	return nil
}
