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

type sComment struct {
}

func (s sComment) Del(ctx context.Context, i int) error {
	uid := service.Context().Get(ctx).User.Id
	role := service.Context().Get(ctx).User.Role
	db := dao.Comment.Ctx(ctx).Where("id", i)
	//messageDb := dao.Message.Ctx(ctx).Where("comment_id", i)

	err := dao.Message.Ctx(ctx).Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		if role == 1 {
			// 管理员
			//id, err := db.Fields("user_id").Value()
			//if err != nil {
			//	g.Log().Error(ctx, err)
			//	return gerror.New("删除评论失败")
			//}
			//_, err = messageDb.Delete("user_id", id)
			//if err != nil {
			//	g.Log().Error(ctx, err)
			//	return gerror.New("删除评论失败")
			//}
			_, err := db.Delete()
			if err != nil {
				g.Log().Error(ctx, err)
				return gerror.New("删除评论失败")
			}
		} else {
			// 普通用户
			if uid == 0 {
				return gerror.New("请登录")
			}
			result, err := db.Where("user_id", uid).Delete()
			if err != nil {
				g.Log().Error(ctx, err)
				return gerror.New("删除评论失败")
			}
			if result == nil {
				return gerror.New("没有权限")
			}
			//_, err = messageDb.Where("user_id", uid).Delete()
			//if err != nil {
			//	g.Log().Error(ctx, err)
			//	return gerror.New("删除消息失败")
			//}
		}
		return nil
	})
	if err != nil {
		g.Log().Error(ctx, err)
		return gerror.New("删除评论失败")
	}

	return nil
}

func (s sComment) GetPostCommentList(ctx context.Context, get api.PostCommentGet) ([]api.PostCommentRes, error) {
	var res []api.PostCommentRes
	db := dao.Comment.Ctx(ctx).Where("post_id", get.PostId).WhereNot("status", 1).Where("parent_id", 0)
	if get.IsOnlyPublisher == 1 {
		value, err := dao.Post.Ctx(ctx).Fields("user_id").Where("id", get.PostId).Value()
		if err != nil {
			g.Log().Error(ctx, err)
			return nil, gerror.New("获取楼主评论列表失败")
		}

		db = db.Where("user_id", value)
	}
	if get.ShowType == 1 {
		db = db.Order("like_count desc")
	} else if get.ShowType == 2 {
		db = db.Order("create_time desc")
	}

	err := db.Order("id asc").Limit(get.PageSize*(get.PageIndex-1), get.PageSize).Scan(&res)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, gerror.New("获取评论列表失败")
	}

	v := service.Context().Get(ctx).User
	if v != nil {
		err := handleComments(ctx, res, v.Id)
		if err != nil {
			g.Log().Error(ctx, err)
			return nil, gerror.New("获取评论列表失败")
		}
	} else {
		err := handleComments(ctx, res, 0)
		if err != nil {
			g.Log().Error(ctx, err)
			return nil, gerror.New("获取评论列表失败")
		}
	}

	return res, nil
}

func (s sComment) Add(ctx context.Context, add api.CommentAdd) error {
	err := dao.Comment.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		uid := service.Context().Get(ctx).User.Id
		var floor int64 = 0
		if add.IsFloor {
			one, err := tx.Ctx(ctx).GetOne("SELECT IFNULL(MAX(floor),0)+1 as floor FROM `comment` WHERE post_id=?", add.PostId)
			if err != nil {
				return err
			}
			floor, _ = one.Map()["floor"].(int64)
		}
		id, err := tx.Ctx(ctx).InsertAndGetId("comment", g.Map{
			"post_id":    add.PostId,
			"user_id":    uid,
			"content":    add.Content,
			"floor":      floor,
			"reply_id":   add.ToUserId,
			"comment_id": add.CommentId,
			"parent_id":  add.ParentId,
			"like_count": 0,
			"status":     0,
		})
		if err != nil {
			return err
		}

		if _, err = tx.Ctx(ctx).Insert("message", g.Map{
			"user_id":    uid,
			"receive_id": add.ToUserId,
			"type":       1,
			"content":    add.Content,
			"post_id":    add.PostId,
			"comment_id": id,
			"is_read":    0,
		}); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		g.Log().Error(ctx, err)
		return gerror.New("评论失败")
	}

	return nil
}

func (s sComment) Like(ctx context.Context, like api.CommentLike) error {
	uid := service.Context().Get(ctx).User.Id

	err := dao.Like.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		if like.Operate == 1 {
			count, err := dao.Like.Ctx(ctx).Where(g.Map{
				"comment_id": like.Id,
				"user_id":    uid,
				"praise_id":  like.ToUserId,
				"post_id":    like.PostId,
			}).Count()
			if err != nil {
				return err
			}

			if count > 0 {
				return gerror.New("已点赞")
			}

			if _, err := tx.Ctx(ctx).Insert("like", g.Map{
				"comment_id": like.Id,
				"user_id":    uid,
				"praise_id":  like.ToUserId,
				"post_id":    like.PostId,
			}); err != nil {
				return err
			}

			if _, err := dao.Comment.Ctx(ctx).Where("id", like.Id).OmitEmpty().Update(g.Map{
				"like_count": gdb.Raw("like_count+1"),
			}); err != nil {
				return err
			}

			if _, err := dao.User.Ctx(ctx).Where("id", like.ToUserId).OmitEmpty().Update(g.Map{
				"like_count": gdb.Raw("like_count+1"),
			}); err != nil {
				return err
			}

			if _, err := dao.Message.Ctx(ctx).Insert(g.Map{
				"user_id":    uid,
				"receive_id": like.ToUserId,
				"type":       2,
				"content":    "点赞了你的评论",
				"comment_id": like.Id,
				"post_id":    like.PostId,
				"is_read":    0,
			}); err != nil {
				return err
			}
		} else if like.Operate == 2 {
			count, err := dao.Like.Ctx(ctx).Where(g.Map{
				"comment_id": like.Id,
				"user_id":    uid,
				"praise_id":  like.ToUserId,
				"post_id":    like.PostId,
			}).Count()
			if err != nil {
				return err
			}
			if count == 0 {
				return gerror.New("未点赞")
			}

			if _, err := tx.Ctx(ctx).Delete("like", g.Map{
				"comment_id": like.Id,
				"user_id":    uid,
				"praise_id":  like.ToUserId,
			}); err != nil {
				return err
			}

			if _, err := dao.Comment.Ctx(ctx).Where("id", like.Id).OmitEmpty().Update(g.Map{
				"like_count": gdb.Raw("like_count-1"),
			}); err != nil {
				return err
			}

			if _, err := dao.User.Ctx(ctx).Where("id", like.ToUserId).OmitEmpty().Update(g.Map{
				"like_count": gdb.Raw("like_count-1"),
			}); err != nil {
				return err
			}

			if _, err := dao.Message.Ctx(ctx).Delete(g.Map{
				"user_id":    uid,
				"receive_id": like.ToUserId,
				"type":       2,
				"comment_id": like.Id,
				"post_id":    like.PostId,
			}); err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		g.Log().Error(ctx, err)
		return gerror.New("点赞失败")
	}

	return nil
}

func init() {
	service.RegisterComment(new(sComment))
}

func handleComments(ctx context.Context, comment []api.PostCommentRes, uid int) error {
	userDb := dao.User.Ctx(ctx)
	likeDb := dao.Like.Ctx(ctx)
	imgDb := dao.Image.Ctx(ctx)
	commentDb := dao.Comment.Ctx(ctx)
	for i := 0; i < len(comment); i++ {
		v, err := userDb.Fields("username").Where("id", comment[i].UserId).Value()
		if err != nil {
			return err
		}
		comment[i].UserName = v.String()

		v, err = userDb.Fields("username").Where("id", comment[i].ReplyId).Value()
		if err != nil {
			return err
		}
		comment[i].ReplyName = v.String()

		v, err = userDb.Fields("avatar").Where("id", comment[i].UserId).Value()
		if err != nil {
			return err
		}
		v, err = imgDb.Fields("url").Where("id", v).Value()
		if err != nil {
			return err
		}
		comment[i].Avatar = v.String()

		if uid > 0 {
			count, err := likeDb.Where("comment_id", comment[i].Id).Where("user_id", uid).Count()
			if err != nil {
				return err
			}
			if count > 0 {
				comment[i].IsLike = 1
			}
			if uid == comment[i].UserId {
				comment[i].IsOwn = 1
			}
		}

		err = commentDb.Where("parent_id", comment[i].Id).Scan(&comment[i].Children)
		if err != nil {
			return err
		}

		for j := 0; j < len(comment[i].Children); j++ {
			v, err = userDb.Fields("username").Where("id", comment[i].Children[j].UserId).Value()
			if err != nil {
				return err
			}
			comment[i].Children[j].UserName = v.String()

			v, err = userDb.Fields("username").Where("id", comment[i].Children[j].ReplyId).Value()
			if err != nil {
				return err
			}
			comment[i].Children[j].ReplyName = v.String()

			v, err = userDb.Fields("avatar").Where("id", comment[i].Children[j].UserId).Value()
			if err != nil {
				return err
			}
			v, err = imgDb.Fields("url").Where("id", v).Value()
			if err != nil {
				return err
			}
			comment[i].Children[j].Avatar = v.String()

			if uid > 0 {
				count, err := likeDb.Where("comment_id", comment[i].Children[j].Id).Where("user_id", uid).Count()
				if err != nil {
					return err
				}
				if count > 0 {
					comment[i].Children[j].IsLike = 1
				}
				if uid == comment[i].Children[j].UserId {
					comment[i].Children[j].IsOwn = 1
				}
			}
		}
	}

	return nil
}
