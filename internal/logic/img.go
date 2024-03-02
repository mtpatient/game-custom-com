package logic

import (
	"context"
	"encoding/json"
	"game-custom-com/internal/consts"
	"game-custom-com/internal/dao"
	"game-custom-com/internal/model/do"
	"game-custom-com/internal/model/entity"
	"game-custom-com/internal/service"
	"game-custom-com/utility"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/guid"
	"strconv"
	"time"
)

type sImg struct {
}

func init() {
	service.RegisterImg(&sImg{})
}

func (s sImg) Update(ctx context.Context, image entity.Image) error {
	db := dao.Image.Ctx(ctx)

	_, err := db.Where("id", image.Id).Update(do.Image{
		Url:  image.Url,
		Name: image.Name,
	})
	jsonData, _ := json.Marshal(image)
	if err != nil {
		service.AdmLog().Save(ctx, consts.LogError, "更新图片失败！"+string(jsonData))
		return err
	}

	service.AdmLog().Save(ctx, consts.LogSuccess, "更新图片成功！"+string(jsonData))
	return nil
}

func (s sImg) DeleteAvatar(ctx context.Context, id uint) error {
	Udb := dao.User.Ctx(ctx)
	one, _ := Udb.One("img", id)
	if one.IsEmpty() == false {
		service.AdmLog().Save(ctx, consts.LogError, "头像已有用户使用！无法删除："+strconv.Itoa(int(id)))
		return gerror.New("该头像已有用户使用！")
	}

	Idb := dao.Image.Ctx(ctx)

	_, err := Idb.Delete("id", id)
	if err != nil {
		service.AdmLog().Save(ctx, consts.LogError, "删除头像失败："+strconv.Itoa(int(id)))
		return err
	}

	service.AdmLog().Save(ctx, consts.LogSuccess, "删除头像成功："+strconv.Itoa(int(id)))
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
	db := dao.Image.Ctx(ctx)

	_, err := db.Insert(images)
	jsonData, _ := json.Marshal(images)
	if err != nil {
		service.AdmLog().Save(ctx, consts.LogError, "保存图片失败！:"+string(jsonData))
		return err
	}
	service.AdmLog().Save(ctx, consts.LogSuccess, "保存图片成功！:"+string(jsonData))
	return nil
}

func (s sImg) GetAllAvatar(ctx context.Context) ([]entity.Image, error) {
	db := dao.Image.Ctx(ctx)
	var images []entity.Image

	err := db.Where("type", 0).Scan(&images)
	if err != nil {
		return nil, err
	}

	return images, nil
}

func (s sImg) GetImageById(ctx context.Context, count uint) (entity.Image, error) {
	db := dao.Image.Ctx(ctx)
	var image entity.Image

	err := db.Where("id", count).Scan(&image)
	if err != nil {
		return image, err
	}

	return image, nil
}
