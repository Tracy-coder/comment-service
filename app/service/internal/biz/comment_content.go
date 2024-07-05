package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type CommentContentRepo interface {
	GetCommentMessage(ctx context.Context, id int64) (string, error)
	GetBatchCommentMessage(ctx context.Context, id []int64) ([]*string, error)
}

type CommentContentCache interface {
	GetCommentMessageCache(ctx context.Context, id int64) (string, error)
}
type CommentContentUsecase struct {
	repo  CommentContentRepo
	cache CommentContentCache
	log   *log.Helper
}

func NewCommentContentUsecase(repo CommentContentRepo, cache CommentContentCache, logger log.Logger) *CommentContentUsecase {
	return &CommentContentUsecase{repo: repo, cache: cache, log: log.NewHelper(logger)}
}

func (uc *CommentContentUsecase) GetCommentMessage(ctx context.Context, id int64) (string, error) {
	return uc.repo.GetCommentMessage(ctx, id)
}

func (uc *CommentContentUsecase) GetBatchCommentMessage(ctx context.Context, ids []int64) ([]*string, error) {
	return uc.repo.GetBatchCommentMessage(ctx, ids)
}

func (uc *CommentContentUsecase) GetCommentMessageCache(ctx context.Context, id int64) (string, error) {
	return uc.cache.GetCommentMessageCache(ctx, id)
}
