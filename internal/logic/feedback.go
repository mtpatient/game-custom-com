package logic

import (
	"context"
	"game-custom-com/api"
	"game-custom-com/internal/dao"
	"game-custom-com/internal/model/do"
	"game-custom-com/internal/service"
	"github.com/gogf/gf/v2/database/gdb"
)

type sFeedback struct {
}

func (s sFeedback) Create(ctx context.Context, add api.FeedbackAdd) error {
	err := dao.Feedback.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		id, err := tx.Ctx(ctx).InsertAndGetId("feedback", add.Feedback)
		if err != nil {
			return err
		}

		imgLen := len(add.Images)
		if imgLen > 0 {
			var images = make([]do.Image, imgLen)
			for i := 0; i < imgLen; i++ {
				images[i].Type = 1
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
		return err
	}

	return nil
}

func init() {
	service.RegisterFeedback(&sFeedback{})
}
