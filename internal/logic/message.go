package logic

import (
	"context"
	"game-custom-com/api"
	"game-custom-com/internal/dao"
	"game-custom-com/internal/model/entity"
	"game-custom-com/internal/service"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

type sMessage struct {
}

func (s sMessage) Read(ctx context.Context, id int) error {
	uid := service.Context().Get(ctx).User.Id

	_, err := dao.Message.Ctx(ctx).Fields("is_read").Where("receive_id", uid).
		Where("is_read", 0).Where("type", id).Update(g.Map{
		"is_read": 1,
	})
	if err != nil {
		g.Log().Error(ctx, err)
		return gerror.New("标记消息为已读失败！")
	}

	return nil
}

func (s sMessage) GetMessageNew(ctx context.Context) ([]int, error) {
	uid := service.Context().Get(ctx).User.Id
	res := []int{0, 0, 0}
	db := dao.Message.Ctx(ctx).Where("receive_id", uid).Where("is_read", 0)

	for i := range res {
		count, err := db.Count("type", i)
		if err != nil {
			g.Log().Error(ctx, err)
			return nil, gerror.New("获取消息数失败")
		}
		res[i] = count
	}

	return res, nil
}

func (s sMessage) GetLikesMessage(ctx context.Context, get api.GetLikesParams) ([]api.LikesMessageVo, error) {
	var res []api.LikesMessageVo

	uid := service.Context().Get(ctx).User.Id

	err := dao.Like.Ctx(ctx).Fields("like.id,`user_id`,`like`.`post_id`,`comment_id`,`praise_id`,`like`.create_time, username, url as avatar").
		LeftJoin("user", "user.id = like.user_id").
		LeftJoin("image", "image.id = user.avatar").Order("create_time desc").
		With(entity.Comment{}, entity.Post{}).Where("praise_id", uid).WhereNot("like.user_id", uid).
		Limit((get.PageIndex-1)*get.PageSize, get.PageSize).Scan(&res)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, gerror.New("获取点赞消息失败")
	}

	return res, nil
}

func init() {
	service.RegisterMessage(new(sMessage))
}
