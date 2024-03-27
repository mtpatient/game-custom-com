package controller

import (
	"game-custom-com/api"
	"game-custom-com/internal/consts"
	"game-custom-com/internal/service"
	"game-custom-com/utility"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type Comment struct {
}

// Add @router /comment [post]
// @Summary 添加评论
// @Description 添加评论
// @Tags comment
// @accept application/json
// @Param body api.CommentAdd
// @Success utility.R{code=0,msg=””,data{}}
func (c *Comment) Add(r *ghttp.Request) {
	var commentAdd api.CommentAdd

	err := r.Parse(&commentAdd)
	g.Log().Info(r.Context(), commentAdd)
	if err != nil {
		r.Response.WriteJsonExit(utility.GetR().Error(consts.RequestErrCode, err.Error()))
	}

	err = service.Comment().Add(r.Context(), commentAdd)

	if err != nil {
		r.Response.WriteJsonExit(utility.GetR().Error(consts.RequestErrCode, err.Error()))
	}

	r.Response.WriteJsonExit(utility.GetR())
}

// GetPostCommentList @router /comment/postCommentList [post]
// @Summary 获取帖子评论列表
// @Description 获取帖子评论列表
// @Tags comment
// @accept application/json
// @Param body api.PostCommentGet
// @Success utility.R{code=0,msg=””,data{comments=[]api.PostCommentRes}}
func (c *Comment) GetPostCommentList(r *ghttp.Request) {
	var postCommentGet api.PostCommentGet

	if err := r.Parse(&postCommentGet); err != nil {
		r.Response.WriteJsonExit(utility.GetR().Error(consts.RequestErrCode, err.Error()))
	}

	comments, err := service.Comment().GetPostCommentList(r.Context(), postCommentGet)

	if err != nil {
		r.Response.WriteJsonExit(utility.GetR().Error(consts.RequestErrCode, err.Error()))
	}

	r.Response.WriteJsonExit(utility.GetR().PUT("comments", comments))
}

// Del @router /comment/:id [delete]
// @Summary 删除评论
// @Description 删除评论
// @Tags comment
// @accept application/json
// @Param id int
// @Success utility.R{code=0,msg=””,data{}}
func (c *Comment) Del(r *ghttp.Request) {
	id := r.Get("id")

	err := service.Comment().Del(r.Context(), id.Int())

	if err != nil {
		r.Response.WriteJsonExit(utility.GetR().Error(consts.RequestErrCode, err.Error()))
	}

	r.Response.WriteJsonExit(utility.GetR())
}

// Like @router /comment/like/:id [post]
// @Summary 点赞评论
// @Description 点赞评论
// @Tags comment
// @accept application/json
// @Param api.CommentLike
// @Success utility.R{code=0,msg=””,data{}}
func (c *Comment) Like(r *ghttp.Request) {
	var like api.CommentLike
	err := r.Parse(&like)
	if err != nil {
		g.Log().Error(r.Context(), err)
		r.Response.WriteJsonExit(utility.GetR().Error(consts.RequestErrCode, err.Error()))
	}

	err = service.Comment().Like(r.Context(), like)

	if err != nil {
		r.Response.WriteJsonExit(utility.GetR().Error(consts.RequestErrCode, err.Error()))
	}

	r.Response.WriteJsonExit(utility.GetR())
}

// GetMineComments @router /comment/getMine [post]
// @Summary 获取个人评论列表
// @Description 获取个人评论列表
// @Tags comment
// @accept application/json
// @Param body api.CommentGet
// @Success utility.R{code=0,msg=””,data{comments=[]api.CommentRes}}
func (c *Comment) GetMineComments(r *ghttp.Request) {
	var get api.CommentGet

	err := r.Parse(&get)
	if err != nil {
		r.Response.WriteJsonExit(utility.GetR().Error(consts.RequestErrCode, err.Error()))
	}

	comments, err := service.Comment().GetMineComments(r.Context(), get)
	if err != nil {
		r.Response.WriteJsonExit(utility.GetR().Error(consts.RequestErrCode, err.Error()))
	}
	r.Response.WriteJsonExit(utility.GetR().PUT("comments", comments))
}

// GetCommentById @router /comment/:id [get]
// @Summary 获取单个评论
// @Description 获取单个评论
// @Tags comment
// @accept application/json
// @Param id
// @Success utility.R{code=0,msg=””,data{comments=[]api.PostCommentRes}}
func (c *Comment) GetCommentById(r *ghttp.Request) {
	id := r.Get("id")

	comment, err := service.Comment().GetCommentById(r.Context(), id.Int())
	if err != nil {
		r.Response.WriteJsonExit(utility.GetR().Error(consts.RequestErrCode, err.Error()))
	}

	r.Response.WriteJsonExit(utility.GetR().PUT("comment", comment))
}

// CommentList @router /comment/list [post]
// @Summary 获取评论列表
// @Description 获取评论列表
// @Tags comment
// @accept application/json
// @Param Object api.CommonParams
// @Success utility.R{code=0,msg=””,data{comments=[]api.CommentRes}}
func (c *Comment) CommentList(r *ghttp.Request) {
	var params api.CommonParams
	if err := r.Parse(&params); err != nil {
		r.Response.WriteJsonExit(utility.GetR().Error(consts.RequestErrCode, err.Error()))
	}
	comments, total, err := service.Comment().CommentList(r.Context(), params)
	if err != nil {
		r.Response.WriteJsonExit(utility.GetR().Error(consts.RequestErrCode, err.Error()))
	}
	r.Response.WriteJsonExit(utility.GetR().PUT("comments", comments).PUT("total", total))
}
