package mp

import (
	"net/http"
	"strconv"
	"io"
	"fmt"
	"bytes"
	"strings"
	"encoding/json"
	"log"
	"github.com/qjw/go-wx-sdk/mp"
	"github.com/qjw/go-wx-sdk-app/const_var"
	"github.com/qjw/kelly"
	"github.com/qjw/kelly/middleware/swagger"
)

func InitializeApiRoutes(grouter kelly.Router, context *mp.Context) {
	urouter := grouter.Group("/mp")
	wechatApi := mp.NewWechatApi(context)
	//------------------------------------------------------------------------
	urouter.Annotation(swagger.Swagger(&swagger.StructParam{
		ResponseData:&swagger.SuccessResp{},
		Tags:[]string{"杂项"},
		Summary:"获取Access Token",
	})).GET("/token",func(c *kelly.Context) {
		resp,err := wechatApi.GetAccessToken()
		end_process(c,resp,err)
	})

	urouter.Annotation(swagger.Swagger(&swagger.StructParam{
		ResponseData:&swagger.SuccessResp{},
		Tags:[]string{"杂项"},
		Summary:"获取Js Ticket",
	})).GET("/jsticket",func(c *kelly.Context) {
		resp,err := wechatApi.GetJsTicket()
		end_process(c,resp,err)
	})

	urouter.Annotation(swagger.Swagger(&swagger.StructParam{
		ResponseData:&swagger.SuccessResp{},
		Tags:[]string{"杂项"},
		Summary:"获取Card Ticket",
	})).GET("/cardticket",func(c *kelly.Context) {
		resp,err := wechatApi.GetCardTicket()
		end_process(c,resp,err)
	})

	urouter.Annotation(swagger.Swagger(&swagger.StructParam{
		ResponseData:&swagger.SuccessResp{},
		Tags:[]string{"杂项"},
		Summary:"获取微信服务器IP地址",
	})).GET("/iplist",func(c *kelly.Context) {
		resp,err := wechatApi.GetIpList()
		end_process(c,resp,err)
	})

	urouter.Annotation(swagger.Swagger(&swagger.StructParam{
		JsonData:&SignJsTicketParam{},
		ResponseData:&swagger.SuccessResp{},
		Tags:[]string{"杂项"},
		Summary:"Js签名",
	})).POST("/sign_js_ticket",func(c *kelly.Context) {
		var form SignJsTicketParam
		if !param_process(c,&form){
			return
		}

		resp,err := wechatApi.SignJsTicket(form.NonceStr,form.Timestamp,form.Url)
		end_process(c,resp,err)
	})

	//------------------------------------------------------------------------

	urouter.Annotation(swagger.Swagger(&swagger.StructParam{
		ResponseData:&swagger.SuccessResp{},
		Tags:[]string{"菜单"},
		Summary:"获取菜单",
	})).GET("/get_menu",func(c *kelly.Context) {
		resp,err := wechatApi.GetMenu()
		end_process(c,resp,err)
	})

	urouter.Annotation(swagger.Swagger(&swagger.StructParam{
		JsonData:&CreateMenuParam{},
		ResponseData:&swagger.SuccessResp{},
		Tags:[]string{"菜单"},
		Summary:"创建菜单",
	})).POST("/create_menu",func(c *kelly.Context) {
		var form CreateMenuParam
		if !param_process(c,&form){
			return
		}

		resp,err := wechatApi.CreateMenu(form.Menus)
		end_process(c,resp,err)
	})

	urouter.Annotation(swagger.Swagger(&swagger.StructParam{
		ResponseData:&swagger.SuccessResp{},
		Tags:[]string{"菜单"},
		Summary:"删除菜单",
	})).DELETE("/delete_menu",func(c *kelly.Context) {
		resp,err := wechatApi.DeleteMenu()
		end_process(c,resp,err)
	})

	//------------------------------------------------------------------------

	urouter.Annotation(swagger.Swagger(&swagger.StructParam{
		FormData:&UpdateKfAccountParam{},
		ResponseData:&swagger.SuccessResp{},
		Tags:[]string{"客服消息"},
		Summary:"添加客服帐号",
	})).POST("/update_kf_account",func(c *kelly.Context) {
		var form UpdateKfAccountParam
		if !param_process(c,&form){
			return
		}

		resp,err := wechatApi.UpdateKfAccount(form.KfAccount,form.Nickname,form.Password)
		end_process(c,resp,err)
	})

	urouter.Annotation(swagger.Swagger(&swagger.StructParam{
		FormData:&DelKfAccountParam{},
		ResponseData:&swagger.SuccessResp{},
		Tags:[]string{"客服消息"},
		Summary:"删除客服帐号",
	})).POST("/del_kf_account",func(c *kelly.Context) {
		var form DelKfAccountParam
		if !param_process(c,&form){
			return
		}

		resp,err := wechatApi.DelKfAccount(form.KfAccount)
		end_process(c,resp,err)
	})

	urouter.Annotation(swagger.Swagger(&swagger.StructParam{
		FormData:&AddKfAccountParam{},
		ResponseData:&swagger.SuccessResp{},
		Tags:[]string{"客服消息"},
		Summary:"添加客服帐号",
	})).POST("/add_kf_account",func(c *kelly.Context) {
		var form AddKfAccountParam
		if !param_process(c,&form){
			return
		}

		resp,err := wechatApi.AddKfAccount(form.KfAccount,form.Nickname,form.Password)
		end_process(c,resp,err)
	})

	urouter.Annotation(swagger.Swagger(&swagger.StructParam{
		ResponseData:&swagger.SuccessResp{},
		Tags:[]string{"客服消息"},
		Summary:"获取所有客服账号",
	})).GET("/get_kf_list",func(c *kelly.Context) {
		resp,err := wechatApi.GetKfList()
		end_process(c,resp,err)
	})

	urouter.Annotation(swagger.Swagger(&swagger.StructParam{
		FormData:&SendKfMsgParam{},
		ResponseData:&swagger.SuccessResp{},
		Tags:[]string{"客服消息"},
		Summary:"发送文本客服消息",
	})).POST("/send_kf_msg",func(c *kelly.Context) {
		var form SendKfMsgParam
		if !param_process(c,&form){
			return
		}

		resp,err := wechatApi.SendKfMessage(form.ToUser,form.Content)
		end_process(c,resp,err)
	})

	urouter.Annotation(swagger.Swagger(&swagger.StructParam{
		FormData:&SendKfImageMsgParam{},
		ResponseData:&swagger.SuccessResp{},
		Tags:[]string{"客服消息"},
		Summary:"发送图片客服消息",
	})).POST("/send_kf_image_msg",func(c *kelly.Context) {
		var form SendKfImageMsgParam
		if !param_process(c,&form){
			return
		}

		resp,err := wechatApi.SendKfImageMessage(form.ToUser,form.MediaID)
		end_process(c,resp,err)
	})

	urouter.Annotation(swagger.Swagger(&swagger.StructParam{
		JsonData:&SendKfVideoMsgParam{},
		ResponseData:&swagger.SuccessResp{},
		Tags:[]string{"客服消息"},
		Summary:"发送视频客服消息",
	})).POST("/send_kf_video_msg",func(c *kelly.Context) {
		var form SendKfVideoMsgParam
		if !param_process(c,&form){
			return
		}

		resp,err := wechatApi.SendKfVideoMessage(form.ToUser,
			form.MediaID,form.ThumbMediaID,form.Name,form.Description)
		end_process(c,resp,err)
	})

	urouter.Annotation(swagger.Swagger(&swagger.StructParam{
		FormData:&SendKfImageMsgParam{},
		ResponseData:&swagger.SuccessResp{},
		Tags:[]string{"客服消息"},
		Summary:"发送图文客服消息 - 基于素材",
	})).POST("/send_kf_mpnews_msg",func(c *kelly.Context) {
		var form SendKfImageMsgParam
		if !param_process(c,&form){
			return
		}

		resp,err := wechatApi.SendKfMpnewsMessage(form.ToUser,form.MediaID)
		end_process(c,resp,err)
	})

	urouter.Annotation(swagger.Swagger(&swagger.StructParam{
		FormData:&SendKfCardMsgParam{},
		ResponseData:&swagger.SuccessResp{},
		Tags:[]string{"客服消息"},
		Summary:"发送卡券客服消息",
	})).POST("/send_kf_card_msg",func(c *kelly.Context) {
		var form SendKfCardMsgParam
		if !param_process(c,&form){
			return
		}

		resp,err := wechatApi.SendKfCardMessage(form.ToUser,form.CardID)
		end_process(c,resp,err)
	})

	urouter.Annotation(swagger.Swagger(&swagger.StructParam{
		JsonData:&SendKfArticleMsgParam{},
		ResponseData:&swagger.SuccessResp{},
		Tags:[]string{"客服消息"},
		Summary:"发送图文客服消息",
	})).POST("/send_kf_article_msg",func(c *kelly.Context) {
		var form SendKfArticleMsgParam
		if !param_process(c,&form){
			return
		}

		resp,err := wechatApi.SendKfArticleMessage(form.ToUser,form.Articles)
		end_process(c,resp,err)
	})

	//-------------------------------------------------------------------------------------------------------------

	urouter.Annotation(swagger.Swagger(&swagger.StructParam{
		ResponseData:&swagger.SuccessResp{},
		Tags:[]string{"用户管理"},
		Summary:"获取用户openid列表",
	})).GET("/get_users",func(c *kelly.Context) {
		resp,err := wechatApi.GetUser("")
		end_process(c,resp,err)
	})

	urouter.Annotation(swagger.Swagger(&swagger.StructParam{
		QueryData:&GetUserDetailParam{},
		ResponseData:&swagger.SuccessResp{},
		Tags:[]string{"用户管理"},
		Summary:"获取用户详情",
	})).GET("/get_user_detail",func(c *kelly.Context) {
		var form GetUserDetailParam
		if !param_process(c,&form){
			return
		}

		resp,err := wechatApi.GetUserDetail(form.OpenID)
		end_process(c,resp,err)
	})

	urouter.Annotation(swagger.Swagger(&swagger.StructParam{
		JsonData:&BatchUserDetail{},
		ResponseData:&swagger.SuccessResp{},
		Tags:[]string{"用户管理"},
		Summary:"批量获取用户详情",
	})).POST("/batch_get_user_detail",func(c *kelly.Context) {
		var form BatchUserDetail
		if !param_process(c,&form){
			return
		}

		resp,err := wechatApi.BatchGetUserDetail(form.OpenIDList)
		end_process(c,resp,err)
	})

	urouter.Annotation(swagger.Swagger(&swagger.StructParam{
		ResponseData:&swagger.SuccessResp{},
		Tags:[]string{"用户管理"},
		Summary:"获取用户所有标签",
	})).GET("/get_tag_list",func(c *kelly.Context) {
		resp,err := wechatApi.GetTagList()
		end_process(c,resp,err)
	})

	urouter.Annotation(swagger.Swagger(&swagger.StructParam{
		FormData:&UpdateTagParam{},
		ResponseData:&swagger.SuccessResp{},
		Tags:[]string{"用户管理"},
		Summary:"更新标签",
	})).PUT("/update_tag",func(c *kelly.Context) {
		var form UpdateTagParam
		if !param_process(c,&form){
			return
		}

		resp,err := wechatApi.UpdateTag(form.ID,form.Name)
		end_process(c,resp,err)
	})

	urouter.Annotation(swagger.Swagger(&swagger.StructParam{
		FormData:&UserRemarkParam{},
		ResponseData:&swagger.SuccessResp{},
		Tags:[]string{"用户管理"},
		Summary:"设置用户备注名",
	})).POST("/update_user_tag",func(c *kelly.Context) {
		var form UserRemarkParam
		if !param_process(c,&form){
			return
		}

		resp,err := wechatApi.SetUserRemark(form.OpenID,form.Remark)
		end_process(c,resp,err)
	})

	urouter.Annotation(swagger.Swagger(&swagger.StructParam{
		FormData:&CreateTagParam{},
		ResponseData:&swagger.SuccessResp{},
		Tags:[]string{"用户管理"},
		Summary:"创建标签",
	})).POST("/create_tag",func(c *kelly.Context) {
		var form CreateTagParam
		if !param_process(c,&form){
			return
		}

		resp,err := wechatApi.CreateTag(form.Name)
		end_process(c,resp,err)
	})

	urouter.Annotation(swagger.Swagger(&swagger.StructParam{
		PathData:&DeleteTagParam{},
		ResponseData:&swagger.SuccessResp{},
		Tags:[]string{"用户管理"},
		Summary:"删除标签",
	})).DELETE("/delete_tag/:id",func(c *kelly.Context) {
		var form DeleteTagParam
		if !param_path_process(c,&form){
			return
		}

		resp,err := wechatApi.DeleteTag(form.ID)
		end_process(c,resp,err)
	})

	urouter.Annotation(swagger.Swagger(&swagger.StructParam{
		PathData:&TagUsers{},
		ResponseData:&swagger.SuccessResp{},
		Tags:[]string{"用户管理"},
		Summary:"获取标签下的所有用户",
	})).GET("/tag_users/:id",func(c *kelly.Context) {
		var form TagUsers
		if !param_path_process(c,&form){
			return
		}

		resp,err := wechatApi.GetTagUsers(form.ID,"")
		end_process(c,resp,err)
	})

	urouter.Annotation(swagger.Swagger(&swagger.StructParam{
		JsonData:&TagAddMember{},
		ResponseData:&swagger.SuccessResp{},
		Tags:[]string{"用户管理"},
		Summary:"批量为用户打标签",
	})).POST("/tag_add_members",func(c *kelly.Context) {
		var form TagAddMember
		if !param_process(c,&form){
			return
		}

		resp,err := wechatApi.AddTagMembers(form.TagID,form.OpenIDList)
		end_process(c,resp,err)
	})

	urouter.Annotation(swagger.Swagger(&swagger.StructParam{
		JsonData:&TagRmMember{},
		ResponseData:&swagger.SuccessResp{},
		Tags:[]string{"用户管理"},
		Summary:"批量为用户取消标签",
	})).POST("/tag_rm_members",func(c *kelly.Context) {
		var form TagRmMember
		if !param_process(c,&form){
			return
		}

		res,err := wechatApi.RemoveTagMembers(form.TagID,form.OpenIDList)
		end_process(c,res,err)
	})

	urouter.Annotation(swagger.Swagger(&swagger.StructParam{
		QueryData:&GetUserDetailParam{},
		ResponseData:&swagger.SuccessResp{},
		Tags:[]string{"用户管理"},
		Summary:"获取用户身上的标签列表",
	})).GET("/get_user_tags",func(c *kelly.Context) {
		var form GetUserDetailParam
		if !param_process(c,&form){
			return
		}

		res,err := wechatApi.GetUserTags(form.OpenID)
		end_process(c,res,err)
	})

	urouter.Annotation(swagger.Swagger(&swagger.StructParam{
		ResponseData:&swagger.SuccessResp{},
		Tags:[]string{"素材管理"},
		Summary:"获取素材总数",
	})).GET("/get_materialcount",func(c *kelly.Context) {
		res,err := wechatApi.GetMaterialCount()
		end_process(c,res,err)
	})

	urouter.Annotation(swagger.Swagger(&swagger.StructParam{
		QueryData:&GetMaterialsParam{},
		ResponseData:&swagger.SuccessResp{},
		Tags:[]string{"素材管理"},
		Summary:"获取永久素材列表",
	})).GET("/get_materials",func(c *kelly.Context) {
		var form GetMaterialsParam
		if !param_process(c,&form){
			return
		}

		res,err := wechatApi.GetMaterials(form.Tp,form.Offset,form.Count)
		end_process(c,res,err)
	})

	urouter.Annotation(swagger.Swagger(&swagger.StructParam{
		QueryData:&GetNewsMaterialsParam{},
		ResponseData:&swagger.SuccessResp{},
		Tags:[]string{"素材管理"},
		Summary:"获取永久图文素材列表",
	})).GET("/get_news_materials",func(c *kelly.Context) {
		var form GetNewsMaterialsParam
		if !param_process(c,&form){
			return
		}

		res,err := wechatApi.GetNewsMaterials(form.Offset,form.Count)
		end_process(c,res,err)
	})

	//--------------------------------------------------------------------------------------------------------------

	urouter.Annotation(swagger.Swagger(&swagger.StructParam{
		QueryData:&GetBlackListParam{},
		ResponseData:&swagger.SuccessResp{},
		Tags:[]string{"黑名单管理"},
		Summary:"获取公众号的黑名单列表",
	})).GET("/get_blacklist",func(c *kelly.Context) {
		var form GetBlackListParam
		if !param_process(c,&form){
			return
		}

		res,err := wechatApi.GetBlacklist(form.NextOpenid)
		end_process(c,res,err)
	})

	urouter.Annotation(swagger.Swagger(&swagger.StructParam{
		JsonData:&BatchUnblacklist{},
		ResponseData:&swagger.SuccessResp{},
		Tags:[]string{"黑名单管理"},
		Summary:"取消拉黑用户",
	})).POST("/batch_unblacklist",func(c *kelly.Context) {
		var form BatchBlacklist
		if !param_process(c,&form){
			return
		}
		resp,err := wechatApi.BatchUnblacklist(form.OpenIDList)
		end_process(c,resp,err)
	})

	urouter.Annotation(swagger.Swagger(&swagger.StructParam{
		JsonData:&BatchBlacklist{},
		ResponseData:&swagger.SuccessResp{},
		Tags:[]string{"黑名单管理"},
		Summary:"拉黑用户",
	})).POST("/batch_blacklist",func(c *kelly.Context) {
		var form BatchBlacklist
		if !param_process(c,&form){
			return
		}
		resp,err := wechatApi.Batchblacklist(form.OpenIDList)
		end_process(c,resp,err)
	})

	//-----------------------------------模板------------------------------------------------------------------------

	urouter.Annotation(swagger.Swagger(&swagger.StructParam{
		ResponseData:&swagger.SuccessResp{},
		Tags:[]string{"模板消息"},
		Summary:"获取模板列表",
	})).GET("/get_all_private_template",func(c *kelly.Context) {
		res,err := wechatApi.GetAllPrivateTemplate()
		end_process(c,res,err)
	})

	urouter.Annotation(swagger.Swagger(&swagger.StructParam{
		ResponseData:&swagger.SuccessResp{},
		Tags:[]string{"模板消息"},
		Summary:"获取设置的行业信息",
	})).GET("/get_industry",func(c *kelly.Context) {
		res,err := wechatApi.GetIndustry()
		end_process(c,res,err)
	})

	urouter.Annotation(swagger.Swagger(&swagger.StructParam{
		FormData:&AddTemplateParam{},
		ResponseData:&swagger.SuccessResp{},
		Tags:[]string{"模板消息"},
		Summary:"添加模板 - 获得模板ID",
	})).POST("/add_template",func(c *kelly.Context) {
		var form AddTemplateParam
		if !param_process(c,&form){
			return
		}
		resp,err := wechatApi.AddTemplate(form.TemplateIDShort)
		end_process(c,resp,err)
	})

	urouter.Annotation(swagger.Swagger(&swagger.StructParam{
		FormData:&DeleteTemplateParam{},
		ResponseData:&swagger.SuccessResp{},
		Tags:[]string{"模板消息"},
		Summary:"删除模板",
	})).POST("/delete_template",func(c *kelly.Context) {
		var form DeleteTemplateParam
		if !param_process(c,&form){
			return
		}
		resp,err := wechatApi.DeleteTemplate(form.TemplateID)
		end_process(c,resp,err)
	})

	urouter.Annotation(swagger.Swagger(&swagger.StructParam{
		JsonData:&     mp.TemplateSend{},
		ResponseData:& swagger.SuccessResp{},
		Tags:         []string{"模板消息"},
		Summary:      "发送模板消息",
	})).POST("/send_template",func(c *kelly.Context) {
		var form mp.TemplateSend
		if !param_process(c,&form){
			return
		}
		resp,err := wechatApi.SendTemplateMsg(&form)
		end_process(c,resp,err)
	})

	//---------------------------------------------------------------------------------------------------

	urouter.Annotation(swagger.Swagger(&swagger.StructParam{
		FormData:&LimitStrQrcodeParam{},
		ResponseData:&swagger.SuccessResp{},
		Tags:[]string{"二维码"},
		Summary:"申请永久字符串二维码",
	})).POST("/create_Limit_str_qrcode",func(c *kelly.Context) {
		var form LimitStrQrcodeParam
		if !param_process(c,&form){
			return
		}

		resp,err := wechatApi.CreateLimitStrQrcode(form.SceneStr)
		end_process(c,resp,err)
	})

	urouter.Annotation(swagger.Swagger(&swagger.StructParam{
		FormData:&LimitQrcodeParam{},
		ResponseData:&swagger.SuccessResp{},
		Tags:[]string{"二维码"},
		Summary:"申请永久二维码",
	})).POST("/create_Limit_qrcode",func(c *kelly.Context) {
		var form LimitQrcodeParam
		if !param_process(c,&form){
			return
		}

		resp,err := wechatApi.CreateLimitQrcode(form.SceneID)
		end_process(c,resp,err)
	})

	urouter.Annotation(swagger.Swagger(&swagger.StructParam{
		FormData:&QrcodeParam{},
		ResponseData:&swagger.SuccessResp{},
		Tags:[]string{"二维码"},
		Summary:"临时二维码",
	})).POST("/create_qrcode",func(c *kelly.Context) {
		var form QrcodeParam
		if !param_process(c,&form){
			return
		}

		resp,err := wechatApi.CreateQrcode(form.ExpireSeconds,form.SceneID)
		end_process(c,resp,err)
	})

	urouter.Annotation(
		swagger.SwaggerFile("mp.yaml:show_qrcode"),
	).GET("/show_qrcode",func(c *kelly.Context) {
		var form ShowQrcodeParam
		if !param_process(c,&form){
			return
		}
		url := wechatApi.ShowQrcode(form.Ticket)
		// don't worry about errors
		response, err := http.Get(url);
		if err != nil{
			c.WriteIndentedJson(http.StatusOK, kelly.H{
				"message": err.Error(),
				"result": http.StatusBadRequest,
			})
			return
		}
		defer response.Body.Close()
		c.SetHeader("Content-Type", "image/jpeg")
		c.SetHeader("Content-Length", strconv.FormatInt(response.ContentLength,10))
		io.Copy(c, response.Body)
	})

	urouter.Annotation(swagger.Swagger(&swagger.StructParam{
		FormData:&ShortUrlParam{},
		ResponseData:&swagger.SuccessResp{},
		Tags:[]string{"二维码"},
		Summary:"长链接转短链接",
	})).POST("/short_url",func(c *kelly.Context) {
		var form ShortUrlParam
		if !param_process(c,&form){
			return
		}

		resp,err := wechatApi.ShortUrl(form.LongUrl)
		end_process(c,resp,err)
	})

	urouter.Annotation(
		swagger.SwaggerFile("mp.yaml:upload_tmp_material"),
	).POST("/upload_tmp_material",func(c *kelly.Context) {
		file, multipart := c.MustGetFileVarible("file")
		defer file.Close()
		tp := c.MustGetFormVarible("type")
		resp,err := wechatApi.UploadTmpMaterial(file,multipart.Filename,tp)
		end_process(c,resp,err)
	})

	urouter.Annotation(
		swagger.SwaggerFile("mp.yaml:get_tmp_material"),
	).GET("/get_tmp_material",func(c *kelly.Context) {
		var form GetTmpMaterialParam
		if !param_process(c,&form){
			return
		}
		url,_ := wechatApi.GetTmpMaterial(form.MediaID)
		// don't worry about errors
		response, err := http.Get(url);
		if err != nil{
			c.WriteIndentedJson(http.StatusOK, kelly.H{
				"message": err.Error(),
				"result": http.StatusBadRequest,
			})
			return
		}
		defer response.Body.Close()
		c.SetHeader("Content-Type", "image/jpeg")
		c.SetHeader("Content-Length", strconv.FormatInt(response.ContentLength,10))
		io.Copy(c, response.Body)
	})

	urouter.Annotation(swagger.Swagger(&swagger.StructParam{
		QueryData:&GetTmpMaterialParam{},
		ResponseData:&swagger.SuccessResp{},
		Tags:[]string{"素材管理"},
		Summary:"获取视频临时素材",
	})).GET("/get_video_tmp_material",func(c *kelly.Context) {
		var form GetTmpMaterialParam
		if !param_process(c,&form){
			return
		}
		res,err := wechatApi.GetVideoTmpMaterial(form.MediaID)
		end_process(c,res,err)
	})

	urouter.Annotation(
		swagger.SwaggerFile("mp.yaml:upload_material"),
	).POST("/upload_material",func(c *kelly.Context) {
		file, multipart := c.MustGetFileVarible("file")
		defer file.Close()
		tp := c.MustGetFormVarible("type")
		resp,err := wechatApi.UploadMaterial(file,multipart.Filename,tp)
		end_process(c,resp,err)
	})

	urouter.Annotation(
		swagger.SwaggerFile("mp.yaml:upload_video_material"),
	).POST("/upload_video_material",func(c *kelly.Context) {
		file, multipart := c.MustGetFileVarible("file")
		defer file.Close()

		title := c.MustGetFormVarible("title")
		description := c.MustGetFormVarible("description")

		resp,err := wechatApi.UploadVideoMaterial(file,multipart.Filename,&mp.VideoMaterialInfo{
			Title: title,
			Description: description,
		})
		end_process(c,resp,err)
	})


	const (
		get_material_temp = `{
			"media_id":"%s"
		}`
	)

	urouter.Annotation(
		swagger.SwaggerFile("mp.yaml:get_material"),
	).GET("/get_material",func(c *kelly.Context) {
		var form GetTmpMaterialParam
		if !param_process(c,&form){
			return
		}
		uri,_ := wechatApi.GetMaterial()
		body := []byte(fmt.Sprintf(get_material_temp, form.MediaID))
		response, err := http.Post(uri, "application/json;charset=utf-8", bytes.NewBuffer(body))
		if err != nil {
			c.WriteIndentedJson(http.StatusOK, kelly.H{
				"message": "post fail",
				"result": 100,
			})
			return
		}
		defer response.Body.Close()
		c.SetHeader("Content-Type", "image/jpeg")
		c.SetHeader("Content-Length", strconv.FormatInt(response.ContentLength,10))
		io.Copy(c, response.Body)
	})

	urouter.Annotation(swagger.Swagger(&swagger.StructParam{
		QueryData:&GetMaterialParam{},
		ResponseData:&swagger.SuccessResp{},
		Tags:[]string{"素材管理"},
		Summary:"获取视频素材",
	})).GET("/get_video_material",func(c *kelly.Context) {
		var form GetMaterialParam
		if !param_process(c,&form){
			return
		}
		res,err := wechatApi.GetVideoMaterial(form.MediaID)
		end_process(c,res,err)
	})

	urouter.Annotation(swagger.Swagger(&swagger.StructParam{
		QueryData:&GetMaterialParam{},
		ResponseData:&swagger.SuccessResp{},
		Tags:[]string{"素材管理"},
		Summary:"获取图文素材",
	})).GET("/get_article_material",func(c *kelly.Context) {
		var form GetMaterialParam
		if !param_process(c,&form){
			return
		}
		res,err := wechatApi.GetArticleMaterial(form.MediaID)
		end_process(c,res,err)
	})

	urouter.Annotation(swagger.Swagger(&swagger.StructParam{
		FormData:&DeleteMaterialParam{},
		ResponseData:&swagger.SuccessResp{},
		Tags:[]string{"素材管理"},
		Summary:"删除永久素材",
	})).POST("/del_video_material",func(c *kelly.Context) {
		var form DeleteMaterialParam
		if !param_process(c,&form){
			return
		}
		res,err := wechatApi.DeleteMaterial(form.MediaID)
		end_process(c,res,err)
	})

	urouter.Annotation(
		swagger.SwaggerFile("mp.yaml:upload_article_image"),
	).POST("/upload_article_image",func(c *kelly.Context) {
		file, multipart := c.MustGetFileVarible("file")
		defer file.Close()
		resp,err := wechatApi.UploadArticleImage(file,multipart.Filename)
		end_process(c,resp,err)
	})

	urouter.Annotation(swagger.Swagger(&swagger.StructParam{
		JsonData:&     mp.ArticleCreate{},
		ResponseData:& swagger.SuccessResp{},
		Tags:         []string{"素材管理"},
		Summary:      "创建图文消息",
	})).POST("/create_article_material",func(c *kelly.Context) {
		var form mp.ArticleCreate
		if !param_process(c,&form){
			return
		}
		resp,err := wechatApi.CreateArticle(form.Articles)
		end_process(c,resp,err)
	})

	urouter.Annotation(swagger.Swagger(&swagger.StructParam{
		FormData:&MassPreviewTextParam{},
		ResponseData:&swagger.SuccessResp{},
		Tags:[]string{"群发管理"},
		Summary:"群发文本预览",
	})).POST("/mass_preview_content",func(c *kelly.Context) {
		var form MassPreviewTextParam
		if !param_process(c,&form){
			return
		}
		resp,err := wechatApi.MassPreviewText(form.OpenID,form.Content)
		end_process(c,resp,err)
	})

	urouter.Annotation(swagger.Swagger(&swagger.StructParam{
		FormData:&MassPreviewParam{},
		ResponseData:&swagger.SuccessResp{},
		Tags:[]string{"群发管理"},
		Summary:"群发预览",
	})).POST("/mass_preview",func(c *kelly.Context) {
		var form MassPreviewParam
		if !param_process(c,&form){
			return
		}
		resp,err := wechatApi.MassPreviewMsg(form.OpenID,form.MediaID,form.Tp)
		end_process(c,resp,err)
	})

	urouter.Annotation(swagger.Swagger(&swagger.StructParam{
		JsonData:&MassSendMpnewsParam{},
		ResponseData:&swagger.SuccessResp{},
		Tags:[]string{"群发管理"},
		Summary:"基于openid列表的群发图文消息",
	})).POST("/mass_send_article",func(c *kelly.Context) {
		var form MassSendMpnewsParam
		if !param_process(c,&form){
			return
		}
		resp,err := wechatApi.MassSendMpnews(form.OpenIDList,form.MediaID,form.SendIgnoreReprint)
		end_process(c,resp,err)
	})

	urouter.Annotation(swagger.Swagger(&swagger.StructParam{
		JsonData:&MassSendTextParam{},
		ResponseData:&swagger.SuccessResp{},
		Tags:[]string{"群发管理"},
		Summary:"基于openid列表的群发文本消息",
	})).POST("/mass_send_text",func(c *kelly.Context) {
		var form MassSendTextParam
		if !param_process(c,&form){
			return
		}
		resp,err := wechatApi.MassSendText(form.OpenIDList,form.Content)
		end_process(c,resp,err)
	})

	urouter.Annotation(swagger.Swagger(&swagger.StructParam{
		JsonData:&MassSendCardParam{},
		ResponseData:&swagger.SuccessResp{},
		Tags:[]string{"群发管理"},
		Summary:"基于openid列表的群发卡券消息",
	})).POST("/mass_send_card",func(c *kelly.Context) {
		var form MassSendCardParam
		if !param_process(c,&form){
			return
		}
		resp,err := wechatApi.MassSendCard(form.OpenIDList,form.CardID)
		end_process(c,resp,err)
	})

	urouter.Annotation(swagger.Swagger(&swagger.StructParam{
		JsonData:&MassSendMsgParam{},
		ResponseData:&swagger.SuccessResp{},
		Tags:[]string{"群发管理"},
		Summary:"基于openid列表的群发voice/image消息",
	})).POST("/mass_send_msg",func(c *kelly.Context) {
		var form MassSendMsgParam
		if !param_process(c,&form){
			return
		}
		resp,err := wechatApi.MassSendMsg(form.OpenIDList,form.MediaID,form.Tp)
		end_process(c,resp,err)
	})

	urouter.Annotation(swagger.Swagger(&swagger.StructParam{
		QueryData:&MassGetParam{},
		ResponseData:&swagger.SuccessResp{},
		Tags:[]string{"群发管理"},
		Summary:"查询群发消息发送状态",
	})).GET("/mass_get",func(c *kelly.Context) {
		var form MassGetParam
		if !param_process(c,&form){
			return
		}
		resp,err := wechatApi.MassGet(form.MsgID)
		end_process(c,resp,err)
	})

	urouter.Annotation(swagger.Swagger(&swagger.StructParam{
		PathData:&MassDeleteParam{},
		ResponseData:&swagger.SuccessResp{},
		Tags:[]string{"群发管理"},
		Summary:"删除群发",
	})).DELETE("/mass_delete/:msg_id",func(c *kelly.Context) {
		var form MassDeleteParam
		if !param_path_process(c,&form){
			return
		}
		resp,err := wechatApi.MassDelete(form.MsgID)
		end_process(c,resp,err)
	})

	// --------

	urouter.Annotation(swagger.Swagger(&swagger.StructParam{
		FormData:&MassSendAllMpnewsParam{},
		ResponseData:&swagger.SuccessResp{},
		Tags:[]string{"群发管理"},
		Summary:"基于tag的群发图文消息",
	})).POST("/mass_sendall_article",func(c *kelly.Context) {
		var form MassSendAllMpnewsParam
		if !param_process(c,&form){
			return
		}
		resp,err := wechatApi.MassSendAllMpnews(form.MediaID,form.TagID,form.IsToall,form.SendIgnoreReprint)
		end_process(c,resp,err)
	})

	urouter.Annotation(swagger.Swagger(&swagger.StructParam{
		FormData:&MassSendAllTextParam{},
		ResponseData:&swagger.SuccessResp{},
		Tags:[]string{"群发管理"},
		Summary:"基于tag的群发文本消息",
	})).POST("/mass_sendall_text",func(c *kelly.Context) {
		var form MassSendAllTextParam
		if !param_process(c,&form){
			return
		}
		resp,err := wechatApi.MassSendAllText(form.Content,form.TagID,form.IsToall)
		end_process(c,resp,err)
	})

	urouter.Annotation(swagger.Swagger(&swagger.StructParam{
		FormData:&MassSendAllCardParam{},
		ResponseData:&swagger.SuccessResp{},
		Tags:[]string{"群发管理"},
		Summary:"基于tag的群发卡券消息",
	})).POST("/mass_sendall_card",func(c *kelly.Context) {
		var form MassSendAllCardParam
		if !param_process(c,&form){
			return
		}
		resp,err := wechatApi.MassSendAllCard(form.CardID,form.TagID,form.IsToall)
		end_process(c,resp,err)
	})

	urouter.Annotation(swagger.Swagger(&swagger.StructParam{
		FormData:&MassSendAllMsgParam{},
		ResponseData:&swagger.SuccessResp{},
		Tags:[]string{"群发管理"},
		Summary:"基于tag的群发voice/image消息",
	})).POST("/mass_sendall_msg",func(c *kelly.Context) {
		var form MassSendAllMsgParam
		if !param_process(c,&form){
			return
		}
		resp,err := wechatApi.MassSendAllMsg(form.MediaID,form.TagID,form.IsToall,form.Tp)
		end_process(c,resp,err)
	})

	urouter.GET("/snsapi_base",func(c *kelly.Context) {
		url := const_var.HostUrl + realPath(urouter,"snsapi_base_callback")
		url1 := wechatApi.AuthorizeBase(url,"snsapi_base_test")
		c.Redirect(http.StatusFound, url1)
	})

	urouter.GET("/snsapi_base_callback",func(c *kelly.Context) {
		var form AuthorizeBaseParam
		if !param_process(c,&form){
			return
		}
		resp,err := wechatApi.GetAuthorizeAccessToken(form.Code)
		end_process(c,resp,err)
	})

	urouter.GET("/snsapi_userinfo",func(c *kelly.Context) {
		url := const_var.HostUrl + realPath(urouter,"snsapi_userinfo_callback")
		url1 := wechatApi.AuthorizeUserinfo(url,"snsapi_userinfo_test")
		c.Redirect(http.StatusFound, url1)
	})

	urouter.GET("/snsapi_userinfo_callback",func(c *kelly.Context) {
		var form AuthorizeBaseParam
		if !param_process(c,&form){
			return
		}
		resp,err := wechatApi.GetAuthorizeAccessToken(form.Code)
		if err != nil{
			c.WriteIndentedJson(http.StatusOK, kelly.H{
				"message": err.Error(),
				"result": http.StatusBadRequest,
			})
			return
		}

		resp2,err := wechatApi.GetAuthorizeSnsUserinfo(resp.AccessToken,resp.OpenID)
		end_process(c,resp2,err)
	})


	//===============================店铺============================================================================

	urouter.Annotation(swagger.Swagger(&swagger.StructParam{
		QueryData:&mp.GetPoiListParam{},
		ResponseData:&swagger.SuccessResp{},
		Tags:[]string{"店铺和卡券"},
		Summary:"获取店铺列表",
	})).GET("/get_poi_list",func(c *kelly.Context) {
		var form mp.GetPoiListParam
		if !param_process(c,&form){
			return
		}

		resp,err := wechatApi.GetPoiList(&form)
		end_process(c,resp,err)
	})

	urouter.Annotation(swagger.Swagger(&swagger.StructParam{
		QueryData:&mp.GetPoiDetailParam{},
		ResponseData:&swagger.SuccessResp{},
		Tags:[]string{"店铺和卡券"},
		Summary:"获取店铺详情",
	})).GET("/get_poi_detail",func(c *kelly.Context) {
		var form mp.GetPoiDetailParam
		if !param_process(c,&form){
			return
		}

		resp,err := wechatApi.GetPoiDetail(&form)
		end_process(c,resp,err)
	})


	//===============================卡券============================================================================

	urouter.Annotation(swagger.Swagger(&swagger.StructParam{
		JsonData:&mp.GetCardListParam{},
		ResponseData:&swagger.SuccessResp{},
		Tags:[]string{"店铺和卡券"},
		Summary:"获取卡券列表",
	})).POST("/get_card_list",func(c *kelly.Context) {
		var form mp.GetCardListParam
		if !param_process(c,&form){
			return
		}

		resp,err := wechatApi.GetCardList(&form)
		end_process(c,resp,err)
	})

	urouter.Annotation(swagger.Swagger(&swagger.StructParam{
		QueryData:&mp.GetCardDetailParam{},
		ResponseData:&swagger.SuccessResp{},
		Tags:[]string{"店铺和卡券"},
		Summary:"获取卡券详情",
	})).GET("/get_card_detail",func(c *kelly.Context) {
		var form mp.GetCardDetailParam
		if !param_process(c,&form){
			return
		}

		resp,err := wechatApi.GetCardDetail(&form)
		end_process(c,resp,err)
	})

	urouter.Annotation(swagger.Swagger(&swagger.StructParam{
		JsonData:&mp.CreateCardQrcodeParam{},
		ResponseData:&swagger.SuccessResp{},
		Tags:[]string{"店铺和卡券"},
		Summary:"创建卡券二维码",
	})).POST("/create_card_qrcode",func(c *kelly.Context) {
		var form mp.CreateCardQrcodeParam
		if !param_process(c,&form){
			return
		}

		resp,err := wechatApi.CreateCardQrcode(&form)
		end_process(c,resp,err)
	})

	urouter.Annotation(swagger.Swagger(&swagger.StructParam{
		JsonData:&mp.CreateCardParam{},
		ResponseData:&swagger.SuccessResp{},
		Tags:[]string{"店铺和卡券"},
		Summary:"创建卡券",
	})).POST("/create_card",func(c *kelly.Context) {
		var form mp.CreateCardParam
		if !param_process(c,&form){
			return
		}

		resp,err := wechatApi.CreateCard(&form)
		end_process(c,resp,err)
	})

	urouter.Annotation(swagger.Swagger(&swagger.StructParam{
		JsonData:&mp.SetCardWhitelistParam{},
		ResponseData:&swagger.SuccessResp{},
		Tags:[]string{"店铺和卡券"},
		Summary:"设置卡券白名单",
	})).POST("/set_card_whitelist",func(c *kelly.Context) {
		var form mp.SetCardWhitelistParam
		if !param_process(c,&form){
			return
		}

		resp,err := wechatApi.SetCardWhitelist(&form)
		end_process(c,resp,err)
	})

	urouter.Annotation(swagger.Swagger(&swagger.StructParam{
		FormData:&mp.CardIDParam{},
		ResponseData:&swagger.SuccessResp{},
		Tags:[]string{"店铺和卡券"},
		Summary:"删除卡券",
	})).POST("/delete_card",func(c *kelly.Context) {
		var form mp.CardIDParam
		if !param_process(c,&form){
			return
		}
		resp,err := wechatApi.DeleteCard(&form)
		end_process(c,resp,err)
	})

	urouter.Annotation(swagger.Swagger(&swagger.StructParam{
		FormData:&mp.UnavailableCardCodeParam{},
		ResponseData:&swagger.SuccessResp{},
		Tags:[]string{"店铺和卡券"},
		Summary:"作废卡券code",
	})).POST("/unavailable_card_code",func(c *kelly.Context) {
		var form mp.UnavailableCardCodeParam
		if !param_process(c,&form){
			return
		}
		resp,err := wechatApi.UnavailableCardCode(&form)
		end_process(c,resp,err)
	})

	urouter.Annotation(swagger.Swagger(&swagger.StructParam{
		QueryData:&mp.CheckCardCodeParam{},
		ResponseData:&swagger.SuccessResp{},
		Tags:[]string{"店铺和卡券"},
		Summary:"获取卡券code状态",
	})).GET("/check_card_code",func(c *kelly.Context) {
		var form mp.CheckCardCodeParam
		if !param_process(c,&form){
			return
		}
		resp,err := wechatApi.CheckCardCode(&form)
		end_process(c,resp,err)
	})

	urouter.Annotation(swagger.Swagger(&swagger.StructParam{
		QueryData:&mp.GetCardUseListParam{},
		ResponseData:&swagger.SuccessResp{},
		Tags:[]string{"店铺和卡券"},
		Summary:"获取用户的领券情况",
	})).GET("/get_card_use_list",func(c *kelly.Context) {
		var form mp.GetCardUseListParam
		if !param_process(c,&form){
			return
		}
		resp,err := wechatApi.GetCardUseList(&form)
		end_process(c,resp,err)
	})

	urouter.Annotation(swagger.Swagger(&swagger.StructParam{
		JsonData:&mp.CardOuterParam{},
		ResponseData:&swagger.SuccessResp{},
		Tags:[]string{"店铺和卡券"},
		Summary:"解码卡券的code",
	})).POST("/decrypt_card_code",func(c *kelly.Context) {
		var form mp.DecryptCardCodeParam
		if !param_process(c,&form){
			return
		}
		resp,err := wechatApi.DecryptCardCode(&form)
		end_process(c,&resp,err)
	})

	urouter.GET("/card_extern_url",func(c *kelly.Context) {
		log.Print(c.Request().RequestURI)
		var form mp.CardOuterParam
		if !param_process(c,&form){
			return
		}

		resp,err := wechatApi.DecryptCardCode(&mp.DecryptCardCodeParam{
			EncryptCode:form.EncryptCode,
		})
		form.CardCode = resp.CardCode
		if err != nil{
			c.WriteIndentedJson(http.StatusOK, kelly.H{
				"message": err.Error(),
				"result": http.StatusBadRequest,
			})
			return
		}

		data,_ := json.MarshalIndent(&form,"", " ")
		log.Print(string(data))


		c.Redirect(http.StatusFound, "http://www.baidu.com")
	})

	urouter.Annotation(swagger.Swagger(&swagger.StructParam{
		JsonData:&mp.CreateCardLandingpageParam{},
		ResponseData:&swagger.SuccessResp{},
		Tags:[]string{"店铺和卡券"},
		Summary:"创建卡券货架",
	})).POST("/create_card_landingpage",func(c *kelly.Context) {
		var form mp.CreateCardLandingpageParam
		if !param_process(c,&form){
			return
		}
		resp,err := wechatApi.CreateCardLandingpage(&form)
		end_process(c,resp,err)
	})

	urouter.Annotation(swagger.Swagger(&swagger.StructParam{
		JsonData:&SignChooseCardParam{},
		ResponseData:&swagger.SuccessResp{},
		Tags:[]string{"杂项"},
		Summary:"wx.chooseCard签名",
	})).POST("/sign_choose_card",func(c *kelly.Context) {
		var form SignChooseCardParam
		if !param_process(c,&form){
			return
		}

		resp,err := wechatApi.SignChooseCard(form.ShopID,form.CardID,form.CardType)
		end_process(c,resp,err)
	})

	urouter.Annotation(swagger.Swagger(&swagger.StructParam{
		JsonData:&mp.SignAddCardParam{},
		ResponseData:&swagger.SuccessResp{},
		Tags:[]string{"杂项"},
		Summary:"wx.addCard签名",
	})).POST("/sign_add_card",func(c *kelly.Context) {
		var form mp.SignAddCardParam
		if !param_process(c,&form){
			return
		}

		resp,err := wechatApi.SignAddCard(&form)
		end_process(c,resp,err)
	})

	urouter.Annotation(swagger.Swagger(&swagger.StructParam{
		FormData:&mp.ConsumeCardCodeParam{},
		ResponseData:&swagger.SuccessResp{},
		Tags:[]string{"店铺和卡券"},
		Summary:"核销卡券",
	})).POST("/consume_card_code",func(c *kelly.Context) {
		var form mp.ConsumeCardCodeParam
		if !param_process(c,&form){
			return
		}
		resp,err := wechatApi.ConsumeCardCode(&form)
		end_process(c,resp,err)
	})
}

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