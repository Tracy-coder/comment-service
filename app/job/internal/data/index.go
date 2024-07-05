package data

import (
	"context"
	"fmt"
	"strconv"

	svcV1 "comment/api/comment/service/v1"
	"comment/app/job/internal/biz"
	"comment/app/job/internal/data/ent"
	"comment/app/job/internal/data/ent/commentindex"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/redis/go-redis/v9"
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

func (r *commentIndexRepo) InsertCommentIndex(ctx context.Context, req *svcV1.CommentMessage, floor int32, contentID int64) error {
	return r.data.db.CommentIndex(ctx).Create().SetContentID(contentID).SetObjID(req.ObjId).SetOwnerID(req.OwnerId).SetFloor(floor).SetRoot(req.Root).SetParent(req.Parent).Exec(ctx)
}

func (r *commentIndexRepo) GetNextFloor(ctx context.Context, contentID int64) (int32, error) {
	v, err := r.data.db.CommentIndex(ctx).Query().Where(commentindex.ContentID(contentID)).First(ctx)
	if v != nil {
		fmt.Println(v.NextFloor)
		return v.NextFloor, nil
	}
	return 0, err
}
func (r *commentIndexRepo) IncrementFloor(ctx context.Context, contentID int64) error {
	return r.data.db.CommentIndex(ctx).Update().Where(commentindex.ContentID(contentID)).AddNextFloor(1).Exec(ctx)
}

func (r *commentIndexRepo) IncrementCount(ctx context.Context, contentID int64) error {
	return r.data.db.CommentIndex(ctx).Update().Where(commentindex.ContentID(contentID)).AddCount(1).Exec(ctx)
}

func (r *commentIndexRepo) DecrementCount(ctx context.Context, contentID int64) error {
	return r.data.db.CommentIndex(ctx).Update().Where(commentindex.ContentID(contentID)).AddCount(-1).Exec(ctx)
}

func (r *commentIndexRepo) ListRootCommentIndex(ctx context.Context, obj_id int64, page_size int32, page_id int32, order_by_floor bool) ([]redis.Z, error) {
	var err error
	if order_by_floor {
		var v []struct {
			ContentID int64 `json:"content_id,omitempty"`
			Score     int64 `json:"next_floor,omitempty"`
		}
		err = r.data.db.CommentIndex(ctx).
			Query().
			Where(commentindex.ObjID(obj_id)).
			Where(commentindex.Root(0)).
			Order(ent.Asc(commentindex.FieldContentID)).
			Limit(int(page_size)).
			Offset((int(page_id)-1)*int(page_size)).
			Select(commentindex.FieldContentID, commentindex.FieldNextFloor).
			Scan(ctx, &v)
		if err != nil {
			return nil, err
		}
		data := make([]redis.Z, len(v))
		for i, item := range v {
			data[i] = redis.Z{Score: float64(item.Score), Member: strconv.FormatInt(item.ContentID, 10)}
		}
		return data, nil

	} else {
		var v []struct {
			ContentID int64 `json:"content_id,omitempty"`
			Score     int64 `json:"like,omitempty"`
		}
		err = r.data.db.CommentIndex(ctx).
			Query().
			Where(commentindex.ObjID(obj_id)).
			Where(commentindex.Root(0)).
			Order(ent.Desc(commentindex.FieldLike)).
			Limit(int(page_size)).
			Offset((int(page_id)-1)*int(page_size)).
			Select(commentindex.FieldContentID, commentindex.FieldLike).
			Scan(ctx, &v)
		if err != nil {
			return nil, err
		}
		data := make([]redis.Z, len(v))
		for i, item := range v {
			data[i] = redis.Z{Score: float64(item.Score), Member: strconv.FormatInt(item.ContentID, 10)}
		}
		return data, nil
	}

}

// {
//     "obj_id": "1",
//     "order_by_floor": false,
//     "page_id": 1,
//     "page_size": 5
// }
