package logic

import (
	"context"
	"game-custom-com/api"
	"game-custom-com/internal/dao"
	"game-custom-com/internal/model/do"
	"game-custom-com/internal/service"
	"game-custom-com/utility"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

type sPost struct {
}

func init() {
	service.RegisterPost(&sPost{})
}

func (s sPost) Update(ctx context.Context, update api.PostAdd) error {
	err := dao.Post.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		_, err := dao.Post.Ctx(ctx).Where("id", update.Post.Id).
			Fields("title,content,section,status").
			Update(g.Map{
				"title":   update.Post.Title,
				"content": update.Post.Content,
				"section": update.Post.Section,
				"status":  update.Post.Status,
			})
		if err != nil {
			g.Log().Error(ctx, err)
			return gerror.New("更新帖子失败！")
		}

		imgDb := dao.Image.Ctx(ctx)

		array, err := imgDb.Fields("url").Where("post_id", update.Post.Id).Array()
		if err != nil {
			g.Log().Error(ctx, err)
			return gerror.New("获取图片失败")
		}

		old := toString(array)
		// 取旧图片集和新图片集的差集，旧图片的差集则为要删除的图片，新图片的差集则为要添加的图片
		delList, insertList := stringsDifference(old, update.Images)
		// 删除图片
		if len(delList) > 0 {
			// 数据库删除
			_, err = imgDb.WhereIn("url", delList).Delete()
			if err != nil {
				g.Log().Error(ctx, err)
				return gerror.New("删除图片失败！")
			}
			// 从COS删除
			err := utility.CosDel(ctx, delList)
			if err != nil {
				g.Log().Error(ctx, err)
				return gerror.New("COS删除对象失败！")
			}
		}
		// 更新帖子图片
		if len(insertList) > 0 {
			images := make([]do.Image, len(insertList))
			for i := range insertList {
				images[i].Type = 1
				images[i].Url = insertList[i]
				images[i].PostId = update.Post.Id
			}

			_, err := imgDb.Insert(images)
			if err != nil {
				g.Log().Error(ctx, err)
				return gerror.New("更新帖子图片失败！")
			}
		}

		return nil
	})
	if err != nil {
		g.Log().Error(ctx, err)
		return gerror.New("更新帖子失败！")
	}

	return nil
}

func (s sPost) Top(ctx context.Context, top api.TopPost) error {
	user := service.Context().Get(ctx).User

	db := dao.Post.Ctx(ctx)
	switch top.Operate {
	case 1:
		// 查询是否已置顶其他帖子
		count, err := db.WhereNot("id", top.Id).Where("user_id", user.Id).Where("top_self", 1).Count()
		if err != nil {
			g.Log().Error(ctx, err)
			return gerror.New("帖子置顶失败！")
		}
		if count > 0 {
			return gerror.New("帖子置顶失败！请先取消其他帖子置顶！")
		}
		_, err = db.Where("id", top.Id).Where("user_id", user.Id).OmitEmpty().Update(g.Map{
			"top_self": 1,
		})
		if err != nil {
			g.Log().Error(ctx, err)
			return gerror.New("帖子置顶失败！")
		}
	case 2:
		_, err := db.Where("id", top.Id).Where("user_id", user.Id).Fields("top_self").Update(g.Map{
			"top_self": 0,
		})
		if err != nil {
			g.Log().Error(ctx, err)
			return gerror.New("帖子取消置顶失败！")
		}
	case 3:
		// 查询是否已置顶其他帖子
		count, err := db.WhereNot("id", top.Id).Where("top_section", 1).Count()
		if err != nil {
			g.Log().Error(ctx, err)
			return gerror.New("帖子置顶失败！")
		}
		if count > 0 {
			return gerror.New("帖子置顶失败！请先取消其他帖子置顶！")
		}

		if user.Role == 1 {
			_, err := db.Where("id", top.Id).OmitEmpty().Update(g.Map{
				"top_section": 1,
			})
			if err != nil {
				g.Log().Error(ctx, err)
				return gerror.New("帖子板块置顶失败！")
			}
		} else {
			return gerror.New("权限不足！")
		}
	case 4:
		if user.Role == 1 {
			_, err := db.Where("id", top.Id).Fields("top_section").Update(g.Map{
				"top_section": 0,
			})
			if err != nil {
				g.Log().Error(ctx, err)
				return gerror.New("帖子板块取消置顶失败！")
			}
		} else {
			return gerror.New("权限不足！")
		}
	default:
		return gerror.New("操作类型错误")
	}

	return nil
}

func (s sPost) Del(ctx context.Context, id int) error {
	uid := service.Context().Get(ctx).User.Id

	_, err := dao.Post.Ctx(ctx).Where("id", id).Where("user_id", uid).Delete()
	if err != nil {
		g.Log().Error(ctx, err)
		return gerror.New("删除失败")
	}

	return nil
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
	if postDetail.Post.Status != 0 && postDetail.Post.Status != 2 {
		return postDetail, gerror.New("帖子不存在或状态异常")
	}
	_, err = postDb.Where("id", id).FieldsEx("update_time").OmitEmpty().Update(g.Map{
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

func (s sPost) Add(ctx context.Context, postAdd api.PostAdd) (int, error) {
	uid := service.Context().Get(ctx).User.Id
	pid := 0
	if uid != postAdd.Post.UserId {
		return pid, gerror.New("用户不一致！")
	}

	err := dao.Post.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		id, err := tx.Ctx(ctx).InsertAndGetId("post", postAdd.Post)
		//g.Log().Info(ctx, "tx-id:", id)
		if err != nil {
			g.Log().Error(ctx, err)
			return gerror.New("添加帖子失败")
		}
		pid = int(id)

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
		g.Log().Error(ctx, err)
		return 0, gerror.New("添加帖子失败")
	}

	//g.Log().Info(ctx, "res-id:", pid)
	return pid, nil
}

func (s sPost) Like(ctx context.Context, like api.PostLike) error {
	// 判断帖子是否存在且为公共可见状态
	if like.Status != 0 && like.Status != 2 {
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

	if collect.Status != 0 && collect.Status != 2 {
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

func (s sPost) GetMinePost(ctx context.Context, get api.GetMinePost) ([]api.PostVo, error) {
	var res []api.PostVo
	db := dao.Post.Ctx(ctx).LeftJoin("user", "post.user_id=user.id").
		InnerJoin("image", "user.avatar = image.id").
		Fields("post.id,title,user_id,content,section,username,view_count,post.like_count,collect_count,top_self," +
			"top_section,post.status,post.create_time, url as avatar")

	if get.ShowType == 1 {
		if user := service.Context().Get(ctx).User; user != nil && user.Id == get.Id {
			db = db.WhereIn("post.status", g.Array{0, 2})
		} else {
			db = db.Where("post.status", 0)
		}
		db = db.Where("user_id", get.Id).Order("top_self desc").Order("create_time desc")
	} else if get.ShowType == 2 {
		array, err := dao.Collect.Ctx(ctx).Fields("post_id").Where("user_id", get.Id).Array()
		if err != nil {
			g.Log().Error(ctx, err)
			return nil, gerror.New("获取收藏帖子失败")
		}
		db = db.Where("post.status", 0).WhereIn("post.id", array).Order("create_time desc")
	} else {
		return nil, gerror.New("参数错误")
	}

	err := db.Limit(get.PageSize*(get.PageIndex-1), get.PageSize).Scan(&res)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, gerror.New("获取帖子列表失败")
	}

	// 评论数、图片列表、判断是否点赞
	commentDb := dao.Comment.Ctx(ctx)
	imgDb := dao.Image.Ctx(ctx)
	var likes []gdb.Value
	user := service.Context().Get(ctx).User
	if user != nil {
		array, err := dao.Like.Ctx(ctx).Fields("post_id").Where("user_id", user.Id).WhereNull("comment_id").Array()
		if err != nil {
			g.Log().Error(ctx, err)
			return nil, gerror.New("获取点赞数失败")
		}
		likes = array
	}
	for i := range res {
		count, err := commentDb.Where("post_id", res[i].Id).Count()
		if err != nil {
			g.Log().Error(ctx, err)
			return nil, gerror.New("获取评论数失败")
		}
		res[i].CommentCount = count

		array, err := imgDb.Fields("url").Where("post_id", res[i].Id).Array()
		if err != nil {
			g.Log().Error(ctx, err)
			return nil, gerror.New("获取图片列表失败")
		}

		res[i].ImgList = toString(array)

		if likes != nil {
			if contains(likes, res[i].Id) {
				res[i].IsLike = 1
			}
		}
	}

	return res, nil
}

func toString(list []gdb.Value) []string {
	var res = make([]string, len(list))
	for i := range list {
		res[i] = list[i].String()
	}

	return res
}

func toInt(list []gdb.Value) []int {
	var res = make([]int, len(list))
	for i := range list {
		res[i] = list[i].Int()
	}

	return res
}

// 获取两个字符串数组的差集
func stringsDifference(arr1, arr2 []string) (diff1, diff2 []string) {
	// 创建两个map分别记录arr1和arr2中的元素
	countMap1 := make(map[string]bool)
	countMap2 := make(map[string]bool)
	for _, val := range arr1 {
		countMap1[val] = true
	}
	for _, val := range arr2 {
		countMap2[val] = true
	}

	// 遍历arr1，如果元素在map2中不存在，则加入差集1
	for _, val := range arr1 {
		if !countMap2[val] {
			diff1 = append(diff1, val)
		}
	}

	// 遍历arr2，如果元素在map1中不存在，则加入差集2
	for _, val := range arr2 {
		if !countMap1[val] {
			diff2 = append(diff2, val)
		}
	}

	return diff1, diff2
}
