package queue

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

type KafkaProducer struct {
	Writer *kafka.Writer
}

type Producer interface {
	WriteMessage(context.Context, kafka.Message) error
}

func (p *KafkaProducer) WriteMessage(ctx context.Context, message kafka.Message) error {

	err := p.Writer.WriteMessages(ctx, message)
	if err != nil {
		log.Println("Error in producing message to queue", err)
		return err
	}

	return nil

}
