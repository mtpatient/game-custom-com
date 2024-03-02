// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Feedback is the golang structure for table feedback.
type Feedback struct {
	Id         int         `json:"id"          ` //
	UserId     int         `json:"user_id"     ` //
	Comment    string      `json:"comment"     ` //
	CreateTime *gtime.Time `json:"create_time" ` //
	DeleteTime *gtime.Time `json:"delete_time" ` //
}
