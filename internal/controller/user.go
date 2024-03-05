package controller

import (
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

	img, _ := service.Img().GetImageById(r.Context(), user.Avatar)

	r.Response.WriteJsonExit(utility.GetR().PUT("user", &api.UserRes{
		Id:          user.Id,
		Username:    user.Username,
		Email:       user.Email,
		Avatar:      img.Url,
		Sex:         user.Sex,
		Signature:   user.Signature,
		Role:        user.Role,
		Status:      user.Status,
		FansCount:   user.FansCount,
		LikeCount:   user.LikeCount,
		FollowCount: user.FollowCount,
		CreateTime:  user.CreateTime,
	}))
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

	img, _ := service.Img().GetImageById(r.Context(), user.Avatar)

	res.WriteJsonExit(utility.GetR().PUT("user", &api.UserRes{
		Id:          user.Id,
		Username:    user.Username,
		Email:       user.Email,
		Avatar:      img.Url,
		Sex:         user.Sex,
		Signature:   user.Signature,
		Role:        user.Role,
		Status:      user.Status,
		FansCount:   user.FansCount,
		LikeCount:   user.LikeCount,
		FollowCount: user.FollowCount,
		CreateTime:  user.CreateTime,
	}).PUT("token", token))
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
