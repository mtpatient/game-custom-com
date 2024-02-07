// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AdminLog is the golang structure for table admin_log.
type AdminLog struct {
	Id         uint        `json:"id"          ` //
	UserId     int         `json:"user_id"     ` //
	Type       int         `json:"type"        ` //
	Content    string      `json:"content"     ` //
	CreateTime *gtime.Time `json:"create_time" ` //
	DeleteTime *gtime.Time `json:"delete_time" ` //
}
