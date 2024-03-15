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

func (s sMiddleware) AuthAdm(r *ghttp.Request) {
	role, err := service.User().UserRole(r.Context())
	g.Log().Info(r.Context(), role, "管理员鉴权")
	if err != nil || role == 0 {
		r.Response.WriteStatusExit(http.StatusForbidden)
	}
	r.Middleware.Next()
}

func (s sMiddleware) HandleHttpRes(r *ghttp.Request) {
	r.Middleware.Next()

	if r.Response.Status >= http.StatusInternalServerError {
		r.Response.ClearBuffer()
		r.Response.Writeln("哎哟我去，服务器居然开小差了，请稍后再试吧！")
	}

	//errStr := ""
	//if err := r.GetError(); err != nil {
	//	errStr = err.Error()
	//}
	//g.Log().Println(r.Response.Status, r.URL.Path, errStr)
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
		_, err := g.Redis().Do(r.Context(), "expire", consts.TokenKey+token, consts.TokenKeyTTL*60)
		if err != nil {
			r.Response.WriteJsonExit(utility.GetR().Error(consts.RedisErrCode, err.Error()))
		}

		user, _ := service.User().GetById(r.Context(), res.Int())

		customCtx.User = &entity.User{
			Id:          user.Id,
			Username:    user.Username,
			Password:    user.Password,
			Email:       user.Email,
			Avatar:      user.Avatar,
			Signature:   user.Signature,
			Sex:         user.Sex,
			Role:        user.Role,
			Status:      user.Status,
			LikeCount:   user.LikeCount,
			FollowCount: user.FollowCount,
			FansCount:   user.FansCount,
			CreateTime:  user.CreateTime,
			UpdateTime:  user.UpdateTime,
		}

		customCtx.Data["token"] = token
	}
	r.Middleware.Next()
}

func (s sMiddleware) Auth(r *ghttp.Request) {
	if ok, _ := service.User().IsLogin(r.Context()); ok {
		r.Middleware.Next()
	} else {
		r.Response.WriteStatusExit(http.StatusForbidden)
	}
}

func (s sMiddleware) CORS(r *ghttp.Request) {
	//TODO implement me
	panic("implement me")
}
