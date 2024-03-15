// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Section is the golang structure for table section.
type Section struct {
	Id         int         `json:"id"          ` //
	Name       string      `json:"name"        ` //
	Dc         string      `json:"dc"          ` // 描述
	Role       int         `json:"role"        ` // 该板块所属用户，0：普通用户，1：仅管理员用户
	Icon       string      `json:"icon"        ` // url链接
	CreateTime *gtime.Time `json:"create_time" ` //
	UpdateTime *gtime.Time `json:"update_time" ` //
	DeleteTime *gtime.Time `json:"delete_time" ` //
}
