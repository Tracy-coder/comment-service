package biz

import (
	svcV1 "comment/api/comment/service/v1"
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type TransactionUsecase struct {
	content CommentContentRepo
	index   CommentIndexRepo
	subject CommentSubjectRepo
	tm      Transaction
	log     *log.Helper
}

func NewTransactionUsecase(content CommentContentRepo, index CommentIndexRepo, subject CommentSubjectRepo, tm Transaction, logger log.Logger) *TransactionUsecase {
	return &TransactionUsecase{content: content, index: index, subject: subject, tm: tm, log: log.NewHelper(logger)}
}

func (s *TransactionUsecase) InsertComment(ctx context.Context, topic string, msg *svcV1.CommentMessage) error {
	if e := s.tm.InTx(ctx, func(ctx context.Context) error {
		if msg.Root == 0 {
			floor, err := s.subject.GetNextFloor(ctx, msg.ObjId)
			if err != nil {
				return err
			}
			contentID, err := s.content.InsertCommentContent(ctx, msg, floor)
			if err != nil {
				s.log.Debug("InsertCommentContent error:", err.Error())
				return err
			}
			if err := s.index.InsertCommentIndex(ctx, msg, floor, contentID); err != nil {
				return err
			}
			if err := s.subject.IncrementFloor(ctx, msg.ObjId); err != nil {
				return err
			}
			if err := s.subject.IncrementCount(ctx, msg.ObjId); err != nil {
				return err
			}
			return nil
		} else {
			floor, err := s.index.GetNextFloor(ctx, msg.Root)
			if err != nil {
				s.log.Debug("GetNextFloor error:", err.Error())
				return err
			}
			contentID, err := s.content.InsertCommentContent(ctx, msg, floor)
			if err != nil {
				s.log.Debug("InsertCommentContent error:", err.Error())
				return err
			}
			if err := s.index.InsertCommentIndex(ctx, msg, floor, contentID); err != nil {
				return err
			}
			if err := s.index.IncrementFloor(ctx, msg.Root); err != nil {
				return err
			}
			if err := s.index.IncrementCount(ctx, msg.Root); err != nil {
				return err
			}
			if err := s.subject.IncrementCount(ctx, msg.ObjId); err != nil {
				return err
			}
			return nil
		}
	}); e != nil {
		return e
	}
	return nil
}
