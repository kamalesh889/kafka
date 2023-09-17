package api

import (
	"Kafka/model"
	"Kafka/queue"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/segmentio/kafka-go"
)

type server struct {
	router     *mux.Router
	db         *model.Database
	httpClient *http.Client
	producer   *queue.KafkaProducer
	consumer   *queue.KafkaConsumer
}

const (
	topic      = "topic-zocket"
	brokerAddr = "localhost:9092"
)

func NewServer(db *model.Database) (*server, error) {

	s := &server{}

	s.router = mux.NewRouter()
	s.db = db
	s.httpClient = &http.Client{}

	s.producer = NewKafkaProducer()
	s.consumer = NewKafkaConsumer()

	return s, nil

}

func NewKafkaProducer() *queue.KafkaProducer {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{brokerAddr},
		Topic:   topic,
	})

	return &queue.KafkaProducer{
		Writer: writer,
	}
}

func NewKafkaConsumer() *queue.KafkaConsumer {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{brokerAddr},
		Topic:     topic,
		Partition: 0,
	})

	return &queue.KafkaConsumer{
		Reader: reader,
	}
}
