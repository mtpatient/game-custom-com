// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Follow is the golang structure of table follow for DAO operations like Where/Data.
type Follow struct {
	g.Meta       `orm:"table:follow, do:true"`
	Id           interface{} //
	UserId       interface{} //
	FollowUserId interface{} //
	CreateTime   *gtime.Time //
	DeleteTime   *gtime.Time //
}
