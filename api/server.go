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
	httpClient HTTPClient
	service    Service
	db         model.Repository
	producer   queue.Producer
	consumer   queue.Consumer
}

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
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

	s.service = s

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
		Brokers: []string{brokerAddr},
		Topic:   topic,
		GroupID: "group-zocket",
	})

	return &queue.KafkaConsumer{
		Reader: reader,
	}
}
