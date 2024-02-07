// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SearchHistory is the golang structure of table search_history for DAO operations like Where/Data.
type SearchHistory struct {
	g.Meta     `orm:"table:search_history, do:true"`
	Id         interface{} //
	UserId     interface{} //
	Keyword    interface{} //
	CreateTime *gtime.Time //
	DeleteTime *gtime.Time //
}
