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

// Add @router /post [put]
// @Summary 添加帖子
// @Description 添加帖子
// @Param Object api.PostAdd
// @accept application/json
// @Produce application/json
// @Success utility.R{code=0,msg="",data={id:int}}
func (c *Post) Add(r *ghttp.Request) {
	var postAdd api.PostAdd

	err := r.Parse(&postAdd)
	res := r.Response
	if err != nil {
		res.WriteJsonExit(utility.GetR().Error(consts.RequestErrCode, err.Error()))
	}

	id, err := service.Post().Add(r.Context(), postAdd)
	if err != nil {
		res.WriteJsonExit(utility.GetR().Error(consts.ServiceErrCode, err.Error()))
	}

	res.WriteJsonExit(utility.GetR().PUT("id", id))
}

func (c *Post) GetPostById(r *ghttp.Request) {
	id := r.Get("id").Int()

	post, err := service.Post().GetById(r.Context(), id)
	if err != nil {
		r.Response.WriteJsonExit(utility.GetR().Error(consts.ServiceErrCode, err.Error()))
	}

	r.Response.WriteJsonExit(utility.GetR().PUT("post", post.Post).PUT("is_like", post.IsLike).
		PUT("is_collect", post.IsCollect).PUT("is_follow", post.IsFollow).PUT("comment_count", post.CommentCount))
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

// GetMinePost @router /post/mine [post]
// @Summary 获取我的帖子
// @Description 获取我的帖子
// @Param Object body get api.GetMinePost
// @accept application/json
// @Produce application/json
// @Success utility.R{code=0,msg="",data={posts:[]api.PostVo}}
func (c *Post) GetMinePost(r *ghttp.Request) {
	var get api.GetMinePost

	err := r.Parse(&get)
	if err != nil {
		r.Response.WriteJsonExit(utility.GetR().Error(consts.RequestErrCode, err.Error()))
	}

	posts, err := service.Post().GetMinePost(r.Context(), get)
	if err != nil {
		r.Response.WriteJsonExit(utility.GetR().Error(consts.ServiceErrCode, err.Error()))
	}

	r.Response.WriteJsonExit(utility.GetR().PUT("posts", posts))
}

// Top @router /post/top [post]
// @Summary 置顶帖子或取消置顶
// @Description 置顶帖子或取消置顶
// @Param Object body get api.TopPost
// @accept application/json
// @Produce application/json
// @Success utility.R{code=0,msg="",data={}}
func (c *Post) Top(r *ghttp.Request) {
	var top api.TopPost

	err := r.Parse(&top)
	if err != nil {
		r.Response.WriteJsonExit(utility.GetR().Error(consts.RequestErrCode, err.Error()))
	}

	err = service.Post().Top(r.Context(), top)

	if err != nil {
		r.Response.WriteJsonExit(utility.GetR().Error(consts.ServiceErrCode, err.Error()))
	}

	r.Response.WriteJsonExit(utility.GetR())
}

// Del @router /post/:id [delete]
// @Summary 删除帖子
// @Description 删除帖子
// @Param id
// @accept application/json
// @Produce application/json
// @Success utility.R{code=0,msg="",data={}}
func (c *Post) Del(r *ghttp.Request) {
	id := r.Get("id")

	err := service.Post().Del(r.Context(), id.Int())
	if err != nil {
		r.Response.WriteJsonExit(utility.GetR().Error(consts.ServiceErrCode, err.Error()))
	}

	r.Response.WriteJsonExit(utility.GetR())
}

// Update @router /post [put]
// @Summary 修改帖子
// @Description 修改帖子
// @Param Object api.PostAdd
// @accept application/json
// @Produce application/json
// @Success utility.R{code=0,msg="",data={}}
func (c *Post) Update(r *ghttp.Request) {
	var update api.PostAdd

	err := r.Parse(&update)
	if err != nil {
		r.Response.WriteJsonExit(utility.GetR().Error(consts.RequestErrCode, err.Error()))
	}

	err = service.Post().Update(r.Context(), update)

	if err != nil {
		r.Response.WriteJsonExit(utility.GetR().Error(consts.ServiceErrCode, err.Error()))
	}

	r.Response.WriteJsonExit(utility.GetR())
}
