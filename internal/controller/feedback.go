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

func (c *Feedback) Add(r *ghttp.Request) {
	var feedbackAdd api.FeedbackAdd
	if err := r.Parse(&feedbackAdd); err != nil {
		r.Response.WriteJsonExit(utility.GetR().Error(consts.RequestErrCode, err.Error()))
	}

	err := service.Feedback().Create(r.Context(), feedbackAdd)
	if err != nil {
		r.Response.WriteJsonExit(utility.GetR().Error(consts.ServiceErrCode, err.Error()))
	}

	r.Response.WriteJsonExit(utility.GetR())
}
