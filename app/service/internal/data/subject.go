package data

import (
	v1 "comment/api/comment/service/v1"
	"comment/app/service/internal/biz"
	"comment/app/service/internal/data/ent/commentsubject"
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var _ biz.CommentSubjectRepo = (*commentSubjectRepo)(nil)

type commentSubjectRepo struct {
	data *Data
	log  *log.Helper
}

// NewGreeterRepo .
func NewCommentSubjectRepo(data *Data, logger log.Logger) biz.CommentSubjectRepo {
	return &commentSubjectRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *commentSubjectRepo) GetSubject(ctx context.Context, req *v1.GetSubjectRequest) (*v1.GetSubjectReply, error) {
	v, err := r.data.db.CommentSubject.Query().Where(commentsubject.ID(req.ObjId)).First(ctx)
	if err != nil {
		return nil, err
	}
	var subject v1.GetSubjectReply
	subject.ObjId = v.ID
	subject.Count = v.Count
	subject.OwnerId = v.OwnerID
	subject.State = int32(v.State)
	subject.CreatedAt = timestamppb.New(v.CreatedAt)
	subject.NextFloor = v.NextFloor
	subject.UpdatedAt = timestamppb.New(v.UpdatedAt)
	return &subject, nil
}
