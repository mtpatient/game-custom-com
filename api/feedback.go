package api

import "game-custom-com/internal/model/entity"

type FeedbackVo struct {
	entity.Feedback
	Images   []string `json:"images"`
	Username string   `json:"username"`
}
