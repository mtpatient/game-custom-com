// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Section is the golang structure for table section.
type Section struct {
	Id         uint        `json:"id"          ` //
	Name       string      `json:"name"        ` //
	Ico        string      `json:"ico"         ` // url链接
	CreateTime *gtime.Time `json:"create_time" ` //
	UpdateTime *gtime.Time `json:"update_time" ` //
	DeleteTime *gtime.Time `json:"delete_time" ` //
}
