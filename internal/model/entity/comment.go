// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Comment is the golang structure for table comment.
type Comment struct {
	Id         int         `json:"id"          ` //
	PostId     int         `json:"post_id"     ` //
	UserId     int         `json:"user_id"     ` //
	ReplyId    int         `json:"reply_id"    ` // 被评论者
	CommentId  int         `json:"comment_id"  ` //
	Floor      int         `json:"floor"       ` // 评论从1开始；0为楼中楼
	ParentId   int         `json:"parent_id"   ` //
	Content    string      `json:"content"     ` //
	LikeCount  int         `json:"like_count"  ` //
	Status     int         `json:"status"      ` // 0为正常；1为被删除
	CreateTime *gtime.Time `json:"create_time" ` //
	UpdateTime *gtime.Time `json:"update_time" ` //
	DeleteTime *gtime.Time `json:"delete_time" ` //
}
