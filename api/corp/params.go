package corp

import "github.com/qjw/go-wx-sdk/corp"

type GetDepartmentsParam struct {
	ParentID *int64 `json:"parent,omitempty" doc:"父部门id，可以为空"`
}

type DeleteDepartmentParam struct {
	ID int64 `json:"id" doc="部门id"`
}

type GetUserParam struct {
	ID string `json:"id" doc="userid"`
}

type DeleteUserParam struct {
	ID string `json:"id" doc="userid"`
}

type SendKfTextMsgParam struct {
	Sender   corp.KfMsgUserObj `json:"sender"`
	Receiver corp.KfMsgUserObj `json:"receiver"`
	Content  string            `json:"content"`
}

type SendKfImageMsgParam struct {
	Sender   corp.KfMsgUserObj `json:"sender"`
	Receiver corp.KfMsgUserObj `json:"receiver"`
	MediaID  string            `json:"mediaid"`
}
type SendKfFileMsgParam SendKfImageMsgParam
type SendKfVoiceMsgParam SendKfImageMsgParam

type GetKfsParam struct {
	Type *string `json:"type,omitempty" doc:"internal/external,不填时，同时返回内部、外部客服列表"`
}

type GetTmpMediaParam struct {
	MediaID string `json:"media_id"`
}

type GetAgentDetailParam struct {
	ID int64 `json:"id" doc="企业号应用id"`
}
type GetMenuParam GetAgentDetailParam
type DeleteMenuParam GetAgentDetailParam
type CreateMenuParam GetAgentDetailParam

type MsgHeaderParam struct {
	AgentID  int64    `json:"id" doc="企业号应用id"`
	ToUsers  []string `json:"tousers"`
	ToPartys []int64  `json:"toparties"`
	ToTag    []int64  `json:"totags"`
}

type MsgTextParam struct {
	MsgHeaderParam
	Content string `json:"content"`
}

type MsgImageParam struct {
	MsgHeaderParam
	MediaID string `json:"media_id"`
}
type MsgVoiceParam MsgImageParam
type MsgFileParam MsgImageParam

type MsgVideoParam struct {
	MsgHeaderParam
	MediaID     string `json:"media_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type MsgNewsParam struct {
	MsgHeaderParam
	News corp.NewsMsgObj `json:"news"`
}

type TagIDParam struct {
	TagID int64 `json:"tagid"`
}

type SignJsTicketParam struct {
	NonceStr  string `json:"nonceStr" form:"nonceStr"`
	Timestamp string `json:"timestamp" form:"timestamp"`
	Url       string `json:"url" form:"url" doc:"不要url#后面的"`
}

type AuthorizeBaseParam struct {
	Code  string `form:"code" json:"code"`
	State string `form:"state" json:"state"`
}

type ToUserIDParam struct {
	OpenID string `json:"openid"`
}
