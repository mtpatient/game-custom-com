// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Comment is the golang structure of table comment for DAO operations like Where/Data.
type Comment struct {
	g.Meta     `orm:"table:comment, do:true"`
	Id         interface{} //
	PostId     interface{} //
	UserId     interface{} //
	ReplyId    interface{} // 被评论者
	Floor      interface{} // 评论从1开始；0为楼中楼
	ParentId   interface{} //
	Content    interface{} //
	LikeCount  interface{} //
	Status     interface{} // 0为正常；1为被删除
	CreateTime *gtime.Time //
	UpdateTime *gtime.Time //
	DeleteTime *gtime.Time //
}
