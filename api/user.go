package api

type User struct {
	Username string `json:"username" v:"required"   ` // 用户名
	Password string `json:"password" v:"required"   ` // 密码
	Repwd    string `json:"repwd"`                    // 重复输入密码
}
