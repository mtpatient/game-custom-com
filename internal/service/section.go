package service

import (
	"context"
	"game-custom-com/internal/model/entity"
)

type ISection interface {
	Add(ctx context.Context, section entity.Section) error
	Update(ctx context.Context, section entity.Section) error
	GetAll(ctx context.Context) ([]entity.Section, error)
	GetById(ctx context.Context, id int) (entity.Section, error)
}

var localSection ISection

func Section() ISection {
	if localSection == nil {
		panic("implement not found for interface ISection, forgot register?")
	}
	return localSection
}

func RegisterSection(i ISection) {
	localSection = i
}
