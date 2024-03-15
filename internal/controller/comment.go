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

func (c *Comment) Add(r *ghttp.Request) {
	var commentAdd api.CommentAdd
	err := r.Parse(&commentAdd)
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
// @Success utility.R{code=0,msg=””,data{comments=api.Comments}}
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
