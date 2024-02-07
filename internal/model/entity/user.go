// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure for table user.
type User struct {
	Id         uint        `json:"id"          ` // 用户id，唯一标识
	Username   string      `json:"username"    ` // 用户名
	Password   string      `json:"password"    ` // 密码
	Email      string      `json:"email"       ` // 邮箱，可通过邮箱找回密码
	Img        uint        `json:"img"         ` // 头像id
	Sex        int         `json:"sex"         ` // 0：女，1：男
	Signature  string      `json:"signature"   ` // 个性签名
	Role       uint        `json:"role"        ` // 管理员：1，,普通用户：0
	Status     uint        `json:"status"      ` // 用户所处状态，0为正常，1为被封禁
	CreateTime *gtime.Time `json:"create_time" ` //
	UpdateTime *gtime.Time `json:"update_time" ` //
	DeleteTime *gtime.Time `json:"delete_time" ` //
}
