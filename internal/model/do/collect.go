// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Collect is the golang structure of table collect for DAO operations like Where/Data.
type Collect struct {
	g.Meta     `orm:"table:collect, do:true"`
	Id         interface{} //
	UserId     interface{} //
	PostId     interface{} //
	CreateTime *gtime.Time //
	UpdateTime *gtime.Time //
}
