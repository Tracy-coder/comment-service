package biz

import (
	"context"
	"time"

	"github.com/google/wire"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(NewCommentSubjectUsecase, NewCommentIndexUsecase, NewCommentContentUsecase, NewKafkaMessageSender, NewKafkaIndexSender)

type Transaction interface {
	InTx(context.Context, func(ctx context.Context) error) error
}

type CommentObj struct {
	ContentID int64
	ObjID     int64
	OwnerID   int64
	Root      int64
	Parent    int64
	Floor     int32
	Like      int32
	Hate      int32
	State     int32
	Message   string
	CreatedAt time.Time
}
