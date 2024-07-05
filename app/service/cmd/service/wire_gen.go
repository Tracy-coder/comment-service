// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"comment/app/service/internal/biz"
	"comment/app/service/internal/conf"
	"comment/app/service/internal/data"
	"comment/app/service/internal/server"
	"comment/app/service/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	client := data.NewEntClient(confData, logger)
	redisClient := data.NewCacheClient(confData, logger)
	dataData, cleanup, err := data.NewData(client, redisClient, logger)
	if err != nil {
		return nil, nil, err
	}
	commentSubjectRepo := data.NewCommentSubjectRepo(dataData, logger)
	commentSubjectUsecase := biz.NewCommentSubjectUsecase(commentSubjectRepo, logger)
	commentIndexRepo := data.NewCommentIndexRepo(dataData, logger)
	commentIndexCache := data.NewCommentIndexCache(dataData, logger)
	commentIndexUsecase := biz.NewCommentIndexUsecase(commentIndexRepo, commentIndexCache, logger)
	commentContentRepo := data.NewCommentContentRepo(dataData, logger)
	commentContentCache := data.NewCommentContentCache(dataData, logger)
	commentContentUsecase := biz.NewCommentContentUsecase(commentContentRepo, commentContentCache, logger)
	kafkaMessageSender := biz.NewKafkaMessageSender(confServer)
	kafkaIndexSender := biz.NewKafkaIndexSender(confServer)
	commentService := service.NewCommentService(commentSubjectUsecase, commentIndexUsecase, commentContentUsecase, logger, kafkaMessageSender, kafkaIndexSender)
	grpcServer := server.NewGRPCServer(confServer, commentService, logger)
	app := newApp(logger, grpcServer)
	return app, func() {
		cleanup()
	}, nil
}
