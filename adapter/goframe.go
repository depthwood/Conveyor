package adapter

import (
	"github.com/conveyor/conveyor/core"
	"github.com/gogf/gf/v2/net/ghttp"
)

// GoFrameAdapter 是Conveyor与GoFrame框架的集成适配器
type GoFrameAdapter struct {
	ctx *core.ConveyorContext
}

// NewGoFrameAdapter 创建一个新的GoFrameAdapter实例
func NewGoFrameAdapter(ctx *core.ConveyorContext) *GoFrameAdapter {
	return &GoFrameAdapter{
		ctx: ctx,
	}
}

// HandleHTTPRequest 将GoFrame请求转换为Conveyor消息并处理
func (gfa *GoFrameAdapter) HandleHTTPRequest(uri string) ghttp.HandlerFunc {
	return func(r *ghttp.Request) {
		// 创建消息
		body := r.GetBody()
		msg := core.NewMessage(body)

		// 设置HTTP相关的消息头
		for key, values := range r.Header {
			if len(values) > 0 {
				msg.SetHeader(key, values[0])
			}
		}

		msg.SetHeader("HTTP_METHOD", r.Method)
		msg.SetHeader("HTTP_PATH", r.URL.Path)
		msg.SetHeader("HTTP_QUERY", r.URL.RawQuery)

		// 创建交换
		exchange := core.NewExchange(core.InOut, msg)

		// 发送到指定URI
		component, ok := gfa.ctx.GetComponent("direct")
		if !ok {
			r.Response.WriteStatus(500)
			r.Response.WriteJson(map[string]interface{}{
				"error": "direct component not found",
			})
			return
		}

		producer, err := component.CreateProducer(uri)
		if err != nil {
			r.Response.WriteStatus(500)
			r.Response.WriteJson(map[string]interface{}{
				"error": err.Error(),
			})
			return
		}

		// 处理消息
		if err := producer.Send(exchange); err != nil {
			r.Response.WriteStatus(500)
			r.Response.WriteJson(map[string]interface{}{
				"error": err.Error(),
			})
			return
		}

		// 返回响应
		if exchange.HasOutMessage() {
			r.Response.WriteStatus(200)
			r.Response.Header().Set("Content-Type", "application/json")
			r.Response.Write(exchange.OutMessage.Body.([]byte))
		} else {
			r.Response.WriteStatus(204)
		}
	}
}
