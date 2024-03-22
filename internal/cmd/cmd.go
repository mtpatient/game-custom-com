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
			cFollow := new(controller.Follow)
			cComment := new(controller.Comment)
			cMessage := new(controller.Message)
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
					group.POST("/search", cUser.SearchUser)
				})
				group.Group("/post", func(group *ghttp.RouterGroup) {
					group.GET("/:id", cPost.GetPostById)
					group.POST("/getMine", cPost.GetMinePost)
					group.GET("/top/:id", cPost.GetTopPost)
					group.POST("/list", cPost.GetPostList)
					group.POST("/search", cPost.SearchPost)
				})
				group.Group("/section", func(group *ghttp.RouterGroup) {
					group.GET("/all", cSection.GetAll)
					group.GET("/:id", cSection.GetById)
				})
				group.Group("/comment", func(group *ghttp.RouterGroup) {
					group.POST("/getPostCommentList", cComment.GetPostCommentList)
					group.POST("/getMineComments", cComment.GetMineComments)
					group.GET("/:id", cComment.GetCommentById)
				})
				group.Group("/follow", func(group *ghttp.RouterGroup) {
					group.GET("/list/:id", cFollow.GetFollowList)
					group.POST("/fans", cFollow.FansList)
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
					group.DELETE("/:id", cPost.Del)
					group.POST("/top", cPost.Top)
					group.PUT("/", cPost.Update)
					group.POST("/follow", cPost.GetFollow)
				})
				group.Group("/feedback", func(group *ghttp.RouterGroup) {
					group.POST("/", cFeedback.Add)
				})
				group.Group("/follow", func(group *ghttp.RouterGroup) {
					group.GET("/isFollow/:id", cFollow.IsFollow)
				})
				group.Group("/comment", func(group *ghttp.RouterGroup) {
					group.POST("/", cComment.Add)
					group.DELETE("/:id", cComment.Del)
					group.POST("/like", cComment.Like)
				})
				group.Group("/message", func(group *ghttp.RouterGroup) {
					group.POST("/likes", cMessage.GetLikesMessage)
					group.GET("/news", cMessage.GetMessageNew)
					group.GET("/read/:id", cMessage.Read)
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
