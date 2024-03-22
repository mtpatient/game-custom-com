package api

import "game-custom-com/internal/model/entity"

type PostAdd struct {
	Post   entity.Post `json:"post"`
	Images []string    `json:"images"`
}

type PostDetail struct {
	Post         entity.Post `json:"post"`
	IsLike       int         `json:"is_like"`
	IsCollect    int         `json:"is_collect"`
	IsFollow     int         `json:"is_follow"`
	CommentCount int         `json:"comment_count"`
}

type PostLike struct {
	PostId   int `json:"post_id" `
	ToUserId int `json:"user_id" `
	Operate  int `json:"operate" ` // like or unlike
	Status   int `json:"status"`
}

type PostCollect struct {
	PostId  int `json:"post_id" `
	Operate int `json:"operate" ` // like or unlike
	Status  int `json:"status"`
}

type GetPostParams struct {
	Id        int `json:"id"`
	PageIndex int `json:"page_index" v:"required"`
	PageSize  int `json:"page_size" v:"required"`
	ShowType  int `json:"show_type"`
}

type PostVo struct {
	entity.Post
	CommentCount int      `json:"comment_count"`
	Username     string   `json:"username"`
	Avatar       string   `json:"avatar"`
	IsLike       int      `json:"isLike"`
	ImgList      []string `json:"imgList"`
}

type TopPost struct {
	Id      int `json:"id" v:"required"`
	Operate int `json:"operate" ` // 1，个人置顶 2，个人取消置顶 3，官方板块置顶 4，官方板块取消置顶
}

type TopPostVo struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}

type SearchParams struct {
	Key       string `json:"key" v:"required"`
	PageIndex int    `json:"page_index" v:"required"`
	PageSize  int    `json:"page_size" v:"required"`
	ShowType  int    `json:"show_type" v:"required"`
}
