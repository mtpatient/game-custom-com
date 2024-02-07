package cmd

import (
	"context"
	"game-custom-com/internal/controller"
	"game-custom-com/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Use(
				ghttp.MiddlewareCORS,
				service.Middleware().Ctx,
			) // 开放域
			/**
			控制器对象
			*/
			cUser := controller.CUser()
			// 不鉴权
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.POST("/register", cUser.Register)
				group.POST("/login", cUser.Login)
				group.Group("/user", func(group *ghttp.RouterGroup) {
					group.GET("/:username", cUser.NameExist)
					group.GET("/is_login", cUser.IsLogin)
				})
			})
			// 鉴权
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(
					service.Middleware().Auth,
				)
				group.Group("/user", func(group *ghttp.RouterGroup) {
					group.GET("/logout", cUser.Logout)
				})
			})
			s.Run()
			return nil
		},
	}
)
