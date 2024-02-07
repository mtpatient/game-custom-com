package logic

import (
	"context"
	"game-custom-com/internal/consts"
	"game-custom-com/internal/model"
	"game-custom-com/internal/model/entity"
	"game-custom-com/internal/service"
	"github.com/gogf/gf/v2/net/ghttp"
)

type sContext struct {
}

func init() {
	service.RegisterContext(&sContext{})
}

func (s sContext) Init(r *ghttp.Request, ctx *model.Context) {
	r.SetCtxVar(consts.ContextKey, ctx)
}

func (s sContext) Get(ctx context.Context) *model.Context {
	value := ctx.Value(consts.ContextKey)
	if value == nil {
		return nil
	}
	if localCtx, ok := value.(*model.Context); ok {
		return localCtx
	}
	return nil
}

func (s sContext) SetUser(ctx context.Context, ctxUser *entity.User) {
	s.Get(ctx).User = ctxUser
}
