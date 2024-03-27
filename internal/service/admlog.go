package service

import (
	"context"
	"game-custom-com/api"
)

type IAdmLog interface {
	Save(ctx context.Context, t string, msg string) error
	GetList(ctx context.Context, params api.CommonParams) ([]api.AdmLgoVo, int, error)
}

var localAdmLog IAdmLog

func AdmLog() IAdmLog {
	if localAdmLog == nil {
		panic("implement not found for interface ISection, forgot register?")
	}
	return localAdmLog
}

func RegisterAdmLog(i IAdmLog) {
	localAdmLog = i
}
