// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Collect is the golang structure for table collect.
type Collect struct {
	Id         int         `json:"id"          ` //
	UserId     int         `json:"user_id"     ` //
	PostId     int         `json:"post_id"     ` //
	CreateTime *gtime.Time `json:"create_time" ` //
	UpdateTime *gtime.Time `json:"update_time" ` //
}
