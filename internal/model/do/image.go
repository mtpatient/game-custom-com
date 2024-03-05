// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Image is the golang structure of table image for DAO operations like Where/Data.
type Image struct {
	g.Meta     `orm:"table:image, do:true"`
	Id         interface{} //
	Url        interface{} // 图片地址
	Type       interface{} // 0：用户头像；1：帖子图片; 2：意见反馈
	Name       interface{} //
	PostId     interface{} // 帖子图片
	FeedbackId interface{} // 反馈图片
	CreateTime *gtime.Time //
	UpdateTime *gtime.Time //
	DeleteTime *gtime.Time //
}
