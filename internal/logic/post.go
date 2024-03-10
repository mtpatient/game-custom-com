package logic

import (
	"context"
	"game-custom-com/api"
	"game-custom-com/internal/consts"
	"game-custom-com/internal/dao"
	"game-custom-com/internal/model/do"
	"game-custom-com/internal/service"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"strconv"
)

type sPost struct {
}

func init() {
	service.RegisterPost(&sPost{})
}

func (s sPost) GetById(ctx context.Context, id int) (api.PostDetail, error) {
	var postDetail api.PostDetail
	postDb := dao.Post.Ctx(ctx)
	err := postDb.Where("id", id).Scan(&postDetail.Post)
	if err != nil {
		return postDetail, err
	}
	// 判断帖子是否存在且为公共可见状态
	if postDetail.Post.Status != 0 {
		return postDetail, gerror.New("帖子不存在或状态异常")
	}
	_, err = postDb.Where("id", id).OmitEmpty().Update(g.Map{
		"view_count": gdb.Raw("view_count+1"),
	})
	if err != nil {
		return postDetail, err
	}
	postDetail.Post.ViewCount++

	postId := strconv.Itoa(id)
	// 获取帖子点赞数、收藏数
	strLen, err := g.Redis().SCard(ctx, consts.PostLikesKey+postId)
	if err != nil {
		return postDetail, err
	}
	postDetail.Post.LikeCount = int(strLen)
	strLen, err = g.Redis().SCard(ctx, consts.PostCollectKey+postId)
	if err != nil {
		return postDetail, err
	}
	postDetail.Post.CollectCount = int(strLen)

	if user := service.Context().Get(ctx).User; user != nil {
		uid := user.Id
		// 判断是否点赞
		member, err := g.Redis().SIsMember(ctx, consts.PostLikesKey+postId, uid)
		if err != nil {
			return postDetail, err
		}
		if member > 0 {
			postDetail.IsLike = 1
		}
		// 判断是否收藏
		member, err = g.Redis().SIsMember(ctx, consts.PostCollectKey+postId, uid)
		if err != nil {
			return postDetail, err
		}
		if member > 0 {
			postDetail.IsCollect = 1
		}
		// 判断是否关注
		followDb := dao.Follow.Ctx(ctx)
		one, err := followDb.Where("user_id", uid).Where("follow_user_id", postDetail.Post.UserId).One()
		if one != nil {
			postDetail.IsFollow = 1
		}
	}

	// TODO 获取评论
	commDb := dao.Comment.Ctx(ctx)
	err = commDb.Where("post_id", id).Order("create_time").Scan(&postDetail.Comments)
	if err != nil {
		return postDetail, nil
	}

	return postDetail, nil
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

func (s sPost) Like(ctx context.Context, like api.PostLike) error {
	// 判断帖子是否存在且为公共可见状态
	if like.Status != 0 {
		return gerror.New("帖子不存在或状态异常")
	}

	uid := service.Context().Get(ctx).User.Id

	// 将点赞信息存入redis
	if like.Operate == 1 {
		_, err := g.Redis().SAdd(ctx, consts.PostLikesKey+strconv.Itoa(like.PostId), uid)
		if err != nil {
			return err
		}

		// 更新数据库
		//go dao.Like.Transaction(context.Background(), func(ctx context.Context, tx gdb.TX) error {
		//	_, err := tx.Ctx(ctx).Update("post", g.Map{
		//		"like_count": gdb.Raw("like_count+1"),
		//	}, g.Map{
		//		"id": like.PostId,
		//	})
		//
		//	if err != nil {
		//		return err
		//	}
		//
		//	_, err = tx.Ctx(ctx).Insert("like", g.Map{
		//		"user_id":   uid,
		//		"praise_id": like.ToUserId,
		//		"post_id":   like.PostId,
		//	})
		//	if err != nil {
		//		return err
		//	}
		//	return nil
		//})
	} else if like.Operate == 2 {
		_, err := g.Redis().SRem(ctx, consts.PostLikesKey+strconv.Itoa(like.PostId), uid)
		if err != nil {
			return err
		}

		// 更新数据库
		//go dao.Like.Transaction(context.Background(), func(ctx context.Context, tx gdb.TX) error {
		//	_, err := tx.Ctx(ctx).Update("post", g.Map{
		//		"like_count": gdb.Raw("like_count-1"),
		//	}, g.Map{
		//		"id": like.PostId,
		//	})
		//
		//	if err != nil {
		//		return err
		//	}
		//
		//	_, err = tx.Ctx(ctx).Delete("like", g.Map{
		//		"user_id":   uid,
		//		"praise_id": like.ToUserId,
		//		"post_id":   like.PostId,
		//	})
		//	if err != nil {
		//		return err
		//	}
		//	return nil
		//})
	} else {
		return gerror.New("操作类型错误")
	}

	return nil
}

func (s sPost) Collect(ctx context.Context, collect api.PostCollect) error {
	// 判断帖子是否存在且为公共可见状态

	if collect.Status != 0 {
		return gerror.New("帖子不存在或状态异常")
	}

	uid := service.Context().Get(ctx).User.Id

	if collect.Operate == 1 {
		// 将收藏信息存入redis
		_, err := g.Redis().SAdd(ctx, consts.PostCollectKey+strconv.Itoa(collect.PostId), uid)
		if err != nil {
			return err
		}

		//go dao.Post.Transaction(context.Background(), func(ctx context.Context, tx gdb.TX) error {
		//	// 更新帖子收藏数
		//	_, err := tx.Update("post", g.Map{
		//		"collect_count": gdb.Raw("collect_count+1"),
		//	}, g.Map{
		//		"id": collect.PostId,
		//	})
		//	if err != nil {
		//		return err
		//	}
		//
		//	// 更新用户收藏表
		//	_, err = tx.Insert("collect", g.Map{
		//		"post_id": collect.PostId,
		//		"user_id": uid,
		//	})
		//	if err != nil {
		//		return err
		//	}
		//	return nil
		//})
	} else if collect.Operate == 2 {
		// 取消收藏
		_, err := g.Redis().SRem(ctx, consts.PostCollectKey+strconv.Itoa(collect.PostId), uid)
		if err != nil {
			return err
		}
		//go dao.Post.Transaction(context.Background(), func(ctx context.Context, tx gdb.TX) error {
		//	// 更新帖子收藏数
		//	_, err := tx.Update("post", g.Map{
		//		"collect_count": gdb.Raw("collect_count-1"),
		//	}, g.Map{
		//		"id": collect.PostId,
		//	})
		//	if err != nil {
		//		return err
		//	}
		//	// 删除用户收藏表
		//	_, err = tx.Delete("collect", g.Map{
		//		"post_id": collect.PostId,
		//		"user_id": uid,
		//	})
		//	if err != nil {
		//		return err
		//	}
		//
		//	return nil
		//})
	} else {
		return gerror.New("操作类型错误")
	}
	return nil
}
