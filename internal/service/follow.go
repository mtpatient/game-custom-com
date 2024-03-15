package service

import "context"

type IFollow interface {
	IsFollow(ctx context.Context, id int) bool
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
