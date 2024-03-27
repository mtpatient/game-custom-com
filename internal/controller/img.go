package controller

import (
	"game-custom-com/api"
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

// PostImgList @router /img/list [post]
// @Summary 获取图片列表
// @Description 获取图片列表
// @Tags image
// @accept application/json
// @Param body api.CommonParams
// @Success utility.R{code=0,msg=””,data{images=[]api.PostImage}}
func (c *Img) PostImgList(r *ghttp.Request) {
	var params api.CommonParams
	if err := r.Parse(&params); err != nil {
		r.Response.WriteJsonExit(utility.GetR().Error(consts.RequestErrCode, err.Error()))
	}
	imgs, total, err := service.Img().PostImgList(r.Context(), params)
	if err != nil {
		r.Response.WriteJsonExit(utility.GetR().Error(consts.ServiceErrCode, err.Error()))
	}
	r.Response.WriteJsonExit(utility.GetR().PUT("images", imgs).PUT("total", total))
}

// Del @router /img/:id [del]
// @Summary 删除图片
// @Description 删除图片
// @Tags image
// @accept application/json
// @Param id
// @Success utility.R{code=0,msg=””,data{}}
func (c *Img) Del(r *ghttp.Request) {
	id := r.Get("id")
	if err := service.Img().Del(r.Context(), id.Int()); err != nil {
		r.Response.WriteJsonExit(utility.GetR().Error(consts.ServiceErrCode, err.Error()))
	}
	r.Response.WriteJsonExit(utility.GetR())
}

// GetSlideshow @router /img/slideshow [get]
// @Summary 获取轮播图列表
// @Description 获取轮播图列表
// @Tags image
// @accept application/json
// @Param
// @Success utility.R{code=0,msg=””,data{images=[]api.PostImage}}
func (c *Img) GetSlideshow(r *ghttp.Request) {
	imgs, err := service.Img().GetSlideshow(r.Context())
	if err != nil {
		r.Response.WriteJsonExit(utility.GetR().Error(consts.ServiceErrCode, err.Error()))
	}
	r.Response.WriteJsonExit(utility.GetR().PUT("images", imgs))
}

// SaveSlideshow @router /img/slideshow [post]
// @Summary 获取轮播图列表
// @Description 获取轮播图列表
// @Tags image
// @accept application/json
// @Param
// @Success utility.R{code=0,msg=””,data{}}
func (c *Img) SaveSlideshow(r *ghttp.Request) {
	var params api.SlideshowParams

	if err := r.Parse(&params); err != nil {
		r.Response.WriteJsonExit(utility.GetR().Error(consts.RequestErrCode, err.Error()))
	}
	if err := service.Img().SaveSlideshow(r.Context(), params); err != nil {
		r.Response.WriteJsonExit(utility.GetR().Error(consts.ServiceErrCode, err.Error()))
	}
	r.Response.WriteJsonExit(utility.GetR())
}

// UpdateSlideshow @router /img/slideshow [put]
// @Summary 获取轮播图列表
// @Description 获取轮播图列表
// @Tags image
// @accept application/json
// @Param
// @Success utility.R{code=0,msg=””,data{}}
func (c *Img) UpdateSlideshow(r *ghttp.Request) {
	var params api.SlideshowParams
	if err := r.Parse(&params); err != nil {
		r.Response.WriteJsonExit(utility.GetR().Error(consts.RequestErrCode, err.Error()))
	}
	if err := service.Img().UpdateSlideshow(r.Context(), params); err != nil {
		r.Response.WriteJsonExit(utility.GetR().Error(consts.ServiceErrCode, err.Error()))
	}
	r.Response.WriteJsonExit(utility.GetR())
}
