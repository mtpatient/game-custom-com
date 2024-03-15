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
