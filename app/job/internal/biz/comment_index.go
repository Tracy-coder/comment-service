package biz

import (
	"context"

	svcV1 "comment/api/comment/service/v1"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/redis/go-redis/v9"
)

type CommentIndexRepo interface {
	//InsertCommentIndex(ctx context.Context, req *svcV1.CommentIndex) error
	InsertCommentIndex(ctx context.Context, req *svcV1.CommentMessage, floor int32, contentID int64) error
	DecrementCount(ctx context.Context, contentID int64) error
	IncrementCount(ctx context.Context, contentID int64) error
	IncrementFloor(ctx context.Context, contentID int64) error
	GetNextFloor(ctx context.Context, contentID int64) (int32, error)
	ListRootCommentIndex(ctx context.Context, obj_id int64, page_size int32, page_id int32, order_by_floor bool) ([]redis.Z, error)
}

type CommentIndexCache interface {
	SetCommentIndexCache(ctx context.Context, key string, z []redis.Z) error
}
type CommentIndexUsecase struct {
	repo  CommentIndexRepo
	cache CommentIndexCache
	log   *log.Helper
}

func NewCommentIndexUsecase(repo CommentIndexRepo, cache CommentIndexCache, logger log.Logger) *CommentIndexUsecase {
	return &CommentIndexUsecase{repo: repo, cache: cache, log: log.NewHelper(logger)}
}

func (uc *CommentIndexUsecase) InsertCommentIndex(ctx context.Context, req *svcV1.CommentMessage, floor int32, contentID int64) error {
	err := uc.repo.InsertCommentIndex(ctx, req, floor, contentID)
	return err
}

func (uc *CommentIndexUsecase) DecrementCount(ctx context.Context, contentID int64) error {
	return uc.repo.DecrementCount(ctx, contentID)
}

func (uc *CommentIndexUsecase) IncrementCount(ctx context.Context, contentID int64) error {
	return uc.repo.IncrementCount(ctx, contentID)
}

func (uc *CommentIndexUsecase) IncrementFloor(ctx context.Context, contentID int64) error {
	return uc.repo.IncrementFloor(ctx, contentID)
}

func (uc *CommentIndexUsecase) GetNextFloor(ctx context.Context, contentID int64) (int32, error) {
	return uc.repo.GetNextFloor(ctx, contentID)
}

func (uc *CommentIndexUsecase) SetCommentIndexCache(ctx context.Context, key string, z []redis.Z) error {
	return uc.cache.SetCommentIndexCache(ctx, key, z)
}

func (uc *CommentIndexUsecase) ListRootCommentIndex(ctx context.Context, obj_id int64, order_by_floor bool) ([]redis.Z, error) {
	return uc.repo.ListRootCommentIndex(ctx, obj_id, -1, 1, order_by_floor)
}
