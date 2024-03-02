package controller

import (
	"game-custom-com/api"
	"game-custom-com/internal/consts"
	"game-custom-com/internal/service"
	"game-custom-com/utility"
	"github.com/gogf/gf/v2/net/ghttp"
)

type Post struct {
}

func (c *Post) CPost() *Post {
	return &Post{}
}

func (c *Post) Add(r *ghttp.Request) {
	var postAdd api.PostAdd

	err := r.Parse(&postAdd)
	res := r.Response
	if err != nil {
		res.WriteJsonExit(utility.GetR().Error(consts.RequestErrCode, err.Error()))
	}

	err = service.Post().Add(r.Context(), postAdd)
	if err != nil {
		res.WriteJsonExit(utility.GetR().Error(consts.ServiceErrCode, err.Error()))
	}

	res.WriteJsonExit(utility.GetR())
}
