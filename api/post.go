package api

import "game-custom-com/internal/model/entity"

type PostAdd struct {
	Post   entity.Post `json:"post"`
	Images []string    `json:"images"`
}
