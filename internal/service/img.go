package service

import (
	"context"
	"game-custom-com/internal/model/entity"
)

type IImg interface {
	GetSignatures(ctx context.Context, count int) ([]string, error)
	Save(ctx context.Context, images []entity.Image) error
	GetAllAvatar(ctx context.Context) ([]entity.Image, error)
	GetImageById(ctx context.Context, count int) (entity.Image, error)
	DeleteAvatar(ctx context.Context, id int) error
	Update(ctx context.Context, image entity.Image) error
}

var localImg IImg

func Img() IImg {
	if localImg == nil {
		panic("implement not found for interface IImg, forgot register?")
	}
	return localImg
}

func RegisterImg(i IImg) {
	localImg = i
}
