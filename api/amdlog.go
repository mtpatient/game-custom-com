package api

import "game-custom-com/internal/model/entity"

type AdmLgoVo struct {
	entity.AdminLog
	UserName string `json:"username"`
}
