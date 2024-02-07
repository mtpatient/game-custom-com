package model

import (
	"game-custom-com/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
)

type Context struct {
	Data g.Map        // 自定义KV变量
	User *entity.User // 用户
}
