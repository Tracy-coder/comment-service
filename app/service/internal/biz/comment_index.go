package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type CommentIndexRepo interface {
	GetSubCommentIndex(ctx context.Context, root int64, limit int) ([]int64, error)
	GetComment(ctx context.Context, id int64) (*CommentObj, error)
	GetBatchComment(ctx context.Context, id []int64) ([]*CommentObj, error)
	ListRootCommentIndex(ctx context.Context, obj_id int64, page_size int32, page_id int32, order_by_floor bool) ([]int64, error)
}

type CommentIndexCache interface {
	ListRootCommentIndexCache(ctx context.Context, obj_id int64, page_size int32, page_id int32, order_by_floor bool) ([]int64, error)
}
type CommentIndexUsecase struct {
	repo  CommentIndexRepo
	cache CommentIndexCache
	log   *log.Helper
}

func NewCommentIndexUsecase(repo CommentIndexRepo, cache CommentIndexCache, logger log.Logger) *CommentIndexUsecase {
	return &CommentIndexUsecase{repo: repo, cache: cache, log: log.NewHelper(logger)}
}

func (uc *CommentIndexUsecase) GetSubCommentIndex(ctx context.Context, id int64) ([]int64, error) {
	return uc.repo.GetSubCommentIndex(ctx, id, 5)
}

func (uc *CommentIndexUsecase) GetComment(ctx context.Context, id int64) (*CommentObj, error) {
	return uc.repo.GetComment(ctx, id)
}
func (uc *CommentIndexUsecase) GetBatchComment(ctx context.Context, id []int64) ([]*CommentObj, error) {
	return uc.repo.GetBatchComment(ctx, id)
}
func (uc *CommentIndexUsecase) ListRootCommentIndex(ctx context.Context, obj_id int64, page_size int32, page_id int32, order_by_floor bool) ([]int64, error) {
	return uc.repo.ListRootCommentIndex(ctx, obj_id, page_size, page_id, order_by_floor)
}

func (uc *CommentIndexUsecase) ListRootCommentIndexCache(ctx context.Context, obj_id int64, page_size int32, page_id int32, order_by_floor bool) ([]int64, error) {
	return uc.cache.ListRootCommentIndexCache(ctx, obj_id, page_size, page_id, order_by_floor)
}

// func convert(comment *CommentObj) *v1.CommentObj {
// 	return &v1.CommentObj{
// 		ContentId: comment.ContentID,
// 		ObjId:     comment.ObjID,
// 		OwnerId:   comment.OwnerID,
// 		Root:      comment.Root,
// 		Parent:    comment.Parent,
// 		Floor:     comment.Floor,
// 		Like:      comment.Like,
// 		Hate:      comment.Hate,
// 		State:     comment.State,
// 		Message:   comment.Message,
// 		CreatedAt: timestamppb.New(comment.CreatedAt),
// 	}
// }
