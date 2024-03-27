package api

import "game-custom-com/internal/model/entity"

type PostImage struct {
	entity.Image
	Title string `json:"title"`
}

type SlideshowParams struct {
	Id  int    `json:"id"`
	Url string `json:"url"`
}
