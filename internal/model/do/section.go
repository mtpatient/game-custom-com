// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Section is the golang structure of table section for DAO operations like Where/Data.
type Section struct {
	g.Meta     `orm:"table:section, do:true"`
	Id         interface{} //
	Name       interface{} //
	Dc         interface{} // 描述
	Role       interface{} // 该板块所属用户，0：普通用户，1：仅管理员用户
	Icon       interface{} // url链接
	CreateTime *gtime.Time //
	UpdateTime *gtime.Time //
	DeleteTime *gtime.Time //
}
