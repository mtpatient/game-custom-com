// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Feedback is the golang structure of table feedback for DAO operations like Where/Data.
type Feedback struct {
	g.Meta     `orm:"table:feedback, do:true"`
	Id         interface{} //
	UserId     interface{} //
	Comment    interface{} //
	CreateTime *gtime.Time //
	DeleteTime *gtime.Time //
}
