package service

import "context"

type IAdmLog interface {
	Save(ctx context.Context, t string, msg string)
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
