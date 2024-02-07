package service

import (
	"context"
	"game-custom-com/internal/model"
	"game-custom-com/internal/model/entity"
	"github.com/gogf/gf/v2/net/ghttp"
)

type IContext interface {
	Init(r *ghttp.Request, ctx *model.Context)
	Get(ctx context.Context) *model.Context
	SetUser(ctx context.Context, ctxUser *entity.User)
}

var localContext IContext

func Context() IContext {
	if localContext == nil {
		panic("implement not found for interface IUser, forget register?")
	}

	return localContext
}

func RegisterContext(i IContext) {
	localContext = i
}
