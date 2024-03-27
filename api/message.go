package api

import (
	"game-custom-com/internal/model/entity"
	"github.com/gogf/gf/v2/util/gmeta"
)

type Params struct {
	PageIndex int `json:"page_index" v:"required"`
	PageSize  int `json:"page_size" v:"required"`
}

type LikesMessageVo struct {
	gmeta.Meta `orm:"table:like"`
	entity.Like
	Comment  *entity.Comment `json:"comment" orm:"with:id=comment_id"`
	Post     *entity.Post    `json:"post" orm:"with:id=post_id"`
	UserName string          `json:"username"`
	Avatar   string          `json:"avatar"`
}
