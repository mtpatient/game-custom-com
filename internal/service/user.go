package service

import (
	"context"
	"game-custom-com/api"
	"game-custom-com/internal/model/entity"
)

type IUser interface {
	Register(ctx context.Context, user api.User) error
	Login(ctx context.Context, user api.User) (entity.User, string, error)
	Logout(ctx context.Context, token string) error
	NameExist(ctx context.Context, username string) (bool, error)
	GetById(ctx context.Context, id int64) (entity.User, error)
	IsLogin(ctx context.Context) (bool, error)
	UserRole(ctx context.Context) (int64, error)
}

var localUser IUser

func User() IUser {
	if localUser == nil {
		panic("implement not found for interface IUser, forget register?")
	}

	return localUser
}

func RegisterUser(i IUser) {
	localUser = i
}
