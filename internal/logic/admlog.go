package logic

import (
	"context"
	"game-custom-com/internal/dao"
	"game-custom-com/internal/model/entity"
	"game-custom-com/internal/service"
)

type sAdmLog struct {
}

func init() {
	service.RegisterAdmLog(&sAdmLog{})
}

func (s sAdmLog) Save(ctx context.Context, t string, msg string) {
	aldb := dao.AdminLog.Ctx(ctx)
	uid := service.Context().Get(ctx).User.Id

	aldb.Insert(&entity.AdminLog{
		UserId:  uid,
		Type:    t,
		Content: msg,
	})
}
