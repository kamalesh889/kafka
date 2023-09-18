package queue

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

type KafkaConsumer struct {
	Reader *kafka.Reader
}

type Consumer interface {
	ReadMessage(ctx context.Context) (kafka.Message, error)
}

func (r *KafkaConsumer) ReadMessage(ctx context.Context) (kafka.Message, error) {

	msg, err := r.Reader.ReadMessage(ctx)
	if err != nil {
		log.Println("Error in Reading message from queue", err)
		return kafka.Message{}, err
	}

	return msg, nil
}
