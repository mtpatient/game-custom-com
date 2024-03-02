package controller

import (
	"game-custom-com/internal/consts"
	"game-custom-com/internal/model/entity"
	"game-custom-com/internal/service"
	"game-custom-com/utility"
	"github.com/gogf/gf/v2/net/ghttp"
)

type Section struct {
}

func (c *Section) Add(r *ghttp.Request) {
	var section entity.Section
	res := r.Response
	err := r.Parse(&section)
	if err != nil {
		res.WriteJsonExit(utility.GetR().Error(consts.RequestErrCode, err.Error()))
	}
	err = service.Section().Add(r.Context(), section)
	if err != nil {
		res.WriteJsonExit(utility.GetR().Error(consts.ServiceErrCode, err.Error()))
	}
	res.WriteJsonExit(utility.GetR())
}

func (c *Section) Update(r *ghttp.Request) {
	var section entity.Section
	res := r.Response
	err := r.Parse(&section)
	if err != nil {
		res.WriteJsonExit(utility.GetR().Error(consts.RequestErrCode, err.Error()))
	}
	err = service.Section().Update(r.Context(), section)
	if err != nil {
		res.WriteJsonExit(utility.GetR().Error(consts.ServiceErrCode, err.Error()))
	}
	res.WriteJsonExit(utility.GetR())
}

func (c *Section) GetAll(r *ghttp.Request) {
	all, err := service.Section().GetAll(r.Context())
	res := r.Response
	if err != nil {
		res.WriteJsonExit(utility.GetR().Error(consts.ServiceErrCode, err.Error()))
	}
	res.WriteJsonExit(utility.GetR().PUT("sections", all))
}
