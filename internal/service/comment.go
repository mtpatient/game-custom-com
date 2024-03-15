package service

import (
	"context"
	"game-custom-com/api"
)

type IComment interface {
	Add(ctx context.Context, add api.CommentAdd) error
	GetPostCommentList(ctx context.Context, get api.PostCommentGet) ([]api.PostCommentRes, error)
	Del(ctx context.Context, i int) error
	Like(ctx context.Context, like api.CommentLike) error
}

var localComment IComment

func RegisterComment(i IComment) {
	localComment = i
}

func Comment() IComment {
	if localComment == nil {
		panic("Comment not Register!")
	}

	return localComment
}
