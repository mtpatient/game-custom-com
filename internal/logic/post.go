package logic

import (
	"context"
	"game-custom-com/api"
	"game-custom-com/internal/dao"
	"game-custom-com/internal/model/do"
	"game-custom-com/internal/service"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
)

type sPost struct {
}

func init() {
	service.RegisterPost(&sPost{})
}

func (s sPost) Add(ctx context.Context, postAdd api.PostAdd) error {
	uid := service.Context().Get(ctx).User.Id

	if uid != postAdd.Post.UserId {
		return gerror.New("用户不一致！")
	}

	err := dao.Post.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		id, err := tx.Ctx(ctx).InsertAndGetId("post", postAdd.Post)
		if err != nil {
			return err
		}

		imgLen := len(postAdd.Images)
		if imgLen > 0 {
			var images = make([]do.Image, imgLen)
			for i := 0; i < imgLen; i++ {
				//tx.Ctx(ctx).Raw("INSERT INTO image(url, type, post_id, create_time) VALUE(?, ?, ?, ?)",
				//	img, 1, id, gtime.Now()).Insert()
				images[i].Type = 1
				images[i].Url = postAdd.Images[i]
				images[i].PostId = id
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
