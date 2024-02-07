// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Collect is the golang structure for table collect.
type Collect struct {
	Id         uint        `json:"id"          ` //
	UserId     uint        `json:"user_id"     ` //
	PostId     uint        `json:"post_id"     ` //
	CreateTime *gtime.Time `json:"create_time" ` //
	UpdateTime *gtime.Time `json:"update_time" ` //
}
