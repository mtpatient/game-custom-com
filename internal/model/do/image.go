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
	Type       interface{} // 0：用户头像；1：网站背景图；2：轮播图；3：用户图床
	UserId     interface{} // 图床图片所属用户
	CreateTime *gtime.Time //
	UpdateTime *gtime.Time //
	DeleteTime *gtime.Time //
}
