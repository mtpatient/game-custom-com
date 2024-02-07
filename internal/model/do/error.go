// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Error is the golang structure of table error for DAO operations like Where/Data.
type Error struct {
	g.Meta     `orm:"table:error, do:true"`
	Id         interface{} //
	Log        interface{} //
	Type       interface{} //
	CreateTime *gtime.Time //
	DeleteTime *gtime.Time //
}
