package mp

import (
	"github.com/qjw/go-wx-sdk/small"
	"github.com/qjw/kelly"
	"log"
	"net/http"
)

type UserCode struct {
	Code string `json:"code"`
}

func InitializeSmallApiRoutes(grouter kelly.Router, context *small.Context, api *small.SmallApi) {
	urouter := grouter.Group("/sm")

	urouter.GET("/jscode2session",
		kelly.BindMiddleware(func() interface{} { return &UserCode{} }),
		func(c *kelly.Context) {
			param := c.GetBindParameter().(*UserCode)
			resp, err := api.Jscode2Session(param.Code)
			if err == nil {
				log.Printf("%s %s %s", resp.OpenID, resp.SessionKey, resp.UnionID)
			}else{
				log.Printf("%s",err.Error())
			}
			end_process(c, resp, err)
		})
}

func end_process(c *kelly.Context, obj interface{}, err error) {
	if err != nil {
		c.WriteIndentedJson(http.StatusOK, kelly.H{
			"message": err.Error(),
			"result":  http.StatusBadRequest,
		})
		return
	}

	c.WriteIndentedJson(http.StatusOK, kelly.H{
		"message": "success",
		"result":  0,
		"data":    obj,
	})
}
