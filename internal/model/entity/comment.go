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
	ReplayId   int         `json:"replay_id"   ` // 被评论者
	Floor      int         `json:"floor"       ` // 1为楼主；评论从2开始；0为楼中楼
	Content    string      `json:"content"     ` //
	LikeCount  int         `json:"like_count"  ` //
	Status     int         `json:"status"      ` // 0为正常；1为被删除
	CreateTime *gtime.Time `json:"create_time" ` //
	DeleteTime *gtime.Time `json:"delete_time" ` //
}
