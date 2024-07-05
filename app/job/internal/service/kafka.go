package service

import (
	svcV1 "comment/api/comment/service/v1"
	"context"
	"fmt"
	"strconv"

	"github.com/tx7do/kratos-transport/broker"
)

func (s *CommentJobService) InsertCommentSubject(ctx context.Context, topic string, headers broker.Headers, msg *svcV1.CommentSubject) error {
	fmt.Println("InsertCommentSubject() Topic: ", topic)

	if err := s.subject.InsertCommentSubject(context.Background(), msg); err != nil {
		s.log.Debug("InsertCommentSubject Insert", err.Error())
		return err
	}

	return nil
}

func (s *CommentJobService) InsertComment(ctx context.Context, topic string, headers broker.Headers, msg *svcV1.CommentMessage) error {
	return s.trans.InsertComment(ctx, topic, msg)
}

func (s *CommentJobService) SetCommentMessageCache(ctx context.Context, topic string, headers broker.Headers, msg *svcV1.CommentMessageCache) error {
	s.log.Debug("set cache:", msg)
	return s.content.SetCommentMessageCache(ctx, msg.GetId(), msg.GetMessage())
}

func (s *CommentJobService) SetCommentIndexCache(ctx context.Context, topic string, headers broker.Headers, msg *svcV1.CommentIndexCache) error {
	key := strconv.FormatInt(msg.GetObjId(), 10) + strconv.FormatBool(msg.GetOrderByFloor())
	data, err := s.index.ListRootCommentIndex(ctx, msg.GetObjId(), msg.GetOrderByFloor())
	if err != nil {
		return err
	}
	err = s.index.SetCommentIndexCache(ctx, key, data)
	if err != nil {
		return err
	}
	return nil
}
