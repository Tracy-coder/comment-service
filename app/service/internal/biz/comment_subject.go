package biz

import (
	v1 "comment/api/comment/service/v1"
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type CommentSubjectRepo interface {
	GetSubject(ctx context.Context, req *v1.GetSubjectRequest) (*v1.GetSubjectReply, error)
}

type CommentSubjectUsecase struct {
	repo CommentSubjectRepo
	log  *log.Helper
}

func NewCommentSubjectUsecase(repo CommentSubjectRepo, logger log.Logger) *CommentSubjectUsecase {
	return &CommentSubjectUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *CommentSubjectUsecase) GetSubject(ctx context.Context, req *v1.GetSubjectRequest) (*v1.GetSubjectReply, error) {
	return uc.repo.GetSubject(ctx, req)
}
