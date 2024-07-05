package data

import (
	"context"

	"comment/app/service/internal/biz"
	"comment/app/service/internal/data/ent"
	"comment/app/service/internal/data/ent/commentindex"

	"github.com/go-kratos/kratos/v2/log"
)

var _ biz.CommentIndexRepo = (*commentIndexRepo)(nil)

type commentIndexRepo struct {
	data *Data
	log  *log.Helper
}

// NewGreeterRepo .
func NewCommentIndexRepo(data *Data, logger log.Logger) biz.CommentIndexRepo {
	return &commentIndexRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *commentIndexRepo) GetComment(ctx context.Context, id int64) (*biz.CommentObj, error) {
	v, err := r.data.db.CommentIndex.Query().Where(commentindex.ContentID(id)).First(ctx)
	if err != nil {
		return nil, err
	}
	var comment biz.CommentObj
	comment.ContentID = v.ContentID
	comment.CreatedAt = v.CreatedAt
	comment.Floor = v.Floor
	comment.Hate = v.Hate
	comment.Like = v.Like
	comment.ObjID = v.ObjID
	comment.OwnerID = v.OwnerID
	comment.Parent = v.Parent
	comment.Root = v.Root
	comment.State = int32(v.State)

	return &comment, nil
}

func (r *commentIndexRepo) GetSubCommentIndex(ctx context.Context, root int64, limit int) ([]int64, error) {
	var v []int64
	err := r.data.db.CommentIndex.Query().Where(commentindex.Root(root)).Order(ent.Asc(commentindex.FieldContentID)).Limit(limit).Select(commentindex.FieldContentID).Scan(ctx, &v)
	if err != nil {
		return nil, err
	}
	return v, nil
	// items := make([]*biz.CommentObj, 0)
	// for _, i := range v {
	// 	item, err := r.GetComment(ctx, i.ContentID)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	items = append(items, item)
	// }
	// return items, nil
}

func (r *commentIndexRepo) GetBatchComment(ctx context.Context, id []int64) ([]*biz.CommentObj, error) {
	objs, err := r.data.db.CommentIndex.Query().Where(commentindex.ContentIDIn(id...)).All(ctx)
	if err != nil {
		return nil, err
	}
	comments := make([]*biz.CommentObj, 0)
	for _, v := range objs {
		var comment biz.CommentObj
		comment.ContentID = v.ContentID
		comment.CreatedAt = v.CreatedAt
		comment.Floor = v.Floor
		comment.Hate = v.Hate
		comment.Like = v.Like
		comment.ObjID = v.ObjID
		comment.OwnerID = v.OwnerID
		comment.Parent = v.Parent
		comment.Root = v.Root
		comment.State = int32(v.State)
		comments = append(comments, &comment)
	}

	return comments, nil
}

func (r *commentIndexRepo) ListRootCommentIndex(ctx context.Context, obj_id int64, page_size int32, page_id int32, order_by_floor bool) ([]int64, error) {
	var v []int64
	var err error
	if order_by_floor {
		err = r.data.db.CommentIndex.
			Query().
			Where(commentindex.ObjID(obj_id)).
			Where(commentindex.Root(0)).
			Order(ent.Asc(commentindex.FieldContentID)).
			Limit(int(page_size)).
			Offset((int(page_id)-1)*int(page_size)).
			Select(commentindex.FieldContentID).
			Scan(ctx, &v)

	} else {
		err = r.data.db.CommentIndex.
			Query().
			Where(commentindex.ObjID(obj_id)).
			Where(commentindex.Root(0)).
			Order(ent.Desc(commentindex.FieldLike)).
			Limit(int(page_size)).
			Offset((int(page_id)-1)*int(page_size)).
			Select(commentindex.FieldContentID).
			Scan(ctx, &v)
	}
	if err != nil {
		return nil, err
	}
	return v, nil
}
