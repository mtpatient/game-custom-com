package service

import (
	"context"
	"game-custom-com/api"
	"game-custom-com/internal/model/entity"
)

type IMessage interface {
	GetLikesMessage(context context.Context, get api.Params) ([]api.LikesMessageVo, error)
	GetMessageNew(ctx context.Context) ([]int, error)
	Read(ctx context.Context, id int) error
	Publish(ctx context.Context, id int, message string) error
	Add(ctx context.Context, message entity.Message) error
	GetNotice(ctx context.Context, params api.Params) ([]entity.Message, error)
}

var localMessage IMessage

func RegisterMessage(i IMessage) {
	localMessage = i
}

func Message() IMessage {
	return localMessage
}
