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

type sPost struct {
}

func init() {
	service.RegisterPost(&sPost{})
}

func (s sPost) GetById(ctx context.Context, id int) (api.PostRes, error) {
	var postRes api.PostRes

	postDb := dao.Post.Ctx(ctx)
	err := postDb.Where("id", id).Scan(&postRes.Post)
	if err != nil {
		return postRes, err
	}
	_, err = postDb.Where("id", id).OmitEmpty().Update(g.Map{
		"view_count": gdb.Raw("view_count+1"),
	})
	if err != nil {
		return postRes, err
	}

	commDb := dao.Comment.Ctx(ctx)
	err = commDb.Where("post_id", id).Order("create_time").Scan(&postRes.Comments)
	if err != nil {
		return postRes, nil
	}

	return postRes, nil
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
