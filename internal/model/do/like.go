// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Like is the golang structure of table like for DAO operations like Where/Data.
type Like struct {
	g.Meta     `orm:"table:like, do:true"`
	Id         interface{} //
	UserId     interface{} //
	PostId     interface{} //
	CommentId  interface{} //
	PraiseId   interface{} // 被赞人
	CreateTime *gtime.Time //
	DeleteTime *gtime.Time //
}
