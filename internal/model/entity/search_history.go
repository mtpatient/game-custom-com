// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SearchHistory is the golang structure for table search_history.
type SearchHistory struct {
	Id         int         `json:"id"          ` //
	UserId     int         `json:"user_id"     ` //
	Keyword    string      `json:"keyword"     ` //
	CreateTime *gtime.Time `json:"create_time" ` //
	DeleteTime *gtime.Time `json:"delete_time" ` //
}
