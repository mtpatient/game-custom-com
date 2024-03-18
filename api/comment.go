package api

import (
	"game-custom-com/internal/model/entity"
)

type CommentAdd struct {
	Content   string `json:"content"`
	IsFloor   bool   `json:"is_floor"`
	ParentId  int    `json:"parent_id"`
	PostId    int    `json:"post_id"`
	ToUserId  int    `json:"to_user_id"`
	CommentId int    `json:"comment_id"`
}

type CommentVo struct {
	entity.Comment
	UserName  string `json:"username" orm:"with:id=user_id"`
	ReplyName string `json:"reply_name" orm:"with:id=reply_id"`
	Avatar    string `json:"avatar" orm:"with:id=user_id"`
	IsLike    int    `json:"isLike"`
	IsOwn     int    `json:"isOwn"`
}

type PostCommentRes struct {
	CommentVo
	Children []CommentVo `json:"child"`
}

type PostCommentGet struct {
	PostId          int `json:"post_id"`           //post_id: this.$route.params.id,
	PageSize        int `json:"page_size"`         //page_size: 20,
	IsOnlyPublisher int `json:"is_only_publisher"` //isOnlyPublisher: 0,
	PageIndex       int `json:"page_index"`        //page_index: 1,
	ShowType        int `json:"show_type"`         // 1: 点赞数排行 2: 发布时间排行
}

type CommentLike struct {
	Id       int `json:"id" v:"required"`
	Operate  int `json:"operate" v:"required"`
	ToUserId int `json:"to_user_id" v:"required"`
	PostId   int `json:"post_id" v:"required"`
}
