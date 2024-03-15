package logic

import (
	"game-custom-com/internal/service"
)

type sLike struct {
}

func init() {
	service.RegisterLike(new(sLike))
}
