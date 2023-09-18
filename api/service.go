package api

import (
	"Kafka/model"
	"context"
	"strconv"
	"time"

	"github.com/segmentio/kafka-go"
)

type Service interface {
	CreateProduct(*ProductRequest) error
	CreateUser(*UserRequest) (uint64, error)
}

func (s *server) CreateProduct(productDetails *ProductRequest) error {

	err := s.db.GetUser(productDetails.UserId)
	if err != nil {
		return err
	}

	product := &model.Product{
		ProductName:   productDetails.ProductName,
		ProductDesc:   productDetails.ProductDesc,
		ProductImages: productDetails.ProductImages,
		ProductPrice:  productDetails.ProductPrice,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	productid, err := s.db.CreateProduct(product)
	if err != nil {
		return err
	}

	// Push into the queue

	kafkamessage := kafka.Message{
		Value: []byte(strconv.Itoa(int(productid))),
	}

	err = s.producer.WriteMessage(context.Background(), kafkamessage)
	if err != nil {
		return err
	}

	return nil
}

func (s *server) CreateUser(userDetails *UserRequest) (uint64, error) {

	user := &model.User{
		Name:      userDetails.Name,
		Mobile:    userDetails.Mobile,
		Latitude:  userDetails.Latitude,
		Longitude: userDetails.Longitude,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	userid, err := s.db.CreateUser(user)
	if err != nil {
		return 0, err
	}

	return userid, nil
}
