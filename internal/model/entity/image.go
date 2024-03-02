// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Image is the golang structure for table image.
type Image struct {
	Id         int         `json:"id"          ` //
	Url        string      `json:"url"         ` // 图片地址
	Type       int         `json:"type"        ` // 0：用户头像；1：帖子图片
	Name       string      `json:"name"        ` //
	PostId     int         `json:"post_id"     ` // 帖子图片
	CreateTime *gtime.Time `json:"create_time" ` //
	UpdateTime *gtime.Time `json:"update_time" ` //
	DeleteTime *gtime.Time `json:"delete_time" ` //
}
