package logic

import "game-custom-com/internal/service"

type sMessage struct {
}

func init() {
	service.RegisterMessage(new(sMessage))
}
