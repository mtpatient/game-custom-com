// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// VisitData is the golang structure for table visit_data.
type VisitData struct {
	Id         int         `json:"id"          ` //
	Date       *gtime.Time `json:"date"        ` //
	ViewCount  int         `json:"view_count"  ` // 每日访客数
	CreateTime *gtime.Time `json:"create_time" ` //
	UpdateTime *gtime.Time `json:"update_time" ` //
	DeleteTime *gtime.Time `json:"delete_time" ` //
}
