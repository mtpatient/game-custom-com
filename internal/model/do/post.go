// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Post is the golang structure of table post for DAO operations like Where/Data.
type Post struct {
	g.Meta       `orm:"table:post, do:true"`
	Id           interface{} //
	UserId       interface{} // 用户id
	Title        interface{} // 标题
	Image        interface{} // 封面图片
	Content      interface{} // 帖子内容
	SectionId    interface{} // 所属板块
	ViewCount    interface{} // 浏览数
	LikeCount    interface{} // 点赞数
	CollectCount interface{} // 被收藏数
	Status       interface{} // 0：正常，1：禁用，2：申请恢复
	CreateTime   *gtime.Time //
	UpdateTime   *gtime.Time //
	DeleteTime   *gtime.Time //
}
