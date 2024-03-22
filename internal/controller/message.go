package controller

import (
	"game-custom-com/api"
	"game-custom-com/internal/consts"
	"game-custom-com/internal/service"
	"game-custom-com/utility"
	"github.com/gogf/gf/v2/net/ghttp"
)

type Message struct {
}

func CMessage() *Message {
	return &Message{}
}

// GetLikesMessage @router /message/likes
// @Summary 返回该用户的点赞列表
// @Description 返回该用户的点赞列表
// @Tags message
// @accept application/json
// @Param body api.GetLikesParams
// @Success utility.R{code=0,msg=””,data{comments=[]api.CommentRes}}
func (c *Message) GetLikesMessage(r *ghttp.Request) {
	var get api.GetLikesParams
	if err := r.Parse(&get); err != nil {
		r.Response.WriteJsonExit(utility.GetR().Error(consts.RequestErrCode, err.Error()))
	}

	likes, err := service.Message().GetLikesMessage(r.Context(), get)
	if err != nil {
		r.Response.WriteJsonExit(utility.GetR().Error(consts.ServiceErrCode, err.Error()))
	}

	r.Response.WriteJsonExit(utility.GetR().PUT("messages", likes))
}

// GetMessageNew @router /message/news
// @Summary 返回该用户的新消息数
// @Description 返回该用户的新消息数
// @Tags message
// @accept application/json
// @Param none
// @Success utility.R{code=0,msg=””,data{news=[]int}}
func (c *Message) GetMessageNew(r *ghttp.Request) {
	news, err := service.Message().GetMessageNew(r.Context())
	if err != nil {
		r.Response.WriteJsonExit(utility.GetR().Error(consts.ServiceErrCode, err.Error()))
	}
	r.Response.WriteJsonExit(utility.GetR().PUT("news", news))
}

// Read @router /message/read/:id
// @Summary 将消息标记为已读
// @Description 将消息标记为已读
// @Tags message
// @accept application/json
// @Param id
// @Success utility.R{code=0,msg=””,data{}}
func (c *Message) Read(r *ghttp.Request) {
	id := r.Get("id").Int()

	err := service.Message().Read(r.Context(), id)
	if err != nil {
		r.Response.WriteJsonExit(utility.GetR().Error(consts.ServiceErrCode, err.Error()))
	}
	r.Response.WriteJsonExit(utility.GetR())
}
