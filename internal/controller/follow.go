package controller

import (
	"game-custom-com/api"
	"game-custom-com/internal/consts"
	"game-custom-com/internal/service"
	"game-custom-com/utility"
	"github.com/gogf/gf/v2/net/ghttp"
)

type Follow struct {
}

func (c *Follow) IsFollow(r *ghttp.Request) {
	id := r.Get("id")

	r.Response.WriteJsonExit(utility.GetR().PUT("isFollow",
		service.Follow().IsFollow(r.Context(), id.Int())))
}

// GetFollowList 获取关注列表
// @router /follow/list/:id [get]
// @Summary 获取关注列表
// @Description 获取关注列表
// @Param id int
// @accept json
// @produce json
// @success  0 utility.R{code=0, msg="", data={"follow_list":[]api.FollowUserVo"}}
func (c *Follow) GetFollowList(r *ghttp.Request) {
	id := r.Get("id")

	list, err := service.Follow().GetFollowList(r.Context(), id.Int())

	if err != nil {
		r.Response.WriteJsonExit(utility.GetR().Error(consts.ServiceErrCode, err.Error()))
	}

	r.Response.WriteJsonExit(utility.GetR().PUT("follow_list", list))
}

// FansList 获取粉丝列表
// @router /follow/fans [post]
// @Summary 获取粉丝列表
// @Description 获取粉丝列表
// @Param {FansGet} api.FansGet
// @accept application/json
// @produce application/json
// success  0 utility.R{code=0, msg="", data={"fans_list":[]api.FansVo}}
func (c *Follow) FansList(r *ghttp.Request) {
	var get api.FansGet

	err := r.Parse(&get)
	if err != nil {
		r.Response.WriteJsonExit(utility.GetR().Error(consts.RequestErrCode, err.Error()))
	}

	fans, err := service.Follow().GetFansList(r.Context(), get)
	if err != nil {
		r.Response.WriteJsonExit(utility.GetR().Error(consts.ServiceErrCode, err.Error()))
	}

	r.Response.WriteJsonExit(utility.GetR().PUT("fans_list", fans))
}
