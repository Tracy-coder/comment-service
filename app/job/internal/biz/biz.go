package biz

import (
	"context"

	"github.com/google/wire"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(NewCommentSubjectUsecase, NewCommentIndexUsecase, NewCommentContentUsecase, NewTransactionUsecase)

type Transaction interface {
	InTx(ctx context.Context, f func(context.Context) error) error
}
