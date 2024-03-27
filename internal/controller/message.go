package controller

import (
	"context"
	"game-custom-com/api"
	"game-custom-com/internal/consts"
	"game-custom-com/internal/service"
	"game-custom-com/utility"
	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
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
// @Param body api.Params
// @Success utility.R{code=0,msg=””,data{comments=[]api.CommentRes}}
func (c *Message) GetLikesMessage(r *ghttp.Request) {
	var get api.Params
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

// Subscription @router /message/subscription [get]
// @Summary 订阅消息
// @Description 订阅消息
// @Tags message
// @accept application/json
// @Param none
// @Success utility.R{code=0,msg=””,data{}}
func (c *Message) Subscription(r *ghttp.Request) {

	r.Response.Header().Set("Content-Type", "text/event-stream")
	r.Response.Header().Set("Cache-Control", "no-cache")
	r.Response.Header().Set("Connection", "keep-alive")
	r.Response.Header().Set("Access-Control-Allow-Origin", "*")

	evId := r.Get("lastEventId")
	if evId != nil {
		r.Response.Write("重新建立链接！\n\n")
		r.Response.Flush()
	}

	ctx := r.Context()
	uid := service.Context().Get(ctx).User.Id
	//uid := r.Get("uid").Int()
	sub, _, _ := g.Redis().Subscribe(ctx, gconv.String(uid))
	defer func(sub gredis.Conn, ctx context.Context) {
		err := sub.Close(ctx)
		//g.Log().Info(ctx, "关闭订阅", err)
		if err != nil {
			g.Log().Error(ctx, err)
		}
	}(sub, ctx)

	//g.Log().Info(ctx, "订阅消息", uid)
	r.Response.Write("建立连接！\n\n")
	r.Response.Flush()

	i := 0
	for {
		//g.Log().Info(ctx, "订阅消息:for")
		msg, err := sub.ReceiveMessage(ctx)
		if err != nil {
			g.Log().Error(ctx, err)
			break
		}
		//g.Log().Info(ctx, "发送消息", msg.Payload)
		i++
		r.Response.Write("id:" + gconv.String(i) + "\n")
		r.Response.Write("event:message\n")
		r.Response.Write("retry:3000\n")
		r.Response.Write("data:" + msg.Payload + "\n\n")
		r.Response.Flush()
		if r.IsExited() {
			//r.Context().Deadline()
			break
		}
	}
}

// GetNotice @router /message/notice [post]
// @Summary 获取通知
// @Description 获取通知
// @Tags message
// @accept application/json
// @Param body api.Params
// @Success utility.R{code=0,msg=””,data{notice=[]entity.Message}}
func (c *Message) GetNotice(r *ghttp.Request) {
	var params api.Params
	if err := r.Parse(&params); err != nil {
		r.Response.WriteJsonExit(utility.GetR().Error(consts.RequestErrCode, err.Error()))
	}
	notice, err := service.Message().GetNotice(r.Context(), params)
	if err != nil {
		r.Response.WriteJsonExit(utility.GetR().Error(consts.ServiceErrCode, err.Error()))
	}
	r.Response.WriteJsonExit(utility.GetR().PUT("notice", notice))
}
