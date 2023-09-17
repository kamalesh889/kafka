package queue

import (
	"context"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

type KafkaConsumer struct {
	Reader *kafka.Reader
}

type Consumer interface {
	FetchMessage(ctx context.Context) (kafka.Message, error)
}

func (r *KafkaConsumer) FetchMessage(ctx context.Context) (kafka.Message, error) {

	msg, err := r.Reader.FetchMessage(ctx)
	if err != nil {
		log.Println("Error in consuming message from queue", err)
	}

	fmt.Println("Message is", msg)

	return msg, nil
}
