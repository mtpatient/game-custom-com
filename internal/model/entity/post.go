// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Post is the golang structure for table post.
type Post struct {
	Id           int         `json:"id"            ` //
	UserId       int         `json:"user_id"       ` // 用户id
	Title        string      `json:"title"         ` // 标题
	Content      string      `json:"content"       ` // 帖子内容
	Section      int         `json:"section"       ` // 所属板块
	ViewCount    int         `json:"view_count"    ` // 浏览数
	LikeCount    int         `json:"like_count"    ` // 点赞数
	CollectCount int         `json:"collect_count" ` // 被收藏数
	IsTop        int         `json:"is_top"        ` // 是否置顶
	Status       uint        `json:"status"        ` // 0：正常，1：禁用，2：仅自己可见，3：申请恢复
	CreateTime   *gtime.Time `json:"create_time"   ` //
	UpdateTime   *gtime.Time `json:"update_time"   ` //
	DeleteTime   *gtime.Time `json:"delete_time"   ` //
}
