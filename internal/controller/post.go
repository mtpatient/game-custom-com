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

// Add @router /post [post]
// @Summary 添加帖子
// @Description 添加帖子
// @Tags post
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

// GetPostById @router /post [put]
// @Summary 根据id获取帖子
// @Description 根据id获取帖子
// @Tags post
// @Param id
// @accept application/json
// @Produce application/json
// @Success utility.R{code=0,msg="",data={api.PostDetail}}
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
// @Tags post
// @Param Object body get api.GetPostParams
// @accept application/json
// @Produce application/json
// @Success utility.R{code=0,msg="",data={posts:[]api.PostVo}}
func (c *Post) GetMinePost(r *ghttp.Request) {
	var get api.GetPostParams

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
// @Tags post
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
// @Tags post
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
// @Tags post
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

// GetTopPost @router /post/top/:id [get]
// @Summary 获取板块置顶的帖子
// @Description 获取板块置顶的帖子
// @Tags post
// @Param int id
// @accept application/json
// @Produce application/json
// @Success utility.R{code=0,msg="",data={[]api.TopPostVo}}
func (c *Post) GetTopPost(r *ghttp.Request) {
	id := r.Get("id").Int()

	posts, err := service.Post().GetTopPost(r.Context(), id)

	if err != nil {
		r.Response.WriteJsonExit(utility.GetR().Error(consts.ServiceErrCode, err.Error()))
	}

	r.Response.WriteJsonExit(utility.GetR().PUT("posts", posts))

}

// GetFollow @router /post/follow [get]
// @Summary 获取关注的人帖子
// @Description 获取关注的人帖子
// @Tags post
// @Param
// @accept application/json
// @Produce application/json
// @Success utility.R{code=0,msg="",data={posts:[]api.PostVo}}
func (c *Post) GetFollow(r *ghttp.Request) {
	var get api.GetPostParams
	err := r.Parse(&get)

	posts, err := service.Post().GetFollow(r.Context(), get)
	if err != nil {
		r.Response.WriteJsonExit(utility.GetR().Error(consts.ServiceErrCode, err.Error()))
	}
	r.Response.WriteJsonExit(utility.GetR().PUT("posts", posts))
}

// GetPostList @router /post/list [post]
// @Summary 获取各个板块的帖子
// @Description 获取各个板块的帖子
// @Tags post
// @Param Object body get api.GetPostParams
// @accept application/json
// @Produce application/json
// @Success utility.R{code=0,msg="",data={posts:[]api.PostVo}}
func (c *Post) GetPostList(r *ghttp.Request) {
	var get api.GetPostParams
	err := r.Parse(&get)
	if err != nil {
		r.Response.WriteJsonExit(utility.GetR().Error(consts.RequestErrCode, err.Error()))
	}

	posts, err := service.Post().GetPostList(r.Context(), get)
	if err != nil {
		r.Response.WriteJsonExit(utility.GetR().Error(consts.ServiceErrCode, err.Error()))
	}

	r.Response.WriteJsonExit(utility.GetR().PUT("posts", posts))
}

// SearchPost @router /post/search [post]
// @Summary 搜索帖子
// @Description 搜索帖子
// @Tags post
// @Param Object body get api.SearchParams
// @accept application/json
// @Produce application/json
// @Success utility.R{code=0,msg="",data={posts:[]api.PostVo}}
func (c *Post) SearchPost(r *ghttp.Request) {
	var get api.SearchParams
	err := r.Parse(&get)
	if err != nil {
		r.Response.WriteJsonExit(utility.GetR().Error(consts.RequestErrCode, err.Error()))
	}

	posts, err := service.Post().SearchPost(r.Context(), get)
	if err != nil {
		r.Response.WriteJsonExit(utility.GetR().Error(consts.ServiceErrCode, err.Error()))
	}

	r.Response.WriteJsonExit(utility.GetR().PUT("posts", posts))
}

// PostList @router /post/bm/list [post]
// @Summary 获取各个板块的帖子
// @Description 获取各个板块的帖子
// @Tags post
// @Param Object body get api.CommonParams
// @accept application/json
// @Produce application/json
// @Success utility.R{code=0,msg="",data={posts=[]entity.Post}}
func (c *Post) PostList(r *ghttp.Request) {
	var get api.CommonParams
	err := r.Parse(&get)
	if err != nil {
		r.Response.WriteJsonExit(utility.GetR().Error(consts.RequestErrCode, err.Error()))
	}

	posts, total, err := service.Post().PostList(r.Context(), get)
	if err != nil {
		r.Response.WriteJsonExit(utility.GetR().Error(consts.ServiceErrCode, err.Error()))
	}

	r.Response.WriteJsonExit(utility.GetR().PUT("posts", posts).PUT("total", total))
}

// UpdateStatus @router /post/bm/status [put]
// @Summary 获取各个板块的帖子
// @Description 获取各个板块的帖子
// @Tags post
// @Param Object body get api.UpdateStatus
// @accept application/json
// @Produce application/json
// @Success utility.R{code=0,msg="",data={}}
func (c *Post) UpdateStatus(r *ghttp.Request) {
	var update api.UpdateStatus
	err := r.Parse(&update)
	if err != nil {
		r.Response.WriteJsonExit(utility.GetR().Error(consts.RequestErrCode, err.Error()))
	}
	err = service.Post().UpdateStatus(r.Context(), update)
	if err != nil {
		r.Response.WriteJsonExit(utility.GetR().Error(consts.ServiceErrCode, err.Error()))
	}
	r.Response.WriteJsonExit(utility.GetR())
}
