package data

import (
	"context"
	"fmt"

	svcV1 "comment/api/comment/service/v1"
	"comment/app/job/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

var _ biz.CommentContentRepo = (*commentContentRepo)(nil)

type commentContentRepo struct {
	data *Data
	log  *log.Helper
}

// NewGreeterRepo .
func NewCommentContentRepo(data *Data, logger log.Logger) biz.CommentContentRepo {
	return &commentContentRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *commentContentRepo) InsertCommentContent(ctx context.Context, req *svcV1.CommentMessage, floor int32) (int64, error) {
	v, err := r.data.db.CommentContent(ctx).Create().
		SetObjID(req.ObjId).
		SetOwnerID(req.OwnerId).
		SetFloor(floor).
		SetRoot(req.Root).
		SetMessage(req.Message).Save(ctx)
	if err != nil {
		return -1, err
	}
	fmt.Println("v:", v)
	return v.ID, nil
}
