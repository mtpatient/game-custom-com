package logic

import (
	"context"
	"game-custom-com/api"
	"game-custom-com/internal/dao"
	"game-custom-com/internal/model/do"
	"game-custom-com/internal/service"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

type sAdmLog struct {
}

func (s sAdmLog) GetList(ctx context.Context, params api.CommonParams) ([]api.AdmLgoVo, int, error) {
	var res []api.AdmLgoVo
	var total int

	db := dao.AdminLog.Ctx(ctx).LeftJoin("user", "user.id = admin_log.user_id").
		Fields("admin_log.id, user_id, username, admin_log.type, admin_log.create_time, content").Order("create_time DESC")

	if params.Keyword != "" {
		db = db.WhereLike("content", "%"+params.Keyword+"%").WhereOrLike("username", "%"+params.Keyword+"%")
	}

	err := db.Limit((params.PageIndex-1)*params.PageSize, params.PageSize).ScanAndCount(&res, &total, false)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, 0, gerror.New("获取失败")
	}

	return res, total, nil
}

func (s sAdmLog) Save(ctx context.Context, t string, msg string) error {
	aldb := dao.AdminLog.Ctx(ctx)
	uid := service.Context().Get(ctx).User.Id

	aldb.Insert(&do.AdminLog{
		UserId:  uid,
		Type:    t,
		Content: msg,
	})
	return nil
}

func init() {
	service.RegisterAdmLog(&sAdmLog{})
}
