package core

// Route 表示一条完整的路由
type Route struct {
	ID        string
	FromURI   string
	ToURI     string
	Processor Processor
}

// NewRoute 创建一个新的路由实例
func NewRoute(id, fromURI, toURI string, processor Processor) *Route {
	return &Route{
		ID:        id,
		FromURI:   fromURI,
		ToURI:     toURI,
		Processor: processor,
	}
}

// RouteBuilder 用于构建路由的构建器
type RouteBuilder interface {
	// Configure 配置路由
	Configure()
}
