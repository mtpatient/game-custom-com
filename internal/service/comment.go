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
	GetMineComments(ctx context.Context, get api.CommentGet) ([]api.CommentRes, error)
	GetCommentById(ctx context.Context, i int) (api.PostCommentRes, error)
	CommentList(ctx context.Context, params api.CommonParams) ([]api.CommentRes, int, error)
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
