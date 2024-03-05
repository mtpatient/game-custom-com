package controller

import (
	"game-custom-com/internal/consts"
	"game-custom-com/internal/model/entity"
	"game-custom-com/internal/service"
	"game-custom-com/utility"
	"github.com/gogf/gf/v2/net/ghttp"
)

type Img struct {
}

func (c *Img) CImg() *Img {
	return &Img{}
}

func (c *Img) GetSignatures(r *ghttp.Request) {
	count := r.Get("count")
	signatures, err := service.Img().GetSignatures(r.Context(), count.Int())
	if err != nil {
		r.Response.WriteJsonExit(utility.GetR().Error(consts.ServiceErrCode, err.Error()))
	}
	r.Response.WriteJsonExit(utility.GetR().PUT("signatures", signatures))
}

func (c *Img) Save(r *ghttp.Request) {
	var imgs []entity.Image

	err := r.Parse(&imgs)
	if err != nil {
		r.Response.WriteJsonExit(utility.GetR().Error(consts.RequestErrCode, err.Error()))
	}
	err = service.Img().Save(r.Context(), imgs)
	if err != nil {
		r.Response.WriteJsonExit(utility.GetR().Error(consts.ServiceErrCode, err.Error()))
	}
	r.Response.WriteJsonExit(utility.GetR())
}

func (c *Img) GetAllAvatar(r *ghttp.Request) {
	avatar, err := service.Img().GetAllAvatar(r.Context())
	if err != nil {
		r.Response.WriteJsonExit(utility.GetR().Error(consts.ServiceErrCode, err.Error()))
	}
	r.Response.WriteJsonExit(utility.GetR().PUT("avatars", avatar))
}

func (c *Img) DeleteAvatar(r *ghttp.Request) {
	id := r.Get("id")

	err := service.Img().DeleteAvatar(r.Context(), id.Int())

	if err != nil {
		r.Response.WriteJsonExit(utility.GetR().Error(consts.ServiceErrCode, err.Error()))
	}

	r.Response.WriteJsonExit(utility.GetR())
}

func (c *Img) Update(r *ghttp.Request) {
	var img entity.Image

	err := r.Parse(&img)
	if err != nil {
		r.Response.WriteJsonExit(utility.GetR().Error(consts.RequestErrCode, err.Error()))
	}

	err = service.Img().Update(r.Context(), img)

	if err != nil {
		r.Response.WriteJsonExit(utility.GetR().Error(consts.ServiceErrCode, err.Error()))
	}

	r.Response.WriteJsonExit(utility.GetR())
}
