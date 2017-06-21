package mch

import (
	"net/http"
	"strings"
	"time"
	"encoding/json"
	"log"
	"encoding/xml"
	"github.com/qjw/go-wx-sdk/mch"
	"github.com/qjw/go-wx-sdk/utils"
	"github.com/qjw/go-wx-sdk-app/const_var"
	"github.com/qjw/kelly"
	"github.com/qjw/kelly/middleware/swagger"
)


func realPath(group kelly.Router, path string) string {
	var burl = group.Path()
	if strings.HasPrefix(path, "/") {
		path = burl + path
	} else {
		path = burl + "/" + path
	}
	return path
}

func param_path_process(c *kelly.Context,form interface{}) bool {
	if err, _ := c.BindPath(form); err != nil {
		c.WriteIndentedJson(http.StatusBadRequest, kelly.H{
			"message": "params error",
			"result": http.StatusBadRequest,
		})
		return false
	}
	return true
}

func param_process(c *kelly.Context,form interface{}) bool{
	if err, _ := c.Bind(form); err != nil {
		c.WriteIndentedJson(http.StatusBadRequest, kelly.H{
			"message": "params error",
			"result": http.StatusBadRequest,
		})
		return false
	}
	return true
}

func end_process(c *kelly.Context,obj interface{},err error){
	if err != nil{
		c.WriteIndentedJson(http.StatusOK, kelly.H{
			"message": err.Error(),
			"result": http.StatusBadRequest,
		})
		return
	}

	c.WriteIndentedJson(http.StatusOK, kelly.H{
		"message": "success",
		"result": 0,
		"data": obj,
	})
}

func createOrder(c *kelly.Context,urouter kelly.Router, wechatApi *mch.PayApi) (*mch.UnifiedordeRes, error){
	var form mch.UnifiedordeObj
	if !param_process(c,&form){
		return nil,nil
	}

	ipaddr := c.Request().RemoteAddr
	if len(ipaddr) > 0{
		ipaddr = strings.Split(ipaddr,":")[0]
	}

	form.SpbillCreateIp = ipaddr
	ctime := time.Now()
	form.TimeStart = ctime.Format("20060102150405")
	form.TimeExpire = ctime.Add(time.Hour).Format("20060102150405")
	form.NotifyUrl = utils.CharData(const_var.HostUrl + realPath(urouter,"pay_callback"))

	return wechatApi.CreateUnifiedOrder(&form)
}

func InitializeApiRoutes(grouter kelly.Router, context *mch.Context) {
	urouter := grouter.Group("/mch")
	wechatApi := mch.NewPayApi(context)

	urouter.Annotation(swagger.Swagger(&swagger.StructParam{
		JsonData:&mch.UnifiedordeObj{},
		ResponseData:&swagger.SuccessResp{},
		Tags:[]string{"微信支付"},
		Summary:"创建订单",
	})).POST("/create_unifiedorder",func(c *kelly.Context){
		resp,err := createOrder(c,urouter,wechatApi)
		if resp != nil || err != nil {
			end_process(c,resp,err)
		}
	})

	urouter.Annotation(swagger.Swagger(&swagger.StructParam{
		FormData:&mch.OrderCloseParam{},
		ResponseData:&swagger.SuccessResp{},
		Tags:[]string{"微信支付"},
		Summary:"关闭订单",
	})).POST("/close_unifiedorder",func(c *kelly.Context) {
		var form mch.OrderCloseParam
		if !param_process(c,&form){
			return
		}

		resp,err := wechatApi.CloseUnifiedOrder(&form)
		end_process(c,resp,err)
	})

	urouter.Annotation(swagger.Swagger(&swagger.StructParam{
		JsonData:&mch.UnifiedordeObj{},
		ResponseData:&swagger.SuccessResp{},
		Tags:[]string{"微信支付"},
		Summary:"创建订单，并返回微信支付相关的参数",
	})).POST("/create_js_unifiedorder",func(c *kelly.Context) {
		resp,err := createOrder(c,urouter,wechatApi)
		if err != nil{
			end_process(c,nil,err)
			return
		}
		if (resp == nil && err == nil) {
			return
		}

		if resp.ReturnCode != "SUCCESS"{
			end_process(c,&mch.JsUnifiedOrderRes{
				OrderCommonError: mch.OrderCommonError{
					ReturnCode:resp.ReturnCode,
					ReturnMsg:resp.ReturnMsg,
				},
			},nil)
			return
		}else if resp.ErrCode != ""{
			end_process(c,&mch.JsUnifiedOrderRes{
				OrderCommonError: mch.OrderCommonError{
					ReturnCode:resp.ErrCode,
					ReturnMsg:resp.ErrCodeDes,
				},
			},nil)
			return
		}

		resp2,err := wechatApi.CreateJsUnifiedOrder(resp)
		end_process(c,resp2,err)
	})

	urouter.POST("/pay_callback",func(c *kelly.Context) {
		var data mch.UnifiedOrderNotifyObj
		decoder := xml.NewDecoder(c.Request().Body)
		if err := decoder.Decode(&data); err != nil {
			log.Print(err.Error())
			c.WriteXml(http.StatusOK, kelly.H{
				"return_msg": err.Error(),
				"return_code": "FAIL",
			})
			return
		}

		sign := data.Sign
		data.Sign = ""
		newSign,_ := wechatApi.Sign(&data)
		if sign != newSign{
			c.WriteXml(http.StatusOK, kelly.H{
				"return_msg": "FAIL",
				"return_code": "SUCCESS",
			})
			return
		}

		jsonData,_ := json.MarshalIndent(&data,"", " ")
		log.Print(string(jsonData))

		c.WriteXml(http.StatusOK, kelly.H{
			"return_msg": "OK",
			"return_code": "SUCCESS",
		})
	})
}