package logic

import (
	"context"
	"encoding/json"
	"game-custom-com/api"
	"game-custom-com/internal/dao"
	"game-custom-com/internal/model/do"
	"game-custom-com/internal/model/entity"
	"game-custom-com/internal/service"
	"game-custom-com/utility"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/guid"
	"strconv"
	"time"
)

type sImg struct {
}

func init() {
	service.RegisterImg(&sImg{})
}

func (s sImg) UpdateSlideshow(ctx context.Context, params api.SlideshowParams) error {
	err := dao.Image.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		one, err := dao.Post.Ctx(ctx).Where("id", params.Id).Where("status", 0).One()
		if err != nil {
			g.Log().Error(ctx, err)
			return gerror.New("查询帖子失败！")
		}
		if one == nil {
			return gerror.New("帖子不存在！")
		}

		one, err = dao.Image.Ctx(ctx).Where("post_id", params.Id).Where("type", 3).One()
		if err != nil {
			g.Log().Error(ctx, err)
			return gerror.New("查询帖子失败！")
		}
		if one != nil {
			return gerror.New("该帖子已关联其它轮播图！")
		}

		if _, err := dao.Image.Ctx(ctx).OmitEmpty().Update(g.Map{
			"url":     params.Url,
			"post_id": params.Id,
			"type":    3,
		}); err != nil {
			g.Log().Error(ctx, err)
			return gerror.New("更新图片失败！")
		}

		if err := service.AdmLog().Save(ctx, "图片", "更新轮播图："+gconv.String(params)); err != nil {
			g.Log().Error(ctx, err)
			return gerror.New("保存管理日志失败！")
		}

		return nil
	})
	if err != nil {
		g.Log().Error(ctx, err)
		return err
	}

	return nil
}

func (s sImg) SaveSlideshow(ctx context.Context, params api.SlideshowParams) error {
	err := dao.Image.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		count, err := dao.Image.Ctx(ctx).Where("type", 3).Count()
		if err != nil {
			g.Log().Error(ctx, err)
			return gerror.New("查询图片失败！")
		}
		if count > 6 {
			return gerror.New("轮播图最多只能有7张！")
		}

		one, err := dao.Post.Ctx(ctx).Where("id", params.Id).Where("status", 0).One()
		if err != nil {
			g.Log().Error(ctx, err)
			return gerror.New("查询帖子失败！")
		}
		if one == nil {
			return gerror.New("帖子不存在！")
		}

		one, err = dao.Image.Ctx(ctx).Where("post_id", params.Id).Where("type", 3).One()
		if err != nil {
			g.Log().Error(ctx, err)
			return gerror.New("查询帖子失败！")
		}
		if one != nil {
			return gerror.New("该帖子已关联其它轮播图！")
		}

		if _, err := dao.Image.Ctx(ctx).Insert(g.Map{
			"url":     params.Url,
			"post_id": params.Id,
			"type":    3,
		}); err != nil {
			g.Log().Error(ctx, err)
			return gerror.New("保存图片失败！")
		}

		if err := service.AdmLog().Save(ctx, "图片", "保存轮播图："+gconv.String(params)); err != nil {
			g.Log().Error(ctx, err)
			return gerror.New("保存管理日志失败！")
		}

		return nil
	})
	if err != nil {
		g.Log().Error(ctx, err)
		return err
	}

	return nil
}

func (s sImg) GetSlideshow(ctx context.Context) ([]api.PostImage, error) {
	var res []api.PostImage

	if err := dao.Image.Ctx(ctx).LeftJoin("post", "image.post_id=post.id").
		Fields("image.id, url, name, post_id, image.create_time, title").
		Where("type", 3).Order("id desc").Scan(&res); err != nil {
		g.Log().Error(ctx, err)
		return nil, gerror.New("获取轮播图失败！")
	}

	return res, nil
}

func (s sImg) Del(ctx context.Context, i int) error {
	if err := dao.Image.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		var img entity.Image
		if err := dao.Image.Ctx(ctx).Where("id", i).Scan(&img); err != nil {
			g.Log().Error(ctx, err)
			return gerror.New("获取图片失败！")
		}

		if err := utility.CosDel(ctx, []string{img.Url}); err != nil {
			return gerror.New("删除COS对象失败！")
		}

		if _, err := dao.Image.Ctx(ctx).Where("id", i).Delete(); err != nil {
			g.Log().Error(ctx, err)
			return gerror.New("删除图片失败！")
		}

		if err := service.AdmLog().Save(ctx, "图片", "删除图片："+gconv.String(img)); err != nil {
			return err
		}

		return nil
	}); err != nil {
		g.Log().Error(ctx, err)
		return gerror.New("删除图片失败！")
	}

	return nil
}

func (s sImg) PostImgList(ctx context.Context, params api.CommonParams) ([]api.PostImage, int, error) {
	var res []api.PostImage
	var total int
	db := dao.Image.Ctx(ctx).LeftJoin("post", "image.post_id=post.id").
		Fields("image.id, url, name, post_id, image.create_time, title").
		Where("type", 1).Order("id desc")
	if params.Keyword != "" {
		db = db.WhereLike("title", "%"+params.Keyword+"%")
	}
	err := db.Limit((params.PageIndex-1)*params.PageSize, params.PageSize).
		ScanAndCount(&res, &total, false)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, 0, gerror.New("获取图片列表失败！")
	}

	return res, total, nil
}

func (s sImg) Update(ctx context.Context, image entity.Image) error {
	db := dao.Image.Ctx(ctx)

	_, err := db.Where("id", image.Id).Update(do.Image{
		Url:  image.Url,
		Name: image.Name,
	})
	jsonData, _ := json.Marshal(image)
	if err != nil {
		g.Log().Error(ctx, err)
		return gerror.New("更新图片失败！")
	}

	if err := service.AdmLog().Save(ctx, "图片", "更新图片成功！"+string(jsonData)); err != nil {
		return err
	}
	return nil
}

func (s sImg) DeleteAvatar(ctx context.Context, id int) error {
	Udb := dao.User.Ctx(ctx)
	one, _ := Udb.One("avatar", id)
	if one.IsEmpty() == false {
		return gerror.New("该头像已有用户使用！")
	}

	Idb := dao.Image.Ctx(ctx)

	_, err := Idb.Delete("id", id)
	if err != nil {
		g.Log().Error(ctx, err)
		return gerror.New("删除头像失败！")
	}

	if err := service.AdmLog().Save(ctx, "头像", "删除头像成功："+strconv.Itoa(id)); err != nil {
		return err
	}
	return nil
}

func (s sImg) GetSignatures(ctx context.Context, count int) ([]string, error) {
	date := time.Now().String()[:10]

	names := make([]string, count)

	for i := 0; i < count; i++ {
		names[i] = "img/" + date + "/" + guid.S()
	}

	return utility.GetCOSSignature(ctx, names, time.Minute)
}

func (s sImg) Save(ctx context.Context, images []entity.Image) error {
	err := dao.Image.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		db := dao.Image.Ctx(ctx)

		_, err := db.Insert(images)
		jsonData, _ := json.Marshal(images)
		if err != nil {
			g.Log().Error(ctx, err)
			return gerror.New("保存图片失败！")
		}

		if err := service.AdmLog().Save(ctx, "图片", "保存图片成功！:"+string(jsonData)); err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		g.Log().Error(ctx, err)
		return gerror.New("保存图片失败！")
	}

	return nil
}

func (s sImg) GetAllAvatar(ctx context.Context) ([]entity.Image, error) {
	db := dao.Image.Ctx(ctx)
	var images []entity.Image

	err := db.Where("type", 0).Scan(&images)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, gerror.New("获取头像失败！")
	}

	return images, nil
}

func (s sImg) GetImageById(ctx context.Context, id int) (entity.Image, error) {
	db := dao.Image.Ctx(ctx)
	var image entity.Image

	err := db.Where("id", id).Scan(&image)
	if err != nil {
		g.Log().Error(ctx, err)
		return image, gerror.New("获取图片失败！")
	}

	return image, nil
}
