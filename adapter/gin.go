package adapter

import (
	"github.com/conveyor/conveyor/core"
	"github.com/gin-gonic/gin"
)

// GinAdapter 是Conveyor与Gin框架的集成适配器
type GinAdapter struct {
	ctx *core.ConveyorContext
}

// NewGinAdapter 创建一个新的GinAdapter实例
func NewGinAdapter(ctx *core.ConveyorContext) *GinAdapter {
	return &GinAdapter{
		ctx: ctx,
	}
}

// HandleHTTPRequest 将Gin请求转换为Conveyor消息并处理
func (ga *GinAdapter) HandleHTTPRequest(uri string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 创建消息
		body, _ := c.GetRawData()
		msg := core.NewMessage(body)

		// 设置HTTP相关的消息头
		for key, values := range c.Request.Header {
			if len(values) > 0 {
				msg.SetHeader(key, values[0])
			}
		}

		msg.SetHeader("HTTP_METHOD", c.Request.Method)
		msg.SetHeader("HTTP_PATH", c.Request.URL.Path)
		msg.SetHeader("HTTP_QUERY", c.Request.URL.RawQuery)

		// 创建交换
		exchange := core.NewExchange(core.InOut, msg)

		// 发送到指定URI
		component, ok := ga.ctx.GetComponent("direct")
		if !ok {
			c.JSON(500, gin.H{"error": "direct component not found"})
			return
		}

		producer, err := component.CreateProducer(uri)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		// 处理消息
		if err := producer.Send(exchange); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		// 返回响应
		if exchange.HasOutMessage() {
			c.Data(200, "application/json", exchange.OutMessage.Body.([]byte))
		} else {
			c.Status(204)
		}
	}
}
