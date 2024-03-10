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

func (c *Post) GetPostById(r *ghttp.Request) {
	id := r.Get("id").Int()

	post, err := service.Post().GetById(r.Context(), id)
	if err != nil {
		r.Response.WriteJsonExit(utility.GetR().Error(consts.ServiceErrCode, err.Error()))
	}

	r.Response.WriteJsonExit(utility.GetR().PUT("post", post.Post).PUT("comments", post.Comments).PUT("is_like", post.IsLike).
		PUT("is_collect", post.IsCollect).PUT("is_follow", post.IsFollow))
}

func (c *Post) Like(r *ghttp.Request) {
	var postLike api.PostLike

	err := r.Parse(&postLike)
	if err != nil {
		r.Response.WriteJsonExit(utility.GetR().Error(consts.RequestErrCode, err.Error()))
	}

	err = service.Post().Like(r.Context(), postLike)
	if err != nil {
		r.Response.WriteJsonExit(utility.GetR().Error(consts.ServiceErrCode, err.Error()))
	}

	r.Response.WriteJsonExit(utility.GetR())
}

func (c *Post) Collect(r *ghttp.Request) {
	var postCollect api.PostCollect

	err := r.Parse(&postCollect)

	if err != nil {
		r.Response.WriteJsonExit(utility.GetR().Error(consts.RequestErrCode, err.Error()))
	}

	err = service.Post().Collect(r.Context(), postCollect)

	if err != nil {
		r.Response.WriteJsonExit(utility.GetR().Error(consts.ServiceErrCode, err.Error()))
	}

	r.Response.WriteJsonExit(utility.GetR())
}
