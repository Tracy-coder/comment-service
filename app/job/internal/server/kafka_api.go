package server

import (
	"context"
	"fmt"

	svcV1 "comment/api/comment/service/v1"

	"github.com/tx7do/kratos-transport/broker"
)

func commentSubjectCreator() broker.Any  { return &svcV1.CommentSubject{} }
func commentMessageCreator() broker.Any  { return &svcV1.CommentMessage{} }
func setMessageCacheCreator() broker.Any { return &svcV1.CommentMessageCache{} }
func setIndexCacheCreator() broker.Any   { return &svcV1.CommentIndexCache{} }

type commentSubjectHandler func(_ context.Context, topic string, headers broker.Headers, msg *svcV1.CommentSubject) error
type commentMessageHandler func(_ context.Context, topic string, headers broker.Headers, msg *svcV1.CommentMessage) error
type setMessageCacheHandler func(_ context.Context, topic string, headers broker.Headers, msg *svcV1.CommentMessageCache) error
type setIndexCacheHandler func(_ context.Context, topic string, headers broker.Headers, msg *svcV1.CommentIndexCache) error

func registerCommentSubjectHandler(fnc commentSubjectHandler) broker.Handler {
	return func(ctx context.Context, event broker.Event) error {
		switch t := event.Message().Body.(type) {
		case *svcV1.CommentSubject:
			if err := fnc(ctx, event.Topic(), event.Message().Headers, t); err != nil {
				return err
			}
		default:
			return fmt.Errorf("unsupported type: %T", t)
		}
		return nil
	}
}

func registerCommentMessageHandler(fnc commentMessageHandler) broker.Handler {
	return func(ctx context.Context, event broker.Event) error {
		switch t := event.Message().Body.(type) {
		case *svcV1.CommentMessage:
			if err := fnc(ctx, event.Topic(), event.Message().Headers, t); err != nil {
				return err
			}
		default:
			return fmt.Errorf("unsupported type: %T", t)
		}
		return nil
	}
}

func registerSetMessageCacheHandler(fnc setMessageCacheHandler) broker.Handler {
	return func(ctx context.Context, event broker.Event) error {
		fmt.Println("event:", event)

		switch t := event.Message().Body.(type) {
		case *svcV1.CommentMessageCache:
			if err := fnc(ctx, event.Topic(), event.Message().Headers, t); err != nil {
				return err
			}
		default:
			return fmt.Errorf("unsupported type: %T", t)
		}
		return nil
	}
}

// {
//   "id":6,
//   "message":"qwerty"
// }

func registerSetIndexCacheHandler(fnc setIndexCacheHandler) broker.Handler {
	return func(ctx context.Context, event broker.Event) error {
		fmt.Println("event:", event)
		switch t := event.Message().Body.(type) {
		case *svcV1.CommentIndexCache:
			if err := fnc(ctx, event.Topic(), event.Message().Headers, t); err != nil {
				return err
			}
		default:
			return fmt.Errorf("unsupported type: %T", t)
		}
		return nil
	}
}
