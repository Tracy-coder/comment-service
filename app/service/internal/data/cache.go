package data

import (
	"comment/app/service/internal/biz"
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

var _ biz.CommentContentCache = (*commentContentCache)(nil)
var _ biz.CommentIndexCache = (*commentIndexCache)(nil)

type commentContentCache struct {
	data *Data
	log  *log.Helper
}

func NewCommentContentCache(data *Data, logger log.Logger) biz.CommentContentCache {
	return &commentContentCache{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *commentContentCache) GetCommentMessageCache(ctx context.Context, id int64) (string, error) {
	wanted, err := r.data.cache.Get(ctx, "comment_id:"+strconv.FormatInt(id, 10)).Result()
	if err != nil {
		return "", err
	}
	return wanted, nil
}

func (r *commentContentCache) SetCommentMessageCache(ctx context.Context, id int64, message string) error {
	_, err := r.data.cache.Set(ctx, "comment_id:"+strconv.FormatInt(id, 10), message, time.Minute).Result()
	return err

}

type commentIndexCache struct {
	data *Data
	log  *log.Helper
}

func NewCommentIndexCache(data *Data, logger log.Logger) biz.CommentIndexCache {
	return &commentIndexCache{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *commentIndexCache) ListRootCommentIndexCache(ctx context.Context, obj_id int64, page_size int32, page_id int32, order_by_floor bool) ([]int64, error) {
	wanted := r.data.cache.ZRange(ctx, strconv.FormatInt(obj_id, 10)+strconv.FormatBool(order_by_floor), (int64(page_id-1) * int64(page_size)), (int64(page_id) * int64(page_size))).Val()
	if len(wanted) == 0 {
		return nil, errors.New("no data in cache")
	}
	ids := make([]int64, len(wanted))
	for i, str_id := range wanted {
		id, err := strconv.ParseInt(str_id, 10, 64)
		if err != nil {
			return nil, err
		}
		ids[i] = id
	}
	return ids, nil
}
