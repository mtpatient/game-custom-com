package logic

import (
	"context"
	"game-custom-com/internal/service"
)

type sCollect struct {
}

func (s sCollect) Add(ctx context.Context) {

}

func init() {
	service.RegisterCollect(new(sCollect))
}
