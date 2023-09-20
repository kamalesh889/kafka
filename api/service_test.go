package api

import (
	"Kafka/model"
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/segmentio/kafka-go"
	"github.com/stretchr/testify/assert"
)

func TestCreateProduct(t *testing.T) {

	product := &ProductRequest{
		UserId:        1,
		ProductName:   "test",
		ProductDesc:   "testing",
		ProductImages: []string{"url1", "url2"},
		ProductPrice:  "100",
	}

	t.Run("Success case", func(t *testing.T) {

		mockrepo.EXPECT().GetUser(gomock.Any()).Return(nil)
		mockrepo.EXPECT().CreateProduct(gomock.Any()).Return(uint64(1), nil)
		mockproducer.EXPECT().WriteMessage(context.Background(), gomock.Any()).Return(nil)
		err := testServer.CreateProduct(product)

		assert.NoError(t, err)

	})

	t.Run("Failure case ", func(t *testing.T) {

		t.Log("Error in getting userdetails")

		mockrepo.EXPECT().GetUser(gomock.Any()).Return(errors.New("user error"))
		err := testServer.CreateProduct(product)

		assert.Error(t, err)

	})

	t.Run("Failure case", func(t *testing.T) {

		t.Log("Error in creating product")

		mockrepo.EXPECT().GetUser(gomock.Any()).Return(nil)
		mockrepo.EXPECT().CreateProduct(gomock.Any()).Return(uint64(0), errors.New("error in creating product"))
		err := testServer.CreateProduct(product)

		assert.Error(t, err)

	})

	t.Run("Failure case", func(t *testing.T) {

		t.Log("Error in pushing product to queue")

		mockrepo.EXPECT().GetUser(gomock.Any()).Return(nil)
		mockrepo.EXPECT().CreateProduct(gomock.Any()).Return(uint64(0), nil)
		mockproducer.EXPECT().WriteMessage(context.Background(), gomock.Any()).Return(errors.New("error in pushing product into queue"))
		err := testServer.CreateProduct(product)

		assert.Error(t, err)

	})
}

func TestCreateUser(t *testing.T) {

	user := &UserRequest{
		Name:      "Avi",
		Mobile:    "9087190876",
		Latitude:  "65.01",
		Longitude: "54.11",
	}

	t.Run("Success case", func(t *testing.T) {

		mockrepo.EXPECT().CreateUser(gomock.Any()).Return(uint64(1), nil)

		userid, err := testServer.CreateUser(user)
		assert.NoError(t, err)
		assert.Equal(t, uint64(1), userid)

	})

	t.Run("Failure case", func(t *testing.T) {

		mockrepo.EXPECT().CreateUser(gomock.Any()).Return(uint64(0), errors.New("error in creating user"))

		userid, err := testServer.CreateUser(user)
		assert.Error(t, err)
		assert.Equal(t, uint64(0), userid)
	})
}

func TestGetProductFromKafka(t *testing.T) {

	msg := kafka.Message{
		Value: []byte("1"),
	}

	product := &model.Product{
		ProductId:     1,
		ProductName:   "testproduct",
		ProductImages: []string{"https://source.unsplash.com/random", "https://source.unsplash.com/random"},
	}

	t.Run("Success case", func(t *testing.T) {

		mockconsumer.EXPECT().ReadMessage(gomock.Any()).Return(msg, nil)
		mockrepo.EXPECT().GetProduct(gomock.Any()).Return(product, nil)
		mockrepo.EXPECT().UpdateProduct(gomock.Any()).Return(nil)

		testServer.GetProductFromKafka(false)
	})

	t.Run("Failure case", func(t *testing.T) {

		mockconsumer.EXPECT().ReadMessage(gomock.Any()).Return(msg, errors.New("queue error"))
		testServer.GetProductFromKafka(false)
	})
}

func TestDownloadAndCompressImage(t *testing.T) {

	successImageUrl := "https://source.unsplash.com/random"
	failureimageurl := "htp.uy"
	imageName := "test1.jpg"
	envTest := false

	expectedPath := "/Users/kamaleshmohapatra/Documents/go/kafka/api/compressed_test1.jpg.gz"

	t.Run("Success case", func(t *testing.T) {

		path, err := testServer.DownloadAndCompressImage(successImageUrl, imageName, envTest)

		assert.Equal(t, expectedPath, path)
		assert.NoError(t, err)

	})

	t.Run("Failure case", func(t *testing.T) {

		t.Log("Error due to passing of invalid url")

		path, err := testServer.DownloadAndCompressImage(failureimageurl, imageName, envTest)

		assert.Equal(t, "", path)
		assert.Error(t, err)

	})

	t.Run("Failure case", func(t *testing.T) {

		t.Log("Error due to passing of wrong enviroment")

		path, err := testServer.DownloadAndCompressImage(successImageUrl, imageName, true)

		assert.Equal(t, "", path)
		assert.Error(t, err)

	})

}
