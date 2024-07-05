package service

import (
	pb "comment/api/comment/job/v1"
	"comment/app/job/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type CommentJobService struct {
	pb.UnimplementedCommentJobServer
	log     *log.Helper
	subject *biz.CommentSubjectUsecase
	index   *biz.CommentIndexUsecase
	content *biz.CommentContentUsecase
	trans   *biz.TransactionUsecase
}

func NewCommentJobService(subject *biz.CommentSubjectUsecase, index *biz.CommentIndexUsecase, content *biz.CommentContentUsecase, trans *biz.TransactionUsecase, logger log.Logger) *CommentJobService {
	return &CommentJobService{
		log:     log.NewHelper(log.With(logger, "module", "service/comment-job")),
		subject: subject,
		index:   index,
		content: content,
		trans:   trans,
	}
}
