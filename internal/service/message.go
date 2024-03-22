package service

import (
	"context"
	"game-custom-com/api"
)

type IMessage interface {
	GetLikesMessage(context context.Context, get api.GetLikesParams) ([]api.LikesMessageVo, error)
	GetMessageNew(ctx context.Context) ([]int, error)
	Read(ctx context.Context, id int) error
}

var localMessage IMessage

func RegisterMessage(i IMessage) {
	localMessage = i
}

func Message() IMessage {
	return localMessage
}
