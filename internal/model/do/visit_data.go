// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// VisitData is the golang structure of table visit_data for DAO operations like Where/Data.
type VisitData struct {
	g.Meta     `orm:"table:visit_data, do:true"`
	Id         interface{} //
	Date       *gtime.Time //
	ViewCount  interface{} // 每日访客数
	CreateTime *gtime.Time //
	UpdateTime *gtime.Time //
	DeleteTime *gtime.Time //
}
