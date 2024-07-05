package service

import (
	pb "comment/api/comment/service/v1"
	v1 "comment/api/comment/service/v1"
	"comment/app/service/internal/biz"
	"context"

	"sync"

	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// GreeterService is a greeter service.
type CommentService struct {
	pb.UnimplementedCommentServer
	log     *log.Helper
	subject *biz.CommentSubjectUsecase
	index   *biz.CommentIndexUsecase
	content *biz.CommentContentUsecase
	writer  *biz.KafkaMessageSender
	writer2 *biz.KafkaIndexSender
}

func NewCommentService(subject *biz.CommentSubjectUsecase, index *biz.CommentIndexUsecase, content *biz.CommentContentUsecase, logger log.Logger, writer *biz.KafkaMessageSender, writer2 *biz.KafkaIndexSender) *CommentService {
	go writer.Work()
	go writer2.Work()
	return &CommentService{
		log:     log.NewHelper(log.With(logger, "module", "service/comment-service")),
		subject: subject,
		index:   index,
		content: content,
		writer:  writer,
		writer2: writer2,
	}
}

func (s *CommentService) GetSubject(ctx context.Context, req *v1.GetSubjectRequest) (*v1.GetSubjectReply, error) {
	return s.subject.GetSubject(ctx, req)
}

func (s *CommentService) GetComment(ctx context.Context, req *v1.GetCommentRequest) (*v1.GetCommentReply, error) {
	rootObj, err := s.index.GetComment(ctx, req.ContentId)
	if err != nil {
		return nil, err
	}
	rootMessage, err := s.content.GetCommentMessageCache(ctx, req.ContentId)
	if err != nil {
		log.Debug("not found in cache")
		rootMessage, err = s.content.GetCommentMessage(ctx, req.ContentId)

		if err != nil {
			return nil, err
		}
		s.writer.Jobs <- v1.CommentMessageCache{
			Id:      req.ContentId,
			Message: rootMessage,
		}
		if err != nil {
			return nil, err
		}
		// 并发 需要写一个后台pick up job的worker
		// 发送信号到kafka 让comment-job处理
	}

	subIndexs, err := s.index.GetSubCommentIndex(ctx, req.ContentId)
	if err != nil {
		return nil, err
	}
	subObjs, err := s.index.GetBatchComment(ctx, subIndexs)
	if err != nil {
		return nil, err
	}
	subMessages, err := s.content.GetBatchCommentMessage(ctx, subIndexs)
	if err != nil {
		return nil, err
	}
	comment := v1.CommentObj{
		ContentId: rootObj.ContentID,
		CreatedAt: timestamppb.New(rootObj.CreatedAt),
		Floor:     rootObj.Floor,
		Hate:      rootObj.Hate,
		Like:      rootObj.Like,
		ObjId:     rootObj.ObjID,
		OwnerId:   rootObj.OwnerID,
		Parent:    rootObj.Parent,
		Root:      rootObj.Root,
		State:     int32(rootObj.State),
		Message:   rootMessage,
	}
	subComment := make([]*v1.CommentObj, 0)
	for i := 0; i < len(subIndexs); i++ {
		subComment = append(subComment, &v1.CommentObj{
			ContentId: subObjs[i].ContentID,
			CreatedAt: timestamppb.New(subObjs[i].CreatedAt),
			Floor:     subObjs[i].Floor,
			Hate:      subObjs[i].Hate,
			Like:      subObjs[i].Like,
			ObjId:     subObjs[i].ObjID,
			OwnerId:   subObjs[i].OwnerID,
			Parent:    subObjs[i].Parent,
			Root:      subObjs[i].Root,
			State:     int32(subObjs[i].State),
			Message:   *subMessages[i],
		})

	}
	return &v1.GetCommentReply{
		Comment:    &comment,
		SubComment: subComment,
	}, nil
}

func (s *CommentService) ListComment(ctx context.Context, req *v1.ListCommentRequest) (*v1.ListCommentReply, error) {
	indexs, err := s.index.ListRootCommentIndexCache(ctx, req.ObjId, req.PageSize, req.PageId, req.OrderByFloor)
	if err == nil {
		log.Debug("get cache success:", indexs)
	}
	if err != nil {
		log.Debug("ListRootCommentIndexCache err:", err)

		indexs, err = s.index.ListRootCommentIndex(ctx, req.ObjId, req.PageSize, req.PageId, req.OrderByFloor)
		if err != nil {
			return nil, err
		}
		if len(indexs) != 0 {
			s.writer2.Jobs <- v1.CommentIndexCache{
				ObjId:        req.GetObjId(),
				OrderByFloor: req.GetOrderByFloor(),
			}
		}

	}

	listReply := make([]*v1.GetCommentReply, len(indexs))
	var wg sync.WaitGroup
	errCh := make(chan error, len(indexs))
	for i, index := range indexs {
		wg.Add(1)
		go func(i int, index int64) {
			defer wg.Done()
			reply, err := s.GetComment(ctx, &v1.GetCommentRequest{ContentId: index})
			if err != nil {
				errCh <- err
				return
			}
			listReply[i] = reply
		}(i, index)
	}

	wg.Wait()
	close(errCh)
	if len(errCh) != 0 {
		return nil, <-errCh
	}
	return &v1.ListCommentReply{Comment: listReply}, nil
}
