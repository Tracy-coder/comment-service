package data

import (
	"comment/app/service/internal/conf"
	"context"

	"comment/app/service/internal/data/ent"

	"github.com/go-kratos/kratos/v2/log"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	_ "github.com/jackc/pgx/v4/stdlib"
	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewCommentSubjectRepo, NewCommentIndexRepo, NewCommentContentRepo, NewEntClient, NewCacheClient, NewCommentContentCache, NewCommentIndexCache)

// Data .
type Data struct {
	// TODO wrapped database client
	db    *ent.Client
	cache *redis.Client
	log   *log.Helper
}

// NewData .
func NewData(entClient *ent.Client, cache *redis.Client, logger log.Logger) (*Data, func(), error) {
	l := log.NewHelper(log.With(logger, "module", "data/logger-service"))
	d := &Data{
		db:    entClient,
		cache: cache,
		log:   l,
	}

	return d, func() {
		l.Info("message", "closing the data resources")
		if err := d.db.Close(); err != nil {
			l.Error(err)
		}

	}, nil
}
func NewEntClient(conf *conf.Data, logger log.Logger) *ent.Client {
	l := log.NewHelper(log.With(logger, "module", "ent/data/logger-service"))

	client, err := ent.Open(
		conf.Database.Driver,
		conf.Database.Source,
	)
	if err != nil {
		l.Fatalf("failed opening connection to db: %v", err)
	}

	return client
}

func NewCacheClient(conf *conf.Data, logger log.Logger) *redis.Client {
	l := log.NewHelper(log.With(logger, "module", "ent/data/logger-service"))
	rdb := redis.NewClient(&redis.Options{
		Addr:         conf.Redis.Addr,
		Password:     conf.Redis.Password,
		DB:           int(conf.Redis.Db),
		DialTimeout:  conf.Redis.DialTimeout.AsDuration(),
		WriteTimeout: conf.Redis.WriteTimeout.AsDuration(),
		ReadTimeout:  conf.Redis.ReadTimeout.AsDuration(),
	})

	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		l.Fatalf("failed opening connection to rdb: %v", err)
	}
	return rdb
}
