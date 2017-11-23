package gateway

import (
	"encoding/json"
	"github.com/qjw/go-wx-sdk/small"
	"github.com/qjw/go-wx-sdk/utils"
	"github.com/qjw/kelly"
	"log"
)

func InitializeSmApiRoutes(grouter kelly.Router, context *small.Context, api *small.SmallApi) {
	grouter.GET("/small", func(c *kelly.Context) {
		server := small.NewServer(c.Request(), c, nil, context)
		server.Ping()
	})

	grouter.POST("/small", func(c *kelly.Context) {
		server := small.NewServer(c.Request(), c, func(msg *small.MixMessage) utils.Reply {
			// 小程序不支持被动回复消息
			if msg.MsgType == "text" {
				log.Print(msg.Content)
				api.SendKfMessage(msg.FromUserName, msg.Content)
				return utils.NewTransfer()
			} else if msg.MsgType == "image" {
				log.Print(msg.PicURL)
				return utils.NewTransfer()
			} else {
				data, _ := json.MarshalIndent(msg, "", " ")
				log.Print(string(data))
				return nil
			}
		}, context)
		server.Serve()
	})
}
