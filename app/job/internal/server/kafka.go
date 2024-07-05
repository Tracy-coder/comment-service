package server

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	"github.com/tx7do/kratos-transport/transport/kafka"

	"comment/app/job/internal/conf"
	"comment/app/job/internal/service"
)

// NewKafkaServer create a kafka server.
func NewKafkaServer(c *conf.Server, _ log.Logger, svc *service.CommentJobService) *kafka.Server {
	ctx := context.Background()

	srv := kafka.NewServer(
		kafka.WithAddress(c.Kafka.Addrs),
		kafka.WithCodec("json"),
	)

	registerKafkaSubscribers(ctx, srv, svc)

	return srv
}

func registerKafkaSubscribers(ctx context.Context, srv *kafka.Server, svc *service.CommentJobService) {

	_ = srv.RegisterSubscriber(ctx,
		"comment.subject",
		"group1",
		false,
		registerCommentSubjectHandler(svc.InsertCommentSubject),
		commentSubjectCreator,
	)
	_ = srv.RegisterSubscriber(ctx,
		"comment.message",
		"group1",
		false,
		registerCommentMessageHandler(svc.InsertComment),
		commentMessageCreator,
	)

	_ = srv.RegisterSubscriber(ctx,
		"cache.message",
		"group1",
		false,
		registerSetMessageCacheHandler(svc.SetCommentMessageCache),
		setMessageCacheCreator,
	)

	_ = srv.RegisterSubscriber(ctx,
		"cache.index",
		"group1",
		false,
		registerSetIndexCacheHandler(svc.SetCommentIndexCache),
		setIndexCacheCreator,
	)
}
