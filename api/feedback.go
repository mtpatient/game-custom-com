package api

import "game-custom-com/internal/model/entity"

type FeedbackAdd struct {
	Feedback entity.Feedback `json:"feedback"`
	Images   []string        `json:"images"`
}
