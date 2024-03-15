// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Message is the golang structure of table message for DAO operations like Where/Data.
type Message struct {
	g.Meta     `orm:"table:message, do:true"`
	Id         interface{} //
	UserId     interface{} //
	ReceiveId  interface{} // 接收消息的用户;为空的话则为管理员向全体用户发布的通知
	Type       interface{} // 0：网站通知；1：回复我的；2：给我点赞的；
	Content    interface{} // 消息内容
	CommentId  interface{} //
	PostId     interface{} //
	IsRead     interface{} // 0：未读；1：已读
	CreateTime *gtime.Time //
	UpdateTime *gtime.Time //
	DeleteTime *gtime.Time //
}
