package controller

import (
	"game-custom-com/api"
	"game-custom-com/internal/consts"
	"game-custom-com/internal/service"
	"game-custom-com/utility"
	"github.com/gogf/gf/v2/net/ghttp"
)

type AdmLog struct {
}

// List @router /img/:id [del]
// @Summary 获取管理日志
// @Description 获取管理日志
// @Tags adm-log
// @accept application/json
// @Param Object api.CommonParams
// @Success utility.R{code=0,msg=””,data{list=[]api.AdmLogVo,total=int}}
func (c *AdmLog) List(r *ghttp.Request) {
	var params api.CommonParams
	if err := r.Parse(&params); err != nil {
		r.Response.WriteJsonExit(utility.GetR().Error(consts.RequestErrCode, err.Error()))
	}

	list, total, err := service.AdmLog().GetList(r.Context(), params)
	if err != nil {
		r.Response.WriteJsonExit(utility.GetR().Error(consts.ServiceErrCode, err.Error()))
	}

	r.Response.WriteJsonExit(utility.GetR().PUT("list", list).PUT("total", total))
}
