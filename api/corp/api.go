package corp

import (
	"strings"
	"net/http"
	"strconv"
	"io"
	"github.com/qjw/go-wx-sdk/corp"
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

func InitializeApiRoutes(grouter kelly.Router, context *corp.Context, agentConfig *corp.AgentConfig) {
	urouter := grouter.Group("/corp")
	wechatApi := corp.NewCorpApi(context)
	//------------------------------------------------------------------------
	urouter.Annotation(swagger.Swagger(&swagger.StructParam{
		ResponseData: &swagger.SuccessResp{},
		Tags:         []string{"企业号-杂项"},
		Summary:      "获取Access Token",
	})).GET("/token", func(c *kelly.Context) {
		resp, err := wechatApi.GetAccessToken()
		end_process(c, resp, err)
	})

	urouter.Annotation(swagger.Swagger(&swagger.StructParam{
		JsonData:&SignJsTicketParam{},
		ResponseData:&swagger.SuccessResp{},
		Tags:[]string{"企业号-杂项"},
		Summary:"Js签名",
	})).POST("/sign_js_ticket",func(c *kelly.Context) {
		var form SignJsTicketParam
		if !param_process(c,&form){
			return
		}

		resp,err := wechatApi.SignJsTicket(form.NonceStr,form.Timestamp,form.Url)
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
		resp,err := wechatApi.Oauth2GetUserInfo(form.Code)
		end_process(c,resp,err)
	})

	urouter.GET("/snsapi_userinfo",func(c *kelly.Context) {
		url := const_var.HostUrl + realPath(urouter,"snsapi_userinfo_callback")
		url1 := wechatApi.AuthorizeUserinfo(url,"snsapi_userinfo_test", agentConfig.AgentID)
		c.Redirect(http.StatusFound, url1)
	})

	urouter.GET("/snsapi_privateinfo",func(c *kelly.Context) {
		url := const_var.HostUrl + realPath(urouter,"snsapi_userinfo_callback")
		url1 := wechatApi.AuthorizePrivateInfo(url,"snsapi_privateinfo_test", agentConfig.AgentID)
		c.Redirect(http.StatusFound, url1)
	})

	urouter.GET("/snsapi_userinfo_callback",func(c *kelly.Context) {
		var form AuthorizeBaseParam
		if !param_process(c,&form){
			return
		}
		resp,err := wechatApi.Oauth2GetUserInfo(form.Code)
		if err != nil{
			c.WriteIndentedJson(http.StatusOK, kelly.H{
				"message": err.Error(),
				"result": http.StatusBadRequest,
			})
			return
		}

		resp2,err := wechatApi.Oauth2GetUserDetail(resp.UserTicket)
		end_process(c,resp2,err)
	})

	urouter.Annotation(swagger.Swagger(&swagger.StructParam{
		ResponseData:&swagger.SuccessResp{},
		Tags:[]string{"企业号-杂项"},
		Summary:"获取Js Ticket",
	})).GET("/jsticket",func(c *kelly.Context) {
		resp,err := wechatApi.GetJsTicket()
		end_process(c,resp,err)
	})

	urouter.Annotation(swagger.Swagger(&swagger.StructParam{
		ResponseData:&swagger.SuccessResp{},
		Tags:[]string{"企业号-杂项"},
		Summary:"获取微信服务器IP地址",
	})).GET("/iplist",func(c *kelly.Context) {
		resp,err := wechatApi.GetIpList()
		end_process(c,resp,err)
	})

	urouter.Annotation(swagger.Swagger(&swagger.StructParam{
		FormData:&corp.Convert2OpenIDObj{},
		ResponseData:&swagger.SuccessResp{},
		Tags:[]string{"企业号-杂项"},
		Summary:"userid转openid",
	})).POST("/convert_to_openid",func(c *kelly.Context) {
		var form corp.Convert2OpenIDObj
		if !param_process(c,&form){
			return
		}

		resp,err := wechatApi.Convert2OpenID(&form)
		end_process(c,resp,err)
	})

	urouter.Annotation(swagger.Swagger(&swagger.StructParam{
		FormData:&ToUserIDParam{},
		ResponseData:&swagger.SuccessResp{},
		Tags:[]string{"企业号-杂项"},
		Summary:"openid转userid",
	})).POST("/convert_to_userid",func(c *kelly.Context) {
		var form ToUserIDParam
		if !param_process(c,&form){
			return
		}

		resp,err := wechatApi.Convert2UserID(form.OpenID)
		end_process(c,resp,err)
	})
	//===============================部门============================================================================
	urouter.Annotation(swagger.Swagger(&swagger.StructParam{
		QueryData:&GetDepartmentsParam{},
		ResponseData:&swagger.SuccessResp{},
		Tags:[]string{"企业号-部门和用户管理"},
		Summary:"获取所有部门列表",
	})).GET("/get_departments",func(c *kelly.Context) {
		var form GetDepartmentsParam
		if !param_process(c,&form){
			return
		}
		resp,err := wechatApi.GetDepartments(form.ParentID)
		end_process(c,resp,err)
	})

	urouter.Annotation(swagger.Swagger(&swagger.StructParam{
		FormData:&corp.DepartmentObj{},
		ResponseData:&swagger.SuccessResp{},
		Tags:[]string{"企业号-部门和用户管理"},
		Summary:"创建部门",
	})).POST("/create_department",func(c *kelly.Context) {
		var form corp.DepartmentObj
		if !param_process(c,&form){
			return
		}
		resp,err := wechatApi.CreateDepartment(&form)
		end_process(c,resp,err)
	})

	urouter.Annotation(swagger.Swagger(&swagger.StructParam{
		FormData:&corp.DepartmentUpdateObj{},
		ResponseData:&swagger.SuccessResp{},
		Tags:[]string{"企业号-部门和用户管理"},
		Summary:"更新部门",
	})).PUT("/update_department",func(c *kelly.Context) {
		var form corp.DepartmentUpdateObj
		if !param_process(c,&form){
			return
		}
		resp,err := wechatApi.UpdateDepartment(&form)
		end_process(c,resp,err)
	})

	urouter.Annotation(swagger.Swagger(&swagger.StructParam{
		PathData:&DeleteDepartmentParam{},
		ResponseData:&swagger.SuccessResp{},
		Tags:[]string{"企业号-部门和用户管理"},
		Summary:"删除部门",
	})).DELETE("/delete_department/:id",func(c *kelly.Context) {
		var form DeleteDepartmentParam
		if !param_path_process(c,&form){
			return
		}
		resp,err := wechatApi.DeleteDepartments(form.ID)
		end_process(c,resp,err)
	})

	urouter.Annotation(swagger.Swagger(&swagger.StructParam{
		QueryData:&corp.SimpleUserlistObj{},
		ResponseData:&swagger.SuccessResp{},
		Tags:[]string{"企业号-部门和用户管理"},
		Summary:"获取部门成员（简单信息）",
	})).GET("/get_department_simple_users",func(c *kelly.Context) {
		var form corp.SimpleUserlistObj
		if !param_process(c,&form){
			return
		}
		resp,err := wechatApi.DepartmentSimpleUserlist(&form)
		end_process(c,resp,err)
	})

	urouter.Annotation(swagger.Swagger(&swagger.StructParam{
		QueryData:&corp.SimpleUserlistObj{},
		ResponseData:&swagger.SuccessResp{},
		Tags:[]string{"企业号-部门和用户管理"},
		Summary:"获取部门成员",
	})).GET("/get_department_users",func(c *kelly.Context) {
		var form corp.SimpleUserlistObj
		if !param_process(c,&form){
			return
		}
		resp,err := wechatApi.DepartmentUserlist(&form)
		end_process(c,resp,err)
	})

	urouter.Annotation(swagger.Swagger(&swagger.StructParam{
		QueryData:&GetUserParam{},
		ResponseData:&swagger.SuccessResp{},
		Tags:[]string{"企业号-部门和用户管理"},
		Summary:"获取用户详情",
	})).GET("/get_user",func(c *kelly.Context) {
		var form GetUserParam
		if !param_process(c,&form){
			return
		}
		resp,err := wechatApi.GetUser(form.ID)
		end_process(c,resp,err)
	})

	urouter.Annotation(swagger.Swagger(&swagger.StructParam{
		PathData:&DeleteDepartmentParam{},
		ResponseData:&swagger.SuccessResp{},
		Tags:[]string{"企业号-部门和用户管理"},
		Summary:"删除用户",
	})).DELETE("/delete_user/:id",func(c *kelly.Context) {
		var form DeleteUserParam
		if !param_path_process(c,&form){
			return
		}
		resp,err := wechatApi.DeleteUser(form.ID)
		end_process(c,resp,err)
	})

	urouter.Annotation(swagger.Swagger(&swagger.StructParam{
		JsonData:&corp.UserCreateObj{},
		ResponseData:&swagger.SuccessResp{},
		Tags:[]string{"企业号-部门和用户管理"},
		Summary:"创建用户",
	})).POST("/create_user",func(c *kelly.Context) {
		var form corp.UserCreateObj
		if !param_process(c,&form){
			return
		}
		resp,err := wechatApi.CreateUser(&form)
		end_process(c,resp,err)
	})

	urouter.Annotation(swagger.Swagger(&swagger.StructParam{
		JsonData:&corp.UserUpdateObj{},
		ResponseData:&swagger.SuccessResp{},
		Tags:[]string{"企业号-部门和用户管理"},
		Summary:"更新用户",
	})).PUT("/update_user",func(c *kelly.Context) {
		var form corp.UserUpdateObj
		if !param_process(c,&form){
			return
		}
		resp,err := wechatApi.UpdateUser(&form)
		end_process(c,resp,err)
	})

	urouter.Annotation(swagger.Swagger(&swagger.StructParam{
		JsonData:&corp.BatchDeleteUserObj{},
		ResponseData:&swagger.SuccessResp{},
		Tags:[]string{"企业号-部门和用户管理"},
		Summary:"批量删除用户",
	})).POST("/batch_delete_user",func(c *kelly.Context) {
		var form corp.BatchDeleteUserObj
		if !param_process(c,&form){
			return
		}
		resp,err := wechatApi.BatchDeleteUser(form.UserIDList)
		end_process(c,resp,err)
	})
	//===============================部门============================================================================
	//====================================Tag=======================================================================
	urouter.Annotation(swagger.Swagger(&swagger.StructParam{
		QueryData:&TagIDParam{},
		ResponseData:&swagger.SuccessResp{},
		Tags:[]string{"企业号-部门和用户管理"},
		Summary:"获取Tag用户",
	})).GET("/get_tag_users",func(c *kelly.Context) {
		var form TagIDParam
		if !param_process(c,&form){
			return
		}
		resp,err := wechatApi.GetTagUsers(form.TagID)
		end_process(c,resp,err)
	})

	urouter.Annotation(swagger.Swagger(&swagger.StructParam{
		ResponseData:&swagger.SuccessResp{},
		Tags:[]string{"企业号-部门和用户管理"},
		Summary:"获取Tag列表",
	})).GET("/get_tags",func(c *kelly.Context) {
		resp,err := wechatApi.GetTags()
		end_process(c,resp,err)
	})

	urouter.Annotation(swagger.Swagger(&swagger.StructParam{
		JsonData:&corp.TagUserUpdateObj{},
		ResponseData:&swagger.SuccessResp{},
		Tags:[]string{"企业号-部门和用户管理"},
		Summary:"增加tag用户",
	})).POST("/add_tag_users",func(c *kelly.Context) {
		var form corp.TagUserUpdateObj
		if !param_process(c,&form){
			return
		}
		resp,err := wechatApi.AddTagUsers(&form)
		end_process(c,resp,err)
	})

	urouter.Annotation(swagger.Swagger(&swagger.StructParam{
		JsonData:&corp.TagUserUpdateObj{},
		ResponseData:&swagger.SuccessResp{},
		Tags:[]string{"企业号-部门和用户管理"},
		Summary:"删除tag用户",
	})).POST("/del_tag_users",func(c *kelly.Context) {
		var form corp.TagUserUpdateObj
		if !param_process(c,&form){
			return
		}
		resp,err := wechatApi.DelTagUsers(&form)
		end_process(c,resp,err)
	})

	urouter.Annotation(swagger.Swagger(&swagger.StructParam{
		JsonData:&corp.TagCreateObj{},
		ResponseData:&swagger.SuccessResp{},
		Tags:[]string{"企业号-部门和用户管理"},
		Summary:"创建Tag",
	})).POST("/create_tag",func(c *kelly.Context) {
		var form corp.TagCreateObj
		if !param_process(c,&form){
			return
		}
		resp,err := wechatApi.CreateTag(&form)
		end_process(c,resp,err)
	})

	urouter.Annotation(swagger.Swagger(&swagger.StructParam{
		JsonData:&corp.TagObj{},
		ResponseData:&swagger.SuccessResp{},
		Tags:[]string{"企业号-部门和用户管理"},
		Summary:"更新Tag",
	})).POST("/update_tag",func(c *kelly.Context) {
		var form corp.TagObj
		if !param_process(c,&form){
			return
		}
		resp,err := wechatApi.UpdateTag(&form)
		end_process(c,resp,err)
	})

	urouter.Annotation(swagger.Swagger(&swagger.StructParam{
		PathData:&TagIDParam{},
		ResponseData:&swagger.SuccessResp{},
		Tags:[]string{"企业号-部门和用户管理"},
		Summary:"删除Tag",
	})).DELETE("/delete_tag/:tagid",func(c *kelly.Context) {
		var form TagIDParam
		if !param_path_process(c,&form){
			return
		}
		resp,err := wechatApi.DeleteTag(form.TagID)
		end_process(c,resp,err)
	})
	//====================================Tag=======================================================================
	//===============================KF=============================================================================
	urouter.Annotation(swagger.Swagger(&swagger.StructParam{
		JsonData:&SendKfTextMsgParam{},
		ResponseData:&swagger.SuccessResp{},
		Tags:[]string{"企业号-客服"},
		Summary:"发送文本消息",
	})).POST("/send_kf_text_msg",func(c *kelly.Context) {
		var form SendKfTextMsgParam
		if !param_process(c,&form){
			return
		}
		resp,err := wechatApi.SendKfText(&form.Sender,&form.Receiver,form.Content)
		end_process(c,resp,err)
	})

	urouter.Annotation(swagger.Swagger(&swagger.StructParam{
		JsonData:&SendKfImageMsgParam{},
		ResponseData:&swagger.SuccessResp{},
		Tags:[]string{"企业号-客服"},
		Summary:"发送图片消息",
	})).POST("/send_kf_image_msg",func(c *kelly.Context) {
		var form SendKfImageMsgParam
		if !param_process(c,&form){
			return
		}
		resp,err := wechatApi.SendKfImage(&form.Sender,&form.Receiver,form.MediaID)
		end_process(c,resp,err)
	})

	urouter.Annotation(swagger.Swagger(&swagger.StructParam{
		JsonData:&SendKfFileMsgParam{},
		ResponseData:&swagger.SuccessResp{},
		Tags:[]string{"企业号-客服"},
		Summary:"发送文件消息",
	})).POST("/send_kf_file_msg",func(c *kelly.Context) {
		var form SendKfFileMsgParam
		if !param_process(c,&form){
			return
		}
		resp,err := wechatApi.SendKfFile(&form.Sender,&form.Receiver,form.MediaID)
		end_process(c,resp,err)
	})

	urouter.Annotation(swagger.Swagger(&swagger.StructParam{
		JsonData:&SendKfVoiceMsgParam{},
		ResponseData:&swagger.SuccessResp{},
		Tags:[]string{"企业号-客服"},
		Summary:"发送语音消息",
	})).POST("/send_kf_voice_msg",func(c *kelly.Context) {
		var form SendKfVoiceMsgParam
		if !param_process(c,&form){
			return
		}
		resp,err := wechatApi.SendKfImage(&form.Sender,&form.Receiver,form.MediaID)
		end_process(c,resp,err)
	})

	urouter.Annotation(swagger.Swagger(&swagger.StructParam{
		QueryData:&GetKfsParam{},
		ResponseData:&swagger.SuccessResp{},
		Tags:[]string{"企业号-客服"},
		Summary:"获得所有客服",
	})).GET("/get_kfs",func(c *kelly.Context) {
		var form GetKfsParam
		if !param_process(c,&form){
			return
		}
		if form.Type == nil{
			emptyTp := ""
			form.Type = &emptyTp
		}
		resp,err := wechatApi.GetKfList(*form.Type)
		end_process(c,resp,err)
	})
	//===============================KF=============================================================================
	//===============================素材============================================================================
	urouter.Annotation(
		swagger.SwaggerFile("corp.yaml:upload_tmp_media"),
	).POST("/upload_tmp_media",func(c *kelly.Context) {

		file, multipart := c.MustGetFileVarible("file")
		defer file.Close()
		tp := c.MustGetFormVarible("type")
		resp,err := wechatApi.UploadTmpMedia(file,multipart.Filename,tp)
		end_process(c,resp,err)
	})

	urouter.Annotation(
		swagger.SwaggerFile("corp.yaml:get_tmp_media"),
	).GET("/get_tmp_media",func(c *kelly.Context) {
		var form GetTmpMediaParam
		if !param_process(c,&form){
			return
		}
		url,_ := wechatApi.GetTmpMedia(form.MediaID)
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
	//===============================素材============================================================================
	//===============================企业号应用=======================================================================
	urouter.Annotation(swagger.Swagger(&swagger.StructParam{
		ResponseData:&swagger.SuccessResp{},
		Tags:[]string{"企业号-应用"},
		Summary:"获取所有应用",
	})).GET("/get_agents",func(c *kelly.Context) {
		resp,err := wechatApi.GetAgentList()
		end_process(c,resp,err)
	})

	urouter.Annotation(swagger.Swagger(&swagger.StructParam{
		QueryData:&GetAgentDetailParam{},
		ResponseData:&swagger.SuccessResp{},
		Tags:[]string{"企业号-应用"},
		Summary:"获取应用详情",
	})).GET("/get_agent",func(c *kelly.Context) {
		var form GetAgentDetailParam
		if !param_process(c,&form){
			return
		}
		resp,err := wechatApi.GetAgentDetail(form.ID)
		end_process(c,resp,err)
	})

	urouter.Annotation(swagger.Swagger(&swagger.StructParam{
		JsonData:&corp.AgentUpdateObj{},
		ResponseData:&swagger.SuccessResp{},
		Tags:[]string{"企业号-应用"},
		Summary:"更新应用",
	})).PUT("/update_agent",func(c *kelly.Context) {
		var form corp.AgentUpdateObj
		if !param_process(c,&form){
			return
		}
		resp,err := wechatApi.UpdateAgent(&form)
		end_process(c,resp,err)
	})
	//===============================企业号应用=======================================================================
	//====================================菜单=======================================================================
	urouter.Annotation(swagger.Swagger(&swagger.StructParam{
		QueryData:&GetMenuParam{},
		ResponseData:&swagger.SuccessResp{},
		Tags:[]string{"企业号-菜单"},
		Summary:"获取应用菜单",
	})).GET("/get_menu",func(c *kelly.Context) {
		var form GetMenuParam
		if !param_process(c,&form){
			return
		}
		resp,err := wechatApi.GetMenu(form.ID)
		end_process(c,resp,err)
	})

	urouter.Annotation(swagger.Swagger(&swagger.StructParam{
		PathData:&CreateMenuParam{},
		JsonData:&corp.MenuCreateObj{},
		ResponseData:&swagger.SuccessResp{},
		Tags:[]string{"企业号-菜单"},
		Summary:"创建菜单",
	})).POST("/create_menu/:id",func(c *kelly.Context) {
		var form CreateMenuParam
		if !param_path_process(c,&form){
			return
		}
		var bodyForm corp.MenuCreateObj
		if !param_process(c,&bodyForm){
			return
		}

		resp,err := wechatApi.CreateMenu(form.ID,&bodyForm)
		end_process(c,resp,err)
	})

	urouter.Annotation(swagger.Swagger(&swagger.StructParam{
		PathData:&DeleteMenuParam{},
		ResponseData:&swagger.SuccessResp{},
		Tags:[]string{"企业号-菜单"},
		Summary:"删除应用菜单",
	})).DELETE("/delete_menu/:id",func(c *kelly.Context) {
		var form DeleteMenuParam
		if !param_path_process(c,&form){
			return
		}
		resp,err := wechatApi.DeleteMenu(form.ID)
		end_process(c,resp,err)
	})
	//====================================菜单=======================================================================
	//====================================消息=======================================================================
	urouter.Annotation(swagger.Swagger(&swagger.StructParam{
		JsonData:&MsgTextParam{},
		ResponseData:&swagger.SuccessResp{},
		Tags:[]string{"企业号-消息"},
		Summary:"发送文本消息",
	})).POST("/send_text_msg",func(c *kelly.Context) {
		var form MsgTextParam
		if !param_process(c,&form){
			return
		}

		resp,err := wechatApi.SendTextMsg(form.ToUsers,form.ToPartys,form.ToTag,form.AgentID,form.Content)
		end_process(c,resp,err)
	})

	urouter.Annotation(swagger.Swagger(&swagger.StructParam{
		JsonData:&MsgImageParam{},
		ResponseData:&swagger.SuccessResp{},
		Tags:[]string{"企业号-消息"},
		Summary:"发送图片消息",
	})).POST("/send_image_msg",func(c *kelly.Context) {
		var form MsgImageParam
		if !param_process(c,&form){
			return
		}

		resp,err := wechatApi.SendImageMsg(form.ToUsers,form.ToPartys,form.ToTag,form.AgentID,form.MediaID)
		end_process(c,resp,err)
	})

	urouter.Annotation(swagger.Swagger(&swagger.StructParam{
		JsonData:&MsgVoiceParam{},
		ResponseData:&swagger.SuccessResp{},
		Tags:[]string{"企业号-消息"},
		Summary:"发送语音消息",
	})).POST("/send_voice_msg",func(c *kelly.Context) {
		var form MsgVoiceParam
		if !param_process(c,&form){
			return
		}

		resp,err := wechatApi.SendVoiceMsg(form.ToUsers,form.ToPartys,form.ToTag,form.AgentID,form.MediaID)
		end_process(c,resp,err)
	})

	urouter.Annotation(swagger.Swagger(&swagger.StructParam{
		JsonData:&MsgFileParam{},
		ResponseData:&swagger.SuccessResp{},
		Tags:[]string{"企业号-消息"},
		Summary:"发送文件消息",
	})).POST("/send_file_msg",func(c *kelly.Context) {
		var form MsgFileParam
		if !param_process(c,&form){
			return
		}

		resp,err := wechatApi.SendFileMsg(form.ToUsers,form.ToPartys,form.ToTag,form.AgentID,form.MediaID)
		end_process(c,resp,err)
	})

	urouter.Annotation(swagger.Swagger(&swagger.StructParam{
		JsonData:&MsgVideoParam{},
		ResponseData:&swagger.SuccessResp{},
		Tags:[]string{"企业号-消息"},
		Summary:"发送视频消息",
	})).POST("/send_video_msg",func(c *kelly.Context) {
		var form MsgVideoParam
		if !param_process(c,&form){
			return
		}

		resp,err := wechatApi.SendVideoMsg(form.ToUsers,form.ToPartys,form.ToTag,form.AgentID,
			form.MediaID,form.Title,form.Description)
		end_process(c,resp,err)
	})

	urouter.Annotation(swagger.Swagger(&swagger.StructParam{
		JsonData:&MsgNewsParam{},
		ResponseData:&swagger.SuccessResp{},
		Tags:[]string{"企业号-消息"},
		Summary:"发送图文消息",
	})).POST("/send_news_msg",func(c *kelly.Context) {
		var form MsgNewsParam
		if !param_process(c,&form){
			return
		}

		resp,err := wechatApi.SendNewsMsg(form.ToUsers,form.ToPartys,form.ToTag,form.AgentID,&form.News)
		end_process(c,resp,err)
	})
	//====================================消息=======================================================================
	//===============================卡券============================================================================
	urouter.Annotation(
		swagger.SwaggerFile("corp.yaml:upload_card_logo"),
	).POST("/upload_card_logo",func(c *kelly.Context) {
		file, multipart := c.MustGetFileVarible("file")
		defer file.Close()
		resp,err := wechatApi.UploadCardLogo(file,multipart.Filename)
		end_process(c,resp,err)
	})

	urouter.Annotation(swagger.Swagger(&swagger.StructParam{
		JsonData:&corp.CouponCardCreateParam{},
		ResponseData:&swagger.SuccessResp{},
		Tags:[]string{"企业号-卡券"},
		Summary:"创建优惠券",
	})).POST("/create_coupon_card",func(c *kelly.Context) {
		var form corp.CouponCardCreateParam
		if !param_process(c,&form){
			return
		}

		resp,err := wechatApi.CreateCouponCard(&form)
		end_process(c,resp,err)
	})

	urouter.Annotation(swagger.Swagger(&swagger.StructParam{
		QueryData:&corp.CardListParam{},
		ResponseData:&corp.CardListRes{},
		Tags:[]string{"企业号-卡券"},
		Summary:"获取卡券列表",
	})).GET("/get_cards",func(c *kelly.Context) {
		var form corp.CardListParam
		if !param_process(c,&form){
			return
		}
		resp,err := wechatApi.GetCards(&form)
		end_process(c,resp,err)
	})
}