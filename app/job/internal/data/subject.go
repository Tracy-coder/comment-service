package data

import (
	"context"
	"fmt"

	svcV1 "comment/api/comment/service/v1"
	"comment/app/job/internal/biz"
	"comment/app/job/internal/data/ent/commentsubject"

	"github.com/go-kratos/kratos/v2/log"
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

func (r *commentSubjectRepo) InsertCommentSubject(ctx context.Context, req *svcV1.CommentSubject) error {
	return r.data.db.CommentSubject(ctx).Create().SetOwnerID(req.OwnerId).Exec(ctx)
}

func (r *commentSubjectRepo) GetNextFloor(ctx context.Context, objID int64) (int32, error) {
	v, err := r.data.db.CommentSubject(ctx).Query().Where(commentsubject.ID(objID)).First(ctx)
	if v != nil {
		fmt.Println(v.NextFloor)
		return v.NextFloor, nil
	}
	return 0, err
}

func (r *commentSubjectRepo) IncrementFloor(ctx context.Context, objID int64) error {
	return r.data.db.CommentSubject(ctx).Update().Where(commentsubject.ID(objID)).AddNextFloor(1).Exec(ctx)
}

func (r *commentSubjectRepo) IncrementCount(ctx context.Context, objID int64) error {
	return r.data.db.CommentSubject(ctx).Update().Where(commentsubject.ID(objID)).AddCount(1).Exec(ctx)
}

func (r *commentSubjectRepo) DecrementCount(ctx context.Context, objID int64) error {
	return r.data.db.CommentSubject(ctx).Update().Where(commentsubject.ID(objID)).AddCount(-1).Exec(ctx)
}
