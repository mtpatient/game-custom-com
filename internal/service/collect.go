package service

import "context"

type ICollect interface {
	Add(ctx context.Context)
}

var localCollect ICollect

func RegisterCollect(i ICollect) {
	localCollect = i
}

func Collect() ICollect {
	if localCollect == nil {
		panic("implement not found for interface ICollect, forgot register?")
	}

	return localCollect
}
