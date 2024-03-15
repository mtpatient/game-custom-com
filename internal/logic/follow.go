package logic

import (
	"context"
	"game-custom-com/internal/dao"
	"game-custom-com/internal/service"
)

type sFollow struct {
}

func (s sFollow) IsFollow(ctx context.Context, id int) bool {
	uid := service.Context().Get(ctx).User.Id

	db := dao.Follow.Ctx(ctx)

	one, _ := db.Where("user_id", uid).Where("follow_user_id", id).One()

	if one != nil {
		return true
	}

	return false
}

func init() {
	service.RegisterFollow(&sFollow{})
}
