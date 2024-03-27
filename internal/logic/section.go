package logic

import (
	"context"
	"encoding/json"
	"game-custom-com/internal/dao"
	"game-custom-com/internal/model/do"
	"game-custom-com/internal/model/entity"
	"game-custom-com/internal/service"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

type sSection struct {
}

func init() {
	service.RegisterSection(&sSection{})
}
func (s sSection) GetById(ctx context.Context, id int) (entity.Section, error) {
	db := dao.Section.Ctx(ctx)

	var section entity.Section
	err := db.Where("id", id).Scan(&section)

	if err != nil {
		return section, err
	}

	return section, nil
}

func (s sSection) GetAll(ctx context.Context) ([]entity.Section, error) {
	db := dao.Section.Ctx(ctx)
	var sections []entity.Section
	err := db.Scan(&sections)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, gerror.New("获取板块失败！")
	}

	return sections, nil
}

func (s sSection) Add(ctx context.Context, section entity.Section) error {
	sdb := dao.Section.Ctx(ctx)

	_, err := sdb.Insert(section)
	jsonData, _ := json.Marshal(section)
	if err != nil {
		return gerror.New("添加板块失败")
	}
	service.AdmLog().Save(ctx, "板块", "添加板块成功:"+string(jsonData))
	return nil
}

func (s sSection) Update(ctx context.Context, section entity.Section) error {
	sdb := dao.Section.Ctx(ctx)

	_, err := sdb.Where("id", section.Id).Update(do.Section{
		Name: section.Name,
		Dc:   section.Dc,
		Icon: section.Icon,
		Role: section.Role,
	})
	jsonData, _ := json.Marshal(section)
	if err != nil {
		return gerror.New("更新板块失败")
	}
	service.AdmLog().Save(ctx, "板块", "更新板块成功："+string(jsonData))
	return nil
}
