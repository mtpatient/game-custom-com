// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Post is the golang structure for table post.
type Post struct {
	Id           uint        `json:"id"            ` //
	UserId       int         `json:"user_id"       ` // 用户id
	Title        string      `json:"title"         ` // 标题
	Image        uint        `json:"image"         ` // 封面图片
	Content      string      `json:"content"       ` // 帖子内容
	SectionId    int         `json:"section_id"    ` // 所属板块
	ViewCount    uint        `json:"view_count"    ` // 浏览数
	LikeCount    uint        `json:"like_count"    ` // 点赞数
	CollectCount uint        `json:"collect_count" ` // 被收藏数
	Status       uint        `json:"status"        ` // 0：正常，1：禁用，2：申请恢复
	CreateTime   *gtime.Time `json:"create_time"   ` //
	UpdateTime   *gtime.Time `json:"update_time"   ` //
	DeleteTime   *gtime.Time `json:"delete_time"   ` //
}
