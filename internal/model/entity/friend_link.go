// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// FriendLink is the golang structure for table friend_link.
type FriendLink struct {
	Id         uint        `json:"id"          ` //
	Name       string      `json:"name"        ` // 友链名称
	Url        string      `json:"url"         ` // 友链地址
	CreateTime *gtime.Time `json:"create_time" ` // 创建时间
	UpdateTime *gtime.Time `json:"update_time" ` //
	DeleteTime *gtime.Time `json:"delete_time" ` //
}
