// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Message is the golang structure for table message.
type Message struct {
	Id         int         `json:"id"          ` //
	UserId     int         `json:"user_id"     ` //
	ReciveId   int         `json:"recive_id"   ` // 接收消息的用户;为空的话则为管理员向全体用户发布的通知
	Type       int         `json:"type"        ` // 0：网站通知；1：回复我的；2：给我点赞的；3：@我的
	Content    string      `json:"content"     ` // 消息内容
	IsRead     int         `json:"is_read"     ` // 0：未读；1：已读
	CreateTime *gtime.Time `json:"create_time" ` //
	UpdateTime *gtime.Time `json:"update_time" ` //
	DeleteTime *gtime.Time `json:"delete_time" ` //
}
