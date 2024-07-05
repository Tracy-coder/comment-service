package biz

import (
	v1 "comment/api/comment/service/v1"
	"comment/app/service/internal/conf"
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

type Event interface {
	Key() string
	Value() []byte
}

type Handler func(context.Context, Event) error

type Sender interface {
	Send(ctx context.Context, msg Event) error
	Close() error
}

type Receiver interface {
	Receive(ctx context.Context, handler Handler) error
	Close() error
}

var (
	_ Sender = (*KafkaMessageSender)(nil)
	_ Event  = (*Message)(nil)
)

type Message struct {
	key   string
	value []byte
}

func (m *Message) Key() string {
	return m.key
}

func (m *Message) Value() []byte {
	return m.value
}

func NewMessage(key string, value []byte) Event {
	return &Message{
		key:   key,
		value: value,
	}
}

type KafkaMessageSender struct {
	writer *kafka.Writer
	Jobs   chan v1.CommentMessageCache
}

type KafkaIndexSender struct {
	writer *kafka.Writer
	Jobs   chan v1.CommentIndexCache
}

func (s *KafkaMessageSender) SendMessageCacheSignal(ctx context.Context, id int64, message string) error {
	data := &v1.CommentMessageCache{
		Id:      id,
		Message: message,
	}

	bytes, err := json.Marshal(data)
	fmt.Println("bytes:", bytes)
	if err != nil {
		return err
	}

	msg := NewMessage("", bytes)
	return s.Send(ctx, msg)
}

func (s *KafkaMessageSender) Work() {
	for job := range s.Jobs {
		log.Printf("I'm working! Job:%v", job)
		err := s.SendMessageCacheSignal(context.Background(), job.GetId(), job.GetMessage())
		if err != nil {
			log.Printf("KafkaMessageSender work err:", err)
		}
	}
}
func (s *KafkaIndexSender) SendIndexCacheSignal(ctx context.Context, obj_id int64, order_by_floor bool) error {
	data := &v1.ListCommentRequest{
		ObjId:        obj_id,
		OrderByFloor: order_by_floor,
	}

	bytes, err := json.Marshal(data)
	fmt.Println("bytes:", bytes)
	if err != nil {
		return err
	}

	msg := NewMessage("", bytes)
	return s.Send(ctx, msg)
}

func (s *KafkaIndexSender) Work() {
	for job := range s.Jobs {
		log.Printf("I'm working! Job:%v", job)
		err := s.SendIndexCacheSignal(context.Background(), job.GetObjId(), job.GetOrderByFloor())
		if err != nil {
			log.Printf("KafkaIndexSender work err:", err)
		}
	}
}

func (s *KafkaMessageSender) Send(ctx context.Context, message Event) error {
	err := s.writer.WriteMessages(ctx, kafka.Message{
		Key:   []byte(message.Key()),
		Value: message.Value(),
	})
	if err != nil {
		return err
	}
	return nil
}

func (s *KafkaMessageSender) Close() error {
	err := s.writer.Close()
	if err != nil {
		return err
	}
	return nil
}

func (s *KafkaIndexSender) Send(ctx context.Context, message Event) error {
	err := s.writer.WriteMessages(ctx, kafka.Message{
		Key:   []byte(message.Key()),
		Value: message.Value(),
	})
	if err != nil {
		return err
	}
	return nil
}

func (s *KafkaIndexSender) Close() error {
	err := s.writer.Close()
	if err != nil {
		return err
	}
	return nil
}

func NewKafkaMessageSender(conf *conf.Server) *KafkaMessageSender {
	return &KafkaMessageSender{
		writer: &kafka.Writer{
			Topic: "cache.message",
			Addr:  kafka.TCP(conf.Kafka.Addrs...),
			// NOTE: When Topic is not defined here, each Message must define it instead.
			Balancer: &kafka.LeastBytes{}},
		Jobs: make(chan v1.CommentMessageCache, 100),
	}

}

func NewKafkaIndexSender(conf *conf.Server) *KafkaIndexSender {
	return &KafkaIndexSender{
		writer: &kafka.Writer{
			Topic: "cache.index",
			Addr:  kafka.TCP(conf.Kafka.Addrs...),
			// NOTE: When Topic is not defined here, each Message must define it instead.
			Balancer: &kafka.LeastBytes{}},
		Jobs: make(chan v1.CommentIndexCache, 100),
	}

}
