// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Follow is the golang structure for table follow.
type Follow struct {
	Id           uint        `json:"id"             ` //
	UserId       uint        `json:"user_id"        ` //
	FollowUserId uint        `json:"follow_user_id" ` //
	CreateTime   *gtime.Time `json:"create_time"    ` //
	DeleteTime   *gtime.Time `json:"delete_time"    ` //
}
