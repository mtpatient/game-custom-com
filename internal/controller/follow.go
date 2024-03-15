package controller

import (
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
