// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure of table user for DAO operations like Where/Data.
type User struct {
	g.Meta      `orm:"table:user, do:true"`
	Id          interface{} // 用户id，唯一标识
	Username    interface{} // 用户名
	Password    interface{} // 密码
	Email       interface{} // 邮箱，可通过邮箱找回密码
	Avatar      interface{} // 头像id
	Sex         interface{} // 2：女，1：男；3：保密
	Signature   interface{} // 个性签名
	Role        interface{} // 管理员：1，,普通用户：0
	Status      interface{} // 用户所处状态，0为正常，1为被封禁
	FansCount   interface{} //
	LikeCount   interface{} //
	FollowCount interface{} //
	CreateTime  *gtime.Time //
	UpdateTime  *gtime.Time //
	DeleteTime  *gtime.Time //
}
