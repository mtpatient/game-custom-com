package service

import (
	"context"
	"game-custom-com/api"
)

type IPost interface {
	Add(ctx context.Context, postAdd api.PostAdd) error
	GetById(ctx context.Context, id int) (api.PostRes, error)
}

var localPost IPost

func RegisterPost(i IPost) {
	localPost = i
}

func Post() IPost {
	if localPost == nil {
		panic("implement not found for interface ISection, forgot register?")
	}
	return localPost
}
