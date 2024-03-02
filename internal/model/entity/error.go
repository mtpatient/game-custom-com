// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Error is the golang structure for table error.
type Error struct {
	Id         int         `json:"id"          ` //
	Log        string      `json:"log"         ` //
	Type       int         `json:"type"        ` //
	CreateTime *gtime.Time `json:"create_time" ` //
	DeleteTime *gtime.Time `json:"delete_time" ` //
}
