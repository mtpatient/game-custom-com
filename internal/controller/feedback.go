package controller

import (
	"game-custom-com/api"
	"game-custom-com/internal/consts"
	"game-custom-com/internal/service"
	"game-custom-com/utility"
	"github.com/gogf/gf/v2/net/ghttp"
)

type Feedback struct {
}

// Add @router /feedback [post]
// @Summary 提交反馈
// @Description 提交反馈
// @Tags feedback
// @accept application/json
// @Param Object api.FeedbackVo
// @Success utility.R{code=0,msg=””,data{}}
func (c *Feedback) Add(r *ghttp.Request) {
	var feedbackAdd api.FeedbackVo
	if err := r.Parse(&feedbackAdd); err != nil {
		r.Response.WriteJsonExit(utility.GetR().Error(consts.RequestErrCode, err.Error()))
	}

	err := service.Feedback().Create(r.Context(), feedbackAdd)
	if err != nil {
		r.Response.WriteJsonExit(utility.GetR().Error(consts.ServiceErrCode, err.Error()))
	}

	r.Response.WriteJsonExit(utility.GetR())
}

// List @router /feedback/list [post]
// @Summary 提交反馈
// @Description 提交反馈
// @Tags feedback
// @accept application/json
// @Param Object api.CommonParams
// @Success utility.R{code=0,msg=””,data{list=[]api.FeedbackVo}}
func (c *Feedback) List(r *ghttp.Request) {
	var params api.CommonParams
	if err := r.Parse(&params); err != nil {
		r.Response.WriteJsonExit(utility.GetR().Error(consts.RequestErrCode, err.Error()))
	}
	list, total, err := service.Feedback().GetList(r.Context(), params)
	if err != nil {
		r.Response.WriteJsonExit(utility.GetR().Error(consts.ServiceErrCode, err.Error()))
	}
	r.Response.WriteJsonExit(utility.GetR().PUT("list", list).PUT("total", total))
}
