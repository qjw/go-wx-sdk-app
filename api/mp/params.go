package mp

import "github.com/qjw/go-wx-sdk/mp"

type SendKfMsgParam struct {
	ToUser  string `form:"touser" json:"touser"`
	Content string `form:"content" json:"content"`
}

type SendKfImageMsgParam struct {
	ToUser  string `form:"touser" json:"touser"`
	MediaID string `form:"media_id" json:"media_id"`
}

type SendKfCardMsgParam struct {
	ToUser string `form:"touser" json:"touser"`
	CardID string `form:"card_id" json:"card_id"`
}

type SendKfVideoMsgParam struct {
	ToUser       string `form:"touser" json:"touser"`
	MediaID      string `form:"media_id" json:"media_id"`
	ThumbMediaID string `form:"media_id" json:"thumb_media_id"`
	Name         string `form:"name" json:"name,omitempty"`
	Description  string `form:"description" json:"description,omitempty"`
}

type SendKfArticleMsgParam struct {
	ToUser   string       `form:"touser" json:"touser"`
	Articles []mp.Article `form:"articles" json:"articles"`
}

type AddKfAccountParam struct {
	KfAccount string `form:"kf_account" json:"kf_account" doc:"不是邮箱帐号，@后面是公众号的微信号"`
	Nickname  string `form:"nickname" json:"nickname"`
	Password  string `form:"password" json:"password"`
}

type UpdateKfAccountParam AddKfAccountParam

type DelKfAccountParam struct {
	KfAccount string `form:"kf_account" json:"kf_account" doc:"不是邮箱帐号，@后面是公众号的微信号"`
}

type GetUserDetailParam struct {
	OpenID string `form:"openid" json:"openid"`
}

type UpdateTagParam struct {
	ID   int    `form:"id" json:"id"`
	Name string `form:"name" json:"name"`
}

type CreateTagParam struct {
	Name string `form:"name" json:"name"`
}

type DeleteTagParam struct {
	ID int `form:"id" json:"id"`
}

type TagUsers DeleteTagParam

type TagAddMember struct {
	TagID      int      `form:"tagid" json:"tagid"`
	OpenIDList []string `form:"openid_list" json:"openid_list"`
}
type TagRmMember TagAddMember

type GetUserTagsParam GetUserDetailParam

type GetMaterialsParam struct {
	Tp     string `form:"type" json:"type" doc:"image/video/voice"`
	Offset int    `form:"offset" json:"offset"`
	Count  int    `form:"count" json:"count"`
}
type GetNewsMaterialsParam struct {
	Offset int `form:"offset" json:"offset"`
	Count  int `form:"count" json:"count"`
}

type UserRemarkParam struct {
	OpenID string `form:"openid" json:"openid"`
	Remark string `form:"remark" json:"remark"`
}

type GetBlackListParam struct {
	NextOpenid string `form:"next_openid" json:"next_openid"`
}

type BatchUnblacklist struct {
	OpenIDList []string `form:"openid_list" json:"openid_list"`
}
type BatchBlacklist BatchUnblacklist

type AddTemplateParam struct {
	TemplateIDShort string `form:"template_id_short" json:"template_id_short"`
}
type DeleteTemplateParam struct {
	TemplateID string `form:"template_id" json:"template_id"`
}

type LimitStrQrcodeParam struct {
	SceneStr string `form:"scene_str" json:"scene_str"`
}
type LimitQrcodeParam struct {
	SceneID int `form:"scene_id" json:"scene_id"`
}
type QrcodeParam struct {
	SceneID       int `form:"scene_id" json:"scene_id"`
	ExpireSeconds int `form:"expire_seconds" json:"expire_seconds"`
}

type ShowQrcodeParam struct {
	Ticket string `form:"ticket" json:"ticket"`
}

type ShortUrlParam struct {
	LongUrl string `form:"long_url" json:"long_url"`
}

type BatchUserDetail BatchUnblacklist

type CreateMenuParam struct {
	Menus []mp.MenuEntry `json:"menus" form:"menus"`
}

type SignJsTicketParam struct {
	NonceStr  string `json:"nonceStr" form:"nonceStr"`
	Timestamp string `json:"timestamp" form:"timestamp"`
	Url       string `json:"url" form:"url" doc:"不要url#后面的"`
}

type GetTmpMaterialParam struct {
	MediaID string `form:"media_id" json:"media_id"`
}

type GetMaterialParam GetTmpMaterialParam
type DeleteMaterialParam GetTmpMaterialParam

type MassPreviewTextParam struct {
	OpenID  string `form:"openid" json:"openid"`
	Content string `form:"content" json:"content"`
}

type MassPreviewParam struct {
	OpenID  string `form:"openid" json:"openid"`
	MediaID string `form:"media_id" json:"media_id"`
	Tp      string `form:"type" json:"type" doc:"mpnews/voice/image/mpvideo"`
}

type MassSendMpnewsParam struct {
	OpenIDList        []string `form:"openids" json:"openids"`
	MediaID           string   `form:"media_id" json:"media_id"`
	SendIgnoreReprint int      `form:"send_ignore_reprint" json:"send_ignore_reprint"`
}

type MassSendTextParam struct {
	OpenIDList []string `form:"openids" json:"openids"`
	Content    string   `form:"content" json:"content"`
}

type MassSendCardParam struct {
	OpenIDList []string `form:"openids" json:"openids"`
	CardID     string   `form:"card_id" json:"card_id"`
}

type MassSendMsgParam struct {
	OpenIDList []string `form:"openids" json:"openids"`
	MediaID    string   `form:"media_id" json:"media_id"`
	Tp         string   `form:"type" json:"type" doc:"voice/image"`
}

type MassGetParam struct {
	MsgID int64 `form:"msg_id" json:"msg_id"`
}

type MassDeleteParam MassGetParam

type MassSendAllMpnewsParam struct {
	IsToall           bool   `form:"is_to_all" json:"is_to_all"`
	TagID             int    `form:"tag_id" json:"tag_id"`
	MediaID           string `form:"media_id" json:"media_id"`
	SendIgnoreReprint int    `form:"send_ignore_reprint" json:"send_ignore_reprint"`
}

type MassSendAllTextParam struct {
	IsToall bool   `form:"is_to_all" json:"is_to_all"`
	TagID   int    `form:"tag_id" json:"tag_id"`
	Content string `form:"content" json:"content"`
}

type MassSendAllCardParam struct {
	IsToall bool   `form:"is_to_all" json:"is_to_all"`
	TagID   int    `form:"tag_id" json:"tag_id"`
	CardID  string `form:"card_id" json:"card_id"`
}

type MassSendAllMsgParam struct {
	IsToall bool   `form:"is_to_all" json:"is_to_all"`
	TagID   int    `form:"tag_id" json:"tag_id"`
	MediaID string `form:"media_id" json:"media_id"`
	Tp      string `form:"type" json:"type" doc:"voice/image"`
}

type AuthorizeBaseParam struct {
	Code  string `form:"code" json:"code"`
	State string `form:"state" json:"state"`
}

type SignChooseCardParam struct {
	ShopID   string `json:"shop_id,omitempty"`
	CardID   string `json:"card_id,omitempty"`
	CardType string `json:"card_type,omitempty"`
}
