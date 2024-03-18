package logic

import (
	"context"
	"game-custom-com/api"
	"game-custom-com/internal/dao"
	"game-custom-com/internal/service"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
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

func (s sFollow) GetFollowList(ctx context.Context, i int) ([]api.FollowVo, error) {
	var res []api.FollowVo
	user := service.Context().Get(ctx).User

	array, err := dao.Follow.Ctx(ctx).Where("user_id", i).Fields("follow_user_id").Array()
	if err != nil {
		goto exit
	}
	if err := dao.User.Ctx(ctx).LeftJoin("image", "image.id=user.avatar").
		Fields("user.id,username,signature,url as avatar").WhereIn("user.id", array).Scan(&res); err != nil {
		goto exit
	}

	if user != nil {
		uid := user.Id
		if uid == i {
			for i := range res {
				res[i].IsFollow = 1
			}
		} else {
			array, err = dao.Follow.Ctx(ctx).Where("user_id", uid).Fields("follow_user_id").Array()
			if err != nil {
				goto exit
			}
			for i := range res {
				if contains(array, res[i].Id) {
					res[i].IsFollow = 1
				}
			}
		}
	}

	return res, nil

exit:
	g.Log().Error(ctx, err)
	return nil, gerror.New("获取关注列表失败")
}

func (s sFollow) GetFansList(ctx context.Context, get api.FansGet) ([]api.FansVo, error) {
	var res []api.FansVo
	user := service.Context().Get(ctx).User

	array, err := dao.Follow.Ctx(ctx).Where("follow_user_id", get.Id).Fields("user_id").Array()
	if err != nil {
		goto exit
	}

	if err := dao.User.Ctx(ctx).LeftJoin("image", "image.id=user.avatar").
		Fields("user.id,username,signature,url as avatar").WhereIn("user.id", array).
		Limit(get.PageSize*(get.PageIndex-1), get.PageSize).Scan(&res); err != nil {
		goto exit
	}

	// 若当前请求已登录，则需要判断当前用户是否已关注
	if user != nil {
		array, err = dao.Follow.Ctx(ctx).Where("user_id", user.Id).Fields("follow_user_id").Array()
		if err != nil {
			goto exit
		}
		for i := range res {
			if contains(array, res[i].Id) {
				res[i].IsFollow = 1
			}
		}
	}

	return res, nil

exit:
	g.Log().Error(ctx, err)
	return nil, gerror.New("获取粉丝列表失败")
}

func init() {
	service.RegisterFollow(&sFollow{})
}

func contains(array []gdb.Value, i int) bool {
	for _, v := range array {
		if v.Int() == i {
			return true
		}
	}

	return false
}
