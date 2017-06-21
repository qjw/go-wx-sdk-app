package gateway

import (
	"encoding/json"
	"log"
	"github.com/qjw/go-wx-sdk/mp"
	"github.com/qjw/go-wx-sdk/utils"
	"github.com/qjw/kelly"
)

func InitializeApiRoutes(grouter kelly.Router,context *mp.Context) {
	grouter.GET("/weixin",func(c *kelly.Context) {
		server := mp.NewServer(c.Request(),c,nil, context)
		server.Ping()
	})


	grouter.POST("/weixin",func(c *kelly.Context) {
		server := mp.NewServer(c.Request(),c,func(msg *mp.MixMessage) utils.Reply {
			//回复消息：演示回复用户发送的消息
			if msg.MsgType == "text"{
				return utils.NewText(msg.Content)
			} else if msg.MsgType == "image"{
				return utils.NewImage(msg.MediaID)
			} else if msg.MsgType == "voice"{
				return utils.NewVoice(msg.MediaID)
			} else if msg.MsgType == "video"{
				return utils.NewVideo(msg.MediaID,"hello","world")
			} else {
				data,_ := json.MarshalIndent(msg,""," ")
				log.Print(string(data))
				return nil
			}
		},context)
		server.Serve()
	})
}

