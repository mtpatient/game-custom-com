package main

import (
	"bufio"
	"context"
	"fmt"
	"game-custom-com/api"
	"game-custom-com/internal/dao"
	"game-custom-com/internal/model/entity"
	"game-custom-com/internal/service"
	"game-custom-com/utility"
	"github.com/go-ego/gse"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"os"
	"testing"
	"time"
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

func TestImgMap(t *testing.T) {
	ctx := context.Background()
	var users []api.FollowUserVo
	array, err := dao.Follow.Ctx(ctx).Where("user_id", 14).Fields("follow_user_id").Array()
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}
	err = g.DB().Model("user").LeftJoin("image", "image.id = user.avatar").
		Fields("user.id, user.username, user.signature, image.url as avatar").WhereIn("user.id", array).Scan(&users)
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}
	for i := 0; i < len(users); i++ {
		fmt.Println(users[i])
	}
}

func TestStringsDifferent(t *testing.T) {
	arr1 := []string{"1", "2", "3", "5", "6"}
	arr2 := []string{"4", "5", "6", "7", "8"}

	diff1, diff2 := stringsDifference(arr1, arr2)

	fmt.Println("diff1", diff1)
	fmt.Println("diff2", diff2)
}

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

func TestSubString(t *testing.T) {
	url := "https://game-custom-1312933264.cos.ap-guangzhou.myqcloud.com/img/2024-03-18/x584fc0wxk0czwl9zyfzmjg10032g4lv"

	str := gstr.ReplaceI(url, "https://game-custom-1312933264.cos.ap-guangzhou.myqcloud.com/", "")
	fmt.Println(str)
}

func TestCosDel(t *testing.T) {
	keys := []string{"https://game-custom-1312933264.cos.ap-guangzhou.myqcloud.com/img/2024-02-29/1a0j290cws0czharmrhtqi01002mo278",
		"https://game-custom-1312933264.cos.ap-guangzhou.myqcloud.com/img/2024-02-18/1a0j290irg0cz87m3lpj594740d6cl59",
		"https://game-custom-1312933264.cos.ap-guangzhou.myqcloud.com/img/2024-02-18/1a0j290irg0cz889zchpdzg79042sfu2",
		"https://game-custom-1312933264.cos.ap-guangzhou.myqcloud.com/img/2024-02-18/1a0j290irg0cz87np6she5075016476f"}

	err := utility.CosDel(context.Background(), keys)
	if err != nil {
		return
	}
}

func TestSeGo(t *testing.T) {
	str := "比安,hhhhhalskdj,卡,空中花园，白毛索拉卡觉得，继续加长，看看效果怎么样，这不是哈可以吗？为什么之前就卡死了"

	var seg gse.Segmenter

	err := seg.LoadDict()
	if err != nil {
		return
	}
	//segments := seg.Cut(str, true)
	res := seg.Cut(str, true)
	//fmt.Println(segments)
	fmt.Println(res)

	//seg.AddToken("比安卡", 100, "nrt")
	//seg.AddToken("花嫁", 100, "n")
	//seg.AddToken("七实", 100, "n")
	//segments := seg.ModeSegment([]byte(str), true)
	//fmt.Println(seg)
	//strs := make([]string, 0)
	//filer := []string{"n", "nr", "v", "vn", "nrt"}
	//for _, segment := range segments {
	//	word := segment.Token().Text()
	//	pos := segment.Token().Pos()
	//	for _, f := range filer {
	//		if pos == f {
	//			strs = append(strs, word)
	//		}
	//	}
	//	fmt.Println(word, pos)
	//}

	//fmt.Println(strs)
}

func TestWith(t *testing.T) {
	ctx := context.Background()
	var res []api.LikesMessageVo
	dao.Like.Ctx(ctx).Fields("like.id,`user_id`,`like`.`post_id`,`comment_id`,`praise_id`,`like`.create_time, username, url as avatar").
		LeftJoin("user", "user.id = like.user_id").
		LeftJoin("image", "image.id = user.avatar").
		With(entity.Comment{}, entity.Post{}).Where("like.id", 1).Scan(&res)

	g.Log().Info(ctx, res)
}

func TestPub(t *testing.T) {
	ctx := context.Background()

	for i := 0; i < 100; i++ {
		publish, err := g.Redis().Publish(ctx, "test", "test"+string(rune(i)))
		if err != nil {
			return
		}
		fmt.Println(publish)
		time.Sleep(time.Second)
	}
}

func TestSub(t *testing.T) {
	ctx := context.Background()

	subscribe, _, err := g.Redis().Subscribe(ctx, "test")
	if err != nil {
		return
	}
	for {
		msg, err := subscribe.ReceiveMessage(ctx)
		if err != nil {
			g.Log().Error(ctx, err)
		}
		g.Log().Info(ctx, msg)
	}

}
