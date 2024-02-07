// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// FriendLink is the golang structure of table friend_link for DAO operations like Where/Data.
type FriendLink struct {
	g.Meta     `orm:"table:friend_link, do:true"`
	Id         interface{} //
	Name       interface{} // 友链名称
	Url        interface{} // 友链地址
	CreateTime *gtime.Time // 创建时间
	UpdateTime *gtime.Time //
	DeleteTime *gtime.Time //
}
