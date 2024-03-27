package service

import (
	"context"
	"game-custom-com/api"
)

type IFeedback interface {
	Create(ctx context.Context, add api.FeedbackVo) error
	GetList(ctx context.Context, params api.CommonParams) ([]api.FeedbackVo, int, error)
}

var localFeedback IFeedback

func RegisterFeedback(i IFeedback) {
	localFeedback = i
}

func Feedback() IFeedback {
	if localFeedback == nil {
		panic("service feedback is not registered")
	}

	return localFeedback
}
