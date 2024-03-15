package main

import (
	"bufio"
	"context"
	"fmt"
	"game-custom-com/api"
	"game-custom-com/internal/dao"
	"game-custom-com/internal/model/entity"
	"game-custom-com/internal/service"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"os"
	"testing"
)

func TestGetSign(t *testing.T) {
	signature, err := service.Img().GetSignatures(context.Background(), 2)
	if err != nil {
		return
	}
	fmt.Println(signature)
}

func TestSaveImages(t *testing.T) {
	images := make([]entity.Image, 2)
	images[0] = entity.Image{
		Type: 0,
		Url:  "test1",
	}

	images[1] = entity.Image{
		Type: 1,
		Url:  "test2",
	}

	service.Img().Save(context.Background(), images)
}

func TestGetAvatar(t *testing.T) {
	avatar, err := service.Img().GetAllAvatar(context.Background())
	if err != nil {
		return
	}
	fmt.Println(avatar)
}

func TestUpdate(t *testing.T) {
	err := service.Img().Update(context.Background(), entity.Image{
		Type: 0,
		Id:   7,
		Url:  "https://game-custom-1312933264.cos.ap-guangzhou.myqcloud.com/img/2024-02-18/1a0j290irg0cz87m3lpj594740d6cl59",
		Name: "test",
	})
	if err != nil {
		g.Log(err.Error())
	}
	return
}

func TestFloor(t *testing.T) {
	dao.Comment.Transaction(context.Background(), func(c context.Context, tx gdb.TX) error {
		one, err := tx.Ctx(c).GetOne("SELECT IFNULL(MAX(floor),0)+1 as floor FROM `comment` WHERE post_id = 1")
		if err != nil {
			return err
		}
		floor, ok := one.Map()["floor"].(int64)
		g.Log().Info(c, one.Map()["floor"], floor, ok)

		return nil
	})
}

func TestStruct(t *testing.T) {
	var comment api.PostCommentRes

	g.Log().Info(context.Background(), comment)
}

func TestWithSelect(t *testing.T) {
	var comment []api.PostCommentRes
	ctx := context.Background()
	err := dao.Comment.Ctx(ctx).Where("post_id", 13).Order("create_time desc").
		Limit(0, 2).Scan(&comment)
	if err != nil {
		g.Log().Info(context.Background(), err)
		return
	}
	for i := 0; i < len(comment); i++ {
		v, err := dao.User.Ctx(ctx).Fields("username").Where("id", comment[i].UserId).Value()
		if err != nil {
			return
		}
		comment[i].UserName = v.String()

		v, err = dao.User.Ctx(ctx).Fields("username").Where("id", comment[i].ReplyId).Value()
		if err != nil {
			return
		}
		comment[i].ReplyName = v.String()

		v, err = dao.User.Ctx(ctx).Fields("avatar").Where("id", comment[i].UserId).Value()
		if err != nil {
			return
		}
		v, err = dao.Image.Ctx(ctx).Fields("url").Where("id", v).Value()
		if err != nil {
			return
		}
		comment[i].Avatar = v.String()

		uid := 17
		count, err := dao.Like.Ctx(ctx).Where("comment_id", comment[i].Id).Where("user_id", uid).Count()
		if err != nil {
			return
		}
		if count > 0 {
			comment[i].IsLike = 1
		}
		if uid == comment[i].UserId {
			comment[i].IsOwn = 1
		}

		err = dao.Comment.Ctx(ctx).Where("parent_id", comment[i].Id).Scan(&comment[i].Children)
		if err != nil {
			return
		}

		for j := 0; j < len(comment[i].Children); j++ {
			v, err = dao.User.Ctx(ctx).Fields("username").Where("id", comment[i].Children[j].UserId).Value()
			if err != nil {
				return
			}
			comment[i].Children[j].UserName = v.String()

			v, err = dao.User.Ctx(ctx).Fields("avatar").Where("id", comment[i].Children[j].UserId).Value()
			if err != nil {
				return
			}
			v, err = dao.Image.Ctx(ctx).Fields("url").Where("id", v).Value()
			if err != nil {
				return
			}
			comment[i].Children[j].Avatar = v.String()

			uid := 17
			count, err = dao.Like.Ctx(ctx).Where("comment_id", comment[i].Children[j].Id).Where("user_id", uid).Count()
			if err != nil {
				return
			}
			if count > 0 {
				comment[i].Children[j].IsLike = 1
			}
			if uid == comment[i].Children[j].UserId {
				comment[i].Children[j].IsOwn = 1
			}
		}
	}
	file, err := os.OpenFile("output.json", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			g.Log().Error(ctx, err)
			return
		}
	}(file)

	res := fmt.Sprintf("%+v", comment)
	write := bufio.NewWriter(file)
	_, err = write.WriteString(res)
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}

	fmt.Println(comment)
}
