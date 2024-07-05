package data

import (
	"comment/app/job/internal/biz"
	"comment/app/job/internal/conf"
	"context"
	"fmt"

	"comment/app/job/internal/data/ent"

	"entgo.io/ent/dialect/sql"
	"github.com/go-kratos/kratos/v2/log"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	_ "github.com/jackc/pgx/v4/stdlib"
	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewCacheClient, NewCommentContentCache, NewCommentIndexCache, NewCommentSubjectRepo, NewCommentIndexRepo, NewCommentContentRepo, NewTransaction)

// Data .
type Data struct {
	// TODO wrapped database client
	db    *ent.Database
	log   *log.Helper
	cache *redis.Client
}

// NewTransaction .
func NewTransaction(data *Data) biz.Transaction {
	return data.db
}

func NewData(conf *conf.Data, logger log.Logger, cache *redis.Client) (*Data, func(), error) {
	log := log.NewHelper(log.With(logger, "module", "data/logger-job"))
	drv, err := sql.Open(
		conf.Database.Driver,
		conf.Database.Source,
	)
	if err != nil {
		log.Errorf("failed opening connection to sqlite: %v", err)
		return nil, nil, err
	}

	d := &Data{
		db:    ent.NewDatabase(ent.Driver(drv)),
		cache: cache,
	}
	fmt.Println(d)

	return d, func() {
		log.Info("message", "closing the data resources")
		if err := drv.Close(); err != nil {
			log.Error(err)
		}
	}, nil
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
