package service

import (
	"context"
	"game-custom-com/api"
	"game-custom-com/internal/model/entity"
)

type IUser interface {
	Register(ctx context.Context, user api.User) error
	Login(ctx context.Context, user api.User) (api.UserRes, string, error)
	Logout(ctx context.Context, token string) error
	NameExist(ctx context.Context, username string) (bool, error)
	GetById(ctx context.Context, id int) (api.UserRes, error)
	IsLogin(ctx context.Context) (bool, error)
	UserRole(ctx context.Context) (int, error)
	Update(ctx context.Context, user entity.User) error
	ReplacePassword(ctx context.Context, rp api.UserReplacePassword) error
	GetAuthCode(ctx context.Context, str string) error
	ResetPwd(ctx context.Context, rs api.ResetPwd) error
	Follow(ctx context.Context, follow api.UserFollow) error
	SearchUser(ctx context.Context, get api.UserSearchParams) ([]api.FollowUserVo, error)
	GetUserList(ctx context.Context, get api.CommonParams) ([]entity.User, int, error)
	Ban(ctx context.Context, ban api.Ban) error
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
