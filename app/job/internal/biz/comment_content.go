package biz

import (
	"context"

	svcV1 "comment/api/comment/service/v1"

	"github.com/go-kratos/kratos/v2/log"
)

type CommentContentRepo interface {
	//InsertCommentContent(ctx context.Context, req *svcV1.CommentContent) error
	InsertCommentContent(ctx context.Context, req *svcV1.CommentMessage, floor int32) (int64, error)
}
type CommentContentCache interface {
	SetCommentMessageCache(ctx context.Context, id int64, message string) error
}
type CommentContentUsecase struct {
	repo  CommentContentRepo
	cache CommentContentCache
	log   *log.Helper
}

func NewCommentContentUsecase(repo CommentContentRepo, cache CommentContentCache, logger log.Logger) *CommentContentUsecase {
	return &CommentContentUsecase{repo: repo, cache: cache, log: log.NewHelper(logger)}
}

func (uc *CommentContentUsecase) InsertCommentContent(ctx context.Context, req *svcV1.CommentMessage, floor int32) (int64, error) {

	return uc.repo.InsertCommentContent(ctx, req, floor)
}

func (uc *CommentContentUsecase) SetCommentMessageCache(ctx context.Context, id int64, message string) error {
	return uc.cache.SetCommentMessageCache(ctx, id, message)
}
