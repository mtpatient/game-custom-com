package logic

import (
	"game-custom-com/internal/consts"
	"game-custom-com/internal/model"
	"game-custom-com/internal/model/entity"
	"game-custom-com/internal/service"
	"game-custom-com/utility"
	"net/http"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type (
	sMiddleware struct{}
)

func init() {
	service.RegisterMiddleware(&sMiddleware{})
}
func (s sMiddleware) Ctx(r *ghttp.Request) {
	customCtx := &model.Context{
		Data: make(g.Map),
	}
	service.Context().Init(r, customCtx)

	var token string
	if r.Header["Token"] != nil {
		//r.Response.WriteStatusExit(http.StatusBadRequest)
		token = r.Header["Token"][0]
	}

	if res, _ := g.Redis().Get(r.Context(), consts.TokenKey+token); res.Int() > 0 {
		_, err := g.Redis().Do(r.Context(), "expire", consts.TokenKey+token, consts.TokenKeyTTL)
		if err != nil {
			r.Response.WriteJsonExit(utility.GetR().Error(consts.RedisErrCode, err.Error()))
		}

		user, _ := service.User().GetById(r.Context(), res.Int64())

		customCtx.User = &entity.User{
			Id:         user.Id,
			Username:   user.Username,
			Password:   user.Password,
			Email:      user.Email,
			Img:        user.Img,
			Signature:  user.Signature,
			Sex:        user.Sex,
			Role:       user.Role,
			Status:     user.Status,
			CreateTime: user.CreateTime,
			UpdateTime: user.UpdateTime,
		}

		customCtx.Data["token"] = token
	}
	r.Middleware.Next()
}

func (s sMiddleware) Auth(r *ghttp.Request) {
	if service.User().IsLogin(r.Context()) {
		r.Middleware.Next()
	} else {
		r.Response.WriteStatus(http.StatusForbidden)
	}
}

func (s sMiddleware) CORS(r *ghttp.Request) {
	//TODO implement me
	panic("implement me")
}
