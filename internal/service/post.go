package service

import (
	"context"
	"game-custom-com/api"
)

type IPost interface {
	Add(ctx context.Context, postAdd api.PostAdd) (int, error)
	GetById(ctx context.Context, id int) (api.PostDetail, error)
	Like(ctx context.Context, like api.PostLike) error
	Collect(ctx context.Context, collect api.PostCollect) error
	GetMinePost(ctx context.Context, get api.GetPostParams) ([]api.PostVo, error)
	Top(ctx context.Context, top api.TopPost) error
	Del(ctx context.Context, id int) error
	Update(ctx context.Context, update api.PostAdd) error
	GetTopPost(ctx context.Context, id int) ([]api.TopPostVo, error)
	GetFollow(ctx context.Context, get api.GetPostParams) ([]api.PostVo, error)
	GetPostList(ctx context.Context, get api.GetPostParams) ([]api.PostVo, error)
	SearchPost(ctx context.Context, get api.SearchParams) ([]api.PostVo, error)
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
