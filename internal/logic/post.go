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

func (s sPost) GetById(ctx context.Context, id int) (api.PostDetail, error) {
	var postDetail api.PostDetail
	postDb := dao.Post.Ctx(ctx)
	err := postDb.Where("id", id).Scan(&postDetail.Post)
	if err != nil {
		g.Log().Error(ctx, err)
		return postDetail, gerror.New("帖子不存在或状态异常")
	}
	// 判断帖子是否存在且为公共可见状态
	if postDetail.Post.Status != 0 {
		return postDetail, gerror.New("帖子不存在或状态异常")
	}
	_, err = postDb.Where("id", id).OmitEmpty().Update(g.Map{
		"view_count": gdb.Raw("view_count+1"),
	})
	if err != nil {
		g.Log().Error(ctx, err)
		return postDetail, gerror.New("帖子不存在或状态异常")
	}
	postDetail.Post.ViewCount++

	//postId := strconv.Itoa(id)
	// 获取帖子点赞数、收藏数
	//strLen, err := g.Redis().SCard(ctx, consts.PostLikesKey+postId)
	//if err != nil {
	//	return postDetail, err
	//}
	//postDetail.Post.LikeCount = int(strLen)
	//strLen, err = g.Redis().SCard(ctx, consts.PostCollectKey+postId)
	//if err != nil {
	//	return postDetail, err
	//}
	//postDetail.Post.CollectCount = int(strLen)

	// 获取评论数
	if postDetail.CommentCount, err = dao.Comment.Ctx(ctx).Where("post_id", id).Count(); err != nil {
		g.Log().Error(ctx, err)
		return postDetail, gerror.New("帖子不存在或状态异常")
	}

	if user := service.Context().Get(ctx).User; user != nil {
		uid := user.Id
		// 判断是否点赞
		//member, err := g.Redis().SIsMember(ctx, consts.PostLikesKey+postId, uid)
		//if err != nil {
		//	return postDetail, err
		//}
		//if member > 0 {
		//	postDetail.IsLike = 1
		//}
		likeDb := dao.Like.Ctx(ctx)
		one, _ := likeDb.Where("user_id", uid).Where("post_id", id).One()
		if one != nil {
			postDetail.IsLike = 1
		}
		// 判断是否收藏
		//member, err = g.Redis().SIsMember(ctx, consts.PostCollectKey+postId, uid)
		//if err != nil {
		//	return postDetail, err
		//}
		//if member > 0 {
		//	postDetail.IsCollect = 1
		//}
		collectDb := dao.Collect.Ctx(ctx)
		one, _ = collectDb.Where("user_id", uid).Where("post_id", id).One()
		if one != nil {
			postDetail.IsCollect = 1
		}
		// 判断是否关注
		followDb := dao.Follow.Ctx(ctx)
		one, err = followDb.Where("user_id", uid).Where("follow_user_id", postDetail.Post.UserId).One()
		if one != nil {
			postDetail.IsFollow = 1
		}
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
			g.Log().Error(ctx, err)
			return gerror.New("添加帖子失败")
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
				g.Log().Error(ctx, err)
				return gerror.New("添加图片失败")
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
		//_, err := g.Redis().SAdd(ctx, consts.PostLikesKey+strconv.Itoa(like.PostId), uid)
		//if err != nil {
		//	return err
		//}

		// 更新数据库
		err := dao.Like.Transaction(context.Background(), func(ctx context.Context, tx gdb.TX) error {
			// 判断是否已经点赞
			one, err := tx.Ctx(ctx).GetOne("SELECT * FROM `like` WHERE user_id=? AND post_id=? AND praise_id=? and delete_time is null LIMIT 1",
				uid, like.PostId, like.ToUserId)

			if one != nil {
				return gerror.New("已经点赞了")
			}

			_, err = tx.Ctx(ctx).Insert("like", g.Map{
				"user_id":   uid,
				"praise_id": like.ToUserId,
				"post_id":   like.PostId,
			})
			if err != nil {
				g.Log().Error(ctx, err)
				return gerror.New("点赞失败")
			}

			_, err = tx.Ctx(ctx).Update("post", g.Map{
				"like_count": gdb.Raw("like_count+1"),
			}, g.Map{
				"id": like.PostId,
			})

			if err != nil {
				g.Log().Error(ctx, err)
				return gerror.New("更新帖子失败")
			}

			_, err = tx.Ctx(ctx).Update("user", g.Map{
				"like_count": gdb.Raw("like_count+1"),
			}, g.Map{
				"id": like.ToUserId,
			})
			if err != nil {
				g.Log().Error(ctx, err)
				return gerror.New("更新用户失败")
			}

			_, err = tx.Ctx(ctx).Insert("message", g.Map{
				"user_id":    uid,
				"receive_id": like.ToUserId,
				"type":       2,
				"post_id":    like.PostId,
				"content":    "点赞了你的帖子",
				"is_read":    0,
			})
			if err != nil {
				g.Log().Error(ctx, err)
				return gerror.New("添加消息失败")
			}
			return nil
		})
		if err != nil {
			return err
		}
	} else if like.Operate == 2 {
		//_, err := g.Redis().SRem(ctx, consts.PostLikesKey+strconv.Itoa(like.PostId), uid)
		//if err != nil {
		//	return err
		//}

		//更新数据库
		err := dao.Like.Transaction(context.Background(), func(ctx context.Context, tx gdb.TX) error {
			// 判断是否未点赞
			one, err := tx.Ctx(ctx).GetOne("SELECT * FROM `like` WHERE user_id=? AND post_id=? AND praise_id=? and delete_time is null LIMIT 1",
				uid, like.PostId, like.ToUserId)

			if one == nil {
				return gerror.New("未点赞！")
			}

			_, err = tx.Ctx(ctx).Delete("like", g.Map{
				"user_id":   uid,
				"praise_id": like.ToUserId,
				"post_id":   like.PostId,
			})
			if err != nil {
				g.Log().Error(ctx, err)
				return gerror.New("删除失败")
			}

			_, err = tx.Ctx(ctx).Update("post", g.Map{
				"like_count": gdb.Raw("like_count-1"),
			}, g.Map{
				"id": like.PostId,
			})
			if err != nil {
				g.Log().Error(ctx, err)
				return gerror.New("更新帖子失败")
			}

			_, err = tx.Ctx(ctx).Update("user", g.Map{
				"like_count": gdb.Raw("like_count-1"),
			}, g.Map{
				"id": like.ToUserId,
			})
			if err != nil {
				g.Log().Error(ctx, err)
				return gerror.New("更新用户失败")
			}

			_, err = tx.Ctx(ctx).Delete("message", g.Map{
				"user_id":    uid,
				"receive_id": like.ToUserId,
				"type":       2,
				"post_id":    like.PostId,
			})
			if err != nil {
				g.Log().Error(ctx, err)
				return gerror.New("删除消息失败")
			}
			return nil
		})
		if err != nil {
			return err
		}
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
		//_, err := g.Redis().SAdd(ctx, consts.PostCollectKey+strconv.Itoa(collect.PostId), uid)
		//if err != nil {
		//	return err
		//}
		err := dao.Post.Transaction(context.Background(), func(ctx context.Context, tx gdb.TX) error {
			// 判断是否已收藏
			one, err := tx.GetOne("SELECT * FROM `collect` WHERE post_id=? AND user_id=? and delete_time is null limit 1", collect.PostId, uid)

			if one != nil {
				return gerror.New("已收藏！")
			}

			// 更新帖子收藏数
			_, err = tx.Update("post", g.Map{
				"collect_count": gdb.Raw("collect_count+1"),
			}, g.Map{
				"id": collect.PostId,
			})
			if err != nil {
				g.Log().Error(ctx, err)
				return gerror.New("更新帖子失败")
			}

			// 更新用户收藏表
			_, err = tx.Insert("collect", g.Map{
				"post_id": collect.PostId,
				"user_id": uid,
			})
			if err != nil {
				g.Log().Error(ctx, err)
				return gerror.New("更新用户失败")
			}
			return nil
		})
		if err != nil {
			return err
		}
	} else if collect.Operate == 2 {
		// 取消收藏
		//_, err := g.Redis().SRem(ctx, consts.PostCollectKey+strconv.Itoa(collect.PostId), uid)
		//if err != nil {
		//	return err
		//}
		err := dao.Post.Transaction(context.Background(), func(ctx context.Context, tx gdb.TX) error {
			// 判断是否已收藏
			one, err := tx.GetOne("SELECT * FROM `collect` WHERE post_id=? AND user_id=? and delete_time is null limit 1", collect.PostId, uid)
			if one == nil {
				return gerror.New("未收藏！")
			}

			// 更新帖子收藏数
			_, err = tx.Update("post", g.Map{
				"collect_count": gdb.Raw("collect_count-1"),
			}, g.Map{
				"id": collect.PostId,
			})
			if err != nil {
				g.Log().Error(ctx, err)
				return gerror.New("更新帖子失败")
			}
			// 删除用户收藏表
			_, err = tx.Delete("collect", g.Map{
				"post_id": collect.PostId,
				"user_id": uid,
			})
			if err != nil {
				g.Log().Error(ctx, err)
				return gerror.New("删除用户失败")
			}

			return nil
		})
		if err != nil {
			return err
		}
	} else {
		return gerror.New("操作类型错误")
	}
	return nil
}
