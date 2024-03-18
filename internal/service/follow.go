package service

import (
	"context"
	"game-custom-com/api"
)

type IFollow interface {
	IsFollow(ctx context.Context, id int) bool
	GetFollowList(ctx context.Context, i int) ([]api.FollowVo, error)
	GetFansList(ctx context.Context, get api.FansGet) ([]api.FansVo, error)
}

var localFollow IFollow

func RegisterFollow(i IFollow) {
	localFollow = i
}

func Follow() IFollow {
	if localFollow == nil {
		panic("implement not found for interface IFollow, forgot register?")
	}

	return localFollow
}
