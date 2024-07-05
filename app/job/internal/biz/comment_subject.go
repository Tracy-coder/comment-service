package biz

import (
	"context"

	svcV1 "comment/api/comment/service/v1"

	"github.com/go-kratos/kratos/v2/log"
)

type CommentSubjectRepo interface {
	InsertCommentSubject(ctx context.Context, req *svcV1.CommentSubject) error
	GetNextFloor(ctx context.Context, obj_id int64) (int32, error)
	IncrementFloor(ctx context.Context, obj_id int64) error
	IncrementCount(ctx context.Context, obj_id int64) error
	DecrementCount(ctx context.Context, obj_id int64) error
}

type CommentSubjectUsecase struct {
	repo CommentSubjectRepo
	log  *log.Helper
}

func NewCommentSubjectUsecase(repo CommentSubjectRepo, logger log.Logger) *CommentSubjectUsecase {
	return &CommentSubjectUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *CommentSubjectUsecase) InsertCommentSubject(ctx context.Context, req *svcV1.CommentSubject) error {
	return uc.repo.InsertCommentSubject(ctx, req)
}

func (uc *CommentSubjectUsecase) GetNextFloor(ctx context.Context, obj_id int64) (int32, error) {
	return uc.repo.GetNextFloor(ctx, obj_id)
}

func (uc *CommentSubjectUsecase) IncrementFloor(ctx context.Context, obj_id int64) error {
	return uc.repo.IncrementFloor(ctx, obj_id)
}

func (uc *CommentSubjectUsecase) IncrementCount(ctx context.Context, obj_id int64) error {
	return uc.repo.IncrementCount(ctx, obj_id)
}

func (uc *CommentSubjectUsecase) DecrementCount(ctx context.Context, obj_id int64) error {
	return uc.repo.DecrementCount(ctx, obj_id)
}
