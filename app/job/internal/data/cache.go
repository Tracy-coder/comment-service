package data

import (
	"comment/app/job/internal/biz"
	"context"
	"strconv"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/redis/go-redis/v9"
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

func (r *commentIndexCache) SetCommentIndexCache(ctx context.Context, key string, z []redis.Z) error {
	_, err := r.data.cache.ZAdd(ctx, key, z...).Result()
	return err
}
