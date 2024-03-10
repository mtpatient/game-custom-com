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
				// 开放域
				ghttp.MiddlewareCORS,
				service.Middleware().Ctx,
				service.Middleware().HandleHttpRes,
			)
			/**
			控制器对象
			*/
			cUser := new(controller.User)
			cImg := new(controller.Img)
			cSection := new(controller.Section)
			cPost := new(controller.Post)
			cFeedback := new(controller.Feedback)
			// 不鉴权
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.POST("/register", cUser.Register)
				group.POST("/login", cUser.Login)
				group.Group("/user", func(group *ghttp.RouterGroup) {
					group.GET("/name-exist/:username", cUser.NameExist)
					group.GET("/is_login", cUser.IsLogin)
					group.GET("/:id", cUser.GetUser)
					group.GET("/authCode/:username", cUser.GetAuthCode)
					group.POST("/resetPwd", cUser.ResetPwd)
				})
				group.Group("/post", func(group *ghttp.RouterGroup) {
					group.GET("/:id", cPost.GetPostById)
				})
				group.Group("/section", func(group *ghttp.RouterGroup) {
					group.GET("/all", cSection.GetAll)
					group.GET("/:id", cSection.GetById)
				})
			})
			// 鉴权: 普通用户
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(
					service.Middleware().Auth,
				)
				group.Group("/user", func(group *ghttp.RouterGroup) {
					group.GET("/logout", cUser.Logout)
					group.PUT("/", cUser.Update)
					group.POST("/password", cUser.ReplacePassword)
					group.POST("/follow", cUser.Follow)
				})
				group.Group("/img", func(group *ghttp.RouterGroup) {
					group.GET("/getSignature/:count", cImg.GetSignatures)
					group.GET("/avatars", cImg.GetAllAvatar)
				})
				group.Group("/post", func(group *ghttp.RouterGroup) {
					group.POST("/", cPost.Add)
					group.POST("/like", cPost.Like)
					group.POST("/collect", cPost.Collect)
				})
				group.Group("/feedback", func(group *ghttp.RouterGroup) {
					group.POST("/", cFeedback.Add)
				})
			})
			// 鉴权: 管理员
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(
					service.Middleware().AuthAdm,
				)
				group.Group("/section", func(group *ghttp.RouterGroup) {
					group.POST("/", cSection.Add)
					group.PUT("/", cSection.Update)
				})
				group.Group("/img", func(group *ghttp.RouterGroup) {
					group.POST("/", cImg.Save)
					group.DELETE("/avatar/:id", cImg.DeleteAvatar)
					group.PUT("/", cImg.Update)
				})
			})
			s.Run()
			return nil
		},
	}
)
