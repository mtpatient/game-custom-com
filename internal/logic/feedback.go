package logic

import (
	"context"
	"game-custom-com/api"
	"game-custom-com/internal/dao"
	"game-custom-com/internal/model/do"
	"game-custom-com/internal/service"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

type sFeedback struct {
}

func (s sFeedback) GetList(ctx context.Context, params api.CommonParams) ([]api.FeedbackVo, int, error) {
	var res []api.FeedbackVo
	var total int

	db := dao.Feedback.Ctx(ctx).LeftJoin("user", "feedback.user_id=user.id").
		Fields("feedback.id, user_id, username, content, feedback.create_time").Order("id desc")
	if params.Keyword != "" {
		db = db.WhereLike("content", "%"+params.Keyword+"%").WhereOrLike("username", "%"+params.Keyword+"%")
	}

	err := db.Limit((params.PageIndex-1)*params.PageSize, params.PageSize).ScanAndCount(&res, &total, false)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, 0, gerror.New("获取失败")
	}

	imgDb := dao.Image.Ctx(ctx)
	for i := 0; i < len(res); i++ {
		array, err := imgDb.Fields("url").Where("feedback_id", res[i].Id).Array()
		if err != nil {
			g.Log().Error(ctx, err)
			return nil, 0, gerror.New("获取失败")
		}
		res[i].Images = toString(array)
	}
	//g.Log().Info(ctx, res)

	return res, total, nil
}

func (s sFeedback) Create(ctx context.Context, add api.FeedbackVo) error {
	err := dao.Feedback.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		id, err := tx.Ctx(ctx).InsertAndGetId("feedback", do.Feedback{
			Content: add.Content,
			UserId:  add.UserId,
		})
		if err != nil {
			return err
		}

		imgLen := len(add.Images)
		if imgLen > 0 {
			var images = make([]do.Image, imgLen)
			for i := 0; i < imgLen; i++ {
				images[i].Type = 2
				images[i].Url = add.Images[i]
				images[i].FeedbackId = id
			}

			_, err = tx.Ctx(ctx).Insert("image", images)
			if err != nil {
				return err
			}
		}
		return nil
	})

	if err != nil {
		g.Log().Error(ctx, err)
		return gerror.New("创建失败")
	}

	return nil
}

func init() {
	service.RegisterFeedback(&sFeedback{})
}
