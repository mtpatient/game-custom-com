package logic

import (
	"context"
	"encoding/json"
	"game-custom-com/internal/consts"
	"game-custom-com/internal/dao"
	"game-custom-com/internal/model/do"
	"game-custom-com/internal/model/entity"
	"game-custom-com/internal/service"
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
		return nil, err
	}

	return sections, nil
}

func (s sSection) Add(ctx context.Context, section entity.Section) error {
	sdb := dao.Section.Ctx(ctx)

	_, err := sdb.Insert(section)
	jsonData, _ := json.Marshal(section)
	if err != nil {
		service.AdmLog().Save(ctx, consts.LogError, "添加板块失败:"+string(jsonData))
		return err
	}
	service.AdmLog().Save(ctx, consts.LogSuccess, "添加板块成功:"+string(jsonData))
	return nil
}

func (s sSection) Update(ctx context.Context, section entity.Section) error {
	sdb := dao.Section.Ctx(ctx)

	_, err := sdb.Where("id", section.Id).Update(do.Section{
		Name: section.Name,
		Dc:   section.Dc,
		Icon: section.Icon,
	})
	jsonData, _ := json.Marshal(section)
	if err != nil {
		service.AdmLog().Save(ctx, consts.LogError, "更新板块失败："+string(jsonData))
		return err
	}
	service.AdmLog().Save(ctx, consts.LogSuccess, "更新板块成功："+string(jsonData))
	return nil
}
