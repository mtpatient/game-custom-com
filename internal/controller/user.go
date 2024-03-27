package controller

import (
	"fmt"
	"game-custom-com/api"
	"game-custom-com/internal/consts"
	"game-custom-com/internal/model/entity"
	"game-custom-com/internal/service"
	"game-custom-com/utility"
	"github.com/gogf/gf/v2/net/ghttp"
	"net/http"
)

type (
	User struct {
	}
)

func CUser() *User {
	return &User{}
}

func (c *User) GetUser(r *ghttp.Request) {
	id := r.Get("id")

	user, err := service.User().GetById(r.Context(), id.Int())
	if err != nil {
		r.Response.WriteJsonExit(utility.GetR().Error(consts.ServiceErrCode, err.Error()))
	}

	r.Response.WriteJsonExit(utility.GetR().PUT("user", user))
}

func (c *User) Login(r *ghttp.Request) {
	var u api.User

	res := r.Response
	err := r.Parse(&u)
	if err != nil {
		res.WriteJsonExit(utility.GetR().Error(consts.RequestErrCode, err.Error()))
	}

	user, token, err := service.User().Login(r.Context(), u)
	if err != nil {
		res.WriteJsonExit(utility.GetR().Error(consts.ServiceErrCode, err.Error()))
	}

	res.WriteJsonExit(utility.GetR().PUT("user", user).PUT("token", token))
}

func (c *User) Register(r *ghttp.Request) {
	var user api.User

	res := r.Response
	err := r.Parse(&user)
	if err != nil {
		res.WriteJsonExit(utility.GetR().Error(consts.RequestErrCode, err.Error()))
	}

	err = service.User().Register(r.Context(), user)
	if err != nil {
		res.WriteJsonExit(utility.GetR().Error(consts.ServiceErrCode, err.Error()))
	}

	res.WriteJsonExit(utility.GetR())
}

func (c *User) Logout(r *ghttp.Request) {
	token := r.Header["Token"]
	if token == nil {
		r.Response.WriteStatusExit(http.StatusForbidden)
	}

	service.User().Logout(r.Context(), token[0])

	r.Response.WriteJsonExit(utility.GetR())
}

func (c *User) NameExist(r *ghttp.Request) {
	username := r.Get("username")

	ok, err := service.User().NameExist(r.Context(), username.String())
	res := r.Response
	if err != nil {
		res.WriteJsonExit(utility.GetR().Error(consts.ServiceErrCode, err.Error()))
	}

	res.WriteJsonExit(utility.GetR().PUT("exist", ok))
}

func (c *User) IsLogin(r *ghttp.Request) {
	ok, err := service.User().IsLogin(r.Context())
	if err != nil {
		r.Response.WriteJsonExit(utility.GetR().Error(consts.ServiceErrCode, err.Error()))
	}
	r.Response.WriteJsonExit(utility.GetR().PUT("is_login", ok))
}

func (c *User) Update(r *ghttp.Request) {
	var user entity.User

	err := r.Parse(&user)
	if err != nil {
		r.Response.WriteJsonExit(utility.GetR().Error(consts.RequestErrCode, err.Error()))
	}
	err = service.User().Update(r.Context(), user)
	if err != nil {
		r.Response.WriteJsonExit(utility.GetR().Error(consts.ServiceErrCode, err.Error()))
	}

	r.Response.WriteJsonExit(utility.GetR())
}

func (c *User) ReplacePassword(r *ghttp.Request) {
	var rp api.UserReplacePassword
	err := r.Parse(&rp)
	if err != nil {
		r.Response.WriteJsonExit(utility.GetR().Error(consts.RequestErrCode, err.Error()))
	}
	err = service.User().ReplacePassword(r.Context(), rp)
	if err != nil {
		r.Response.WriteJsonExit(utility.GetR().Error(consts.ServiceErrCode, err.Error()))
	}

	r.Response.WriteJsonExit(utility.GetR())
}

func (c *User) GetAuthCode(r *ghttp.Request) {
	username := r.Get("username")
	fmt.Println(username.String())
	err := service.User().GetAuthCode(r.Context(), username.String())
	if err != nil {
		r.Response.WriteJsonExit(utility.GetR().Error(consts.ServiceErrCode, err.Error()))
	}

	r.Response.WriteJsonExit(utility.GetR())
}

func (c *User) ResetPwd(r *ghttp.Request) {
	var rs api.ResetPwd
	err := r.Parse(&rs)
	if err != nil {
		r.Response.WriteJsonExit(utility.GetR().Error(consts.RequestErrCode, err.Error()))
	}

	err = service.User().ResetPwd(r.Context(), rs)
	if err != nil {
		r.Response.WriteJsonExit(utility.GetR().Error(consts.ServiceErrCode, err.Error()))
	}

	r.Response.WriteJsonExit(utility.GetR())
}

func (c *User) Follow(r *ghttp.Request) {
	var follow api.UserFollow

	err := r.Parse(&follow)
	if err != nil {
		r.Response.WriteJsonExit(utility.GetR().Error(consts.RequestErrCode, err.Error()))
	}

	err = service.User().Follow(r.Context(), follow)
	if err != nil {
		r.Response.WriteJsonExit(utility.GetR().Error(consts.ServiceErrCode, err.Error()))
	}
	r.Response.WriteJsonExit(utility.GetR())
}

func (c *User) SearchUser(r *ghttp.Request) {
	var get api.UserSearchParams

	err := r.Parse(&get)
	if err != nil {
		r.Response.WriteJsonExit(utility.GetR().Error(consts.RequestErrCode, err.Error()))
	}

	users, err := service.User().SearchUser(r.Context(), get)
	if err != nil {
		r.Response.WriteJsonExit(utility.GetR().Error(consts.ServiceErrCode, err.Error()))
	}
	r.Response.WriteJsonExit(utility.GetR().PUT("users", users))
}

// GetUserList @router /user/list [post]
// @Summary 获取用户列表
// @Description 获取用户列表
// @Tags user
// @accept application/json
// @Param body api.CommonParams
// @Success utility.R{code=0,msg=””,data{users=[]entity.User,total=int}}
func (c *User) GetUserList(r *ghttp.Request) {
	var get api.CommonParams
	err := r.Parse(&get)
	if err != nil {
		r.Response.WriteJsonExit(utility.GetR().Error(consts.RequestErrCode, err.Error()))
	}
	users, total, err := service.User().GetUserList(r.Context(), get)
	if err != nil {
		r.Response.WriteJsonExit(utility.GetR().Error(consts.ServiceErrCode, err.Error()))
	}

	r.Response.WriteJsonExit(utility.GetR().PUT("users", users).PUT("total", total))
}

// Ban @router /user/ban [post]
// @Summary 获取用户列表
// @Description 获取用户列表
// @Tags user
// @accept application/json
// @Param body api.Ban
// @Success utility.R{code=0,msg=””,data{}}
func (c *User) Ban(r *ghttp.Request) {
	var ban api.Ban
	err := r.Parse(&ban)
	if err != nil {
		r.Response.WriteJsonExit(utility.GetR().Error(consts.RequestErrCode, err.Error()))
	}
	err = service.User().Ban(r.Context(), ban)
	if err != nil {
		r.Response.WriteJsonExit(utility.GetR().Error(consts.ServiceErrCode, err.Error()))
	}
	r.Response.WriteJsonExit(utility.GetR())
}
