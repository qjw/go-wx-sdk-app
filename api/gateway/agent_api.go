package gateway

import (
	"github.com/qjw/go-wx-sdk/corp"
	"github.com/qjw/go-wx-sdk/utils"
	"github.com/qjw/kelly"
)

func InitializeAgentApiRoutes(grouter kelly.Router,context *corp.Context,
	config *corp.AgentConfig, kfConfig *corp.KfConfig) {
	grouter.GET("/weixin",func(c *kelly.Context) {
		server := corp.NewServer(c.Request(),c,nil, context,config)
		server.Ping()
	})


	grouter.POST("/weixin",func(c *kelly.Context) {
		server := corp.NewServer(c.Request(),c,func(msg *corp.MixMessage) utils.Reply {
			//回复消息：演示回复用户发送的消息
			if msg.MsgType == "text"{
				return utils.NewText(msg.Content)
			} else if msg.MsgType == "image"{
				return utils.NewImage(msg.MediaID)
			} else if msg.MsgType == "voice"{
				return utils.NewVoice(msg.MediaID)
			} else if msg.MsgType == "video"{
				return utils.NewVideo(msg.MediaID,"hello","world")
			} else{
				return nil
			}
		},context,config)
		server.Serve()
	})

	grouter.GET("/kf",func(c *kelly.Context) {
		server := corp.NewKfServer(c.Request(),c,nil, context,kfConfig)
		server.Ping()
	})


	grouter.POST("/kf",func(c *kelly.Context) {
		server := corp.NewKfServer(c.Request(),c,func(msg *corp.KfMixMessage) (){

		},context,kfConfig)
		server.Serve()
	})
}

