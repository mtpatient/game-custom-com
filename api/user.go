package api

import "github.com/gogf/gf/v2/os/gtime"

type User struct {
	Username string `json:"username" v:"required"   ` // 用户名
	Password string `json:"password" v:"required"   ` // 密码
	Repwd    string `json:"repwd"`                    // 重复输入密码
}

type UserRes struct {
	Id          int         `json:"id"           ` // 用户id，唯一标识
	Username    string      `json:"username"     ` // 用户名
	Email       string      `json:"email"        ` // 邮箱，可通过邮箱找回密码
	Avatar      string      `json:"avatar"       ` // 头像id
	Sex         int         `json:"sex"          ` // 0：女，1：男；2：保密
	Signature   string      `json:"signature"    ` // 个性签名
	Role        int         `json:"role"         ` // 管理员：1，,普通用户：0
	Status      int         `json:"status"       ` // 用户所处状态，0为正常，1为被封禁
	FansCount   int         `json:"fans_count"   ` //
	LikeCount   int         `json:"like_count"   ` //
	FollowCount int         `json:"follow_count" ` //
	CreateTime  *gtime.Time `json:"create_time"  ` //
	UpdateTime  *gtime.Time `json:"update_time"  ` //
	DeleteTime  *gtime.Time `json:"delete_time"  ` //
}

type UserReplacePassword struct {
	Id         int    `json:"user_id"`
	CurPwd     string `json:"cur_pwd" v:"required"`
	NewPwd     string `json:"new_pwd" v:"required"`
	ConfirmPwd string `json:"confirm_pwd" v:"required"`
}
