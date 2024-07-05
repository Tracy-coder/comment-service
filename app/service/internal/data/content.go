package data

import (
	"comment/app/service/internal/biz"
	"comment/app/service/internal/data/ent"
	"comment/app/service/internal/data/ent/commentcontent"
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

var _ biz.CommentContentRepo = (*commentContentRepo)(nil)

type commentContentRepo struct {
	data *Data
	log  *log.Helper
}

func NewCommentContentRepo(data *Data, logger log.Logger) biz.CommentContentRepo {
	return &commentContentRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *commentContentRepo) GetCommentMessage(ctx context.Context, id int64) (string, error) {
	v, err := r.data.db.CommentContent.Query().Where(commentcontent.ID(id)).Select(commentcontent.FieldMessage).First(ctx)
	if err != nil {
		return "", err
	}
	return v.Message, nil
}

func (r *commentContentRepo) GetBatchCommentMessage(ctx context.Context, ids []int64) ([]*string, error) {
	var v []*string
	err := r.data.db.CommentContent.Query().Where(commentcontent.IDIn(ids...)).Order(ent.Asc(commentcontent.FieldID)).Select(commentcontent.FieldMessage).Scan(ctx, &v)
	return v, err
}
