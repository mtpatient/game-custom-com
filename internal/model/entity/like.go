// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Like is the golang structure for table like.
type Like struct {
	Id         int         `json:"id"          ` //
	UserId     int         `json:"user_id"     ` //
	PostId     int         `json:"post_id"     ` //
	CommentId  int         `json:"comment_id"  ` //
	PraiseId   int         `json:"praise_id"   ` // 被赞人
	CreateTime *gtime.Time `json:"create_time" ` //
	UpdateTime *gtime.Time `json:"update_time" ` //
	DeleteTime *gtime.Time `json:"delete_time" ` //
}
