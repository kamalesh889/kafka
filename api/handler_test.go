package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	repomock "Kafka/model/mock"
	qmock "Kafka/queue/mock"

	gomock "github.com/golang/mock/gomock"
	"github.com/gorilla/mux"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	testServer   *server
	mockservice  *MockService
	mockproducer *qmock.MockProducer
	mockconsumer *qmock.MockConsumer
	mockrepo     *repomock.MockRepository
)

func TestServer(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockproducer = qmock.NewMockProducer(ctrl)
	mockconsumer = qmock.NewMockConsumer(ctrl)
	mockrepo = repomock.NewMockRepository(ctrl)
	mockservice = NewMockService(ctrl)

	testServer = &server{}

	testServer.router = mux.NewRouter()
	testServer.producer = mockproducer
	testServer.consumer = mockconsumer
	testServer.db = mockrepo
	testServer.service = mockservice
	testServer.httpClient = &http.Client{}

}

func init() {

	t := &testing.T{}
	TestServer(t)
}

func TestCreateUserHandler(t *testing.T) {

	reqbody := UserRequest{
		Name:      "Avi",
		Mobile:    "9087190876",
		Latitude:  "65.01",
		Longitude: "54.11",
	}

	mockbyt, _ := json.Marshal(reqbody)

	t.Run("Success case", func(t *testing.T) {

		req := httptest.NewRequest(http.MethodPost, "/create-user", bytes.NewBuffer(mockbyt))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		mockservice.EXPECT().CreateUser(gomock.Any()).Return(uint64(1), nil)

		testServer.CreateUserHandler(rec, req)
		assert.Equal(t, http.StatusCreated, rec.Code)

		var userResp UserResponse
		err := json.NewDecoder(rec.Body).Decode(&userResp)
		require.NoError(t, err)
		assert.Equal(t, reqbody.Name, userResp.Name)

	})
	t.Run("Failure case", func(t *testing.T) {

		req := httptest.NewRequest(http.MethodPost, "/create-user", bytes.NewBuffer(mockbyt))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		mockservice.EXPECT().CreateUser(gomock.Any()).Return(uint64(0), errors.New("Error from database"))

		testServer.CreateUserHandler(rec, req)
		assert.Equal(t, http.StatusInternalServerError, rec.Code)

	})
}

func TestCreateProductHandler(t *testing.T) {

	reqbody := ProductRequest{
		UserId:        1,
		ProductName:   "test",
		ProductDesc:   "testing",
		ProductImages: []string{"url1", "url2"},
		ProductPrice:  "100",
	}

	mockbyt, _ := json.Marshal(reqbody)

	t.Run("Success case", func(t *testing.T) {

		req := httptest.NewRequest(http.MethodPost, "/create-product", bytes.NewBuffer(mockbyt))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		mockservice.EXPECT().CreateProduct(gomock.Any()).Return(nil)

		testServer.CreateProductHandler(rec, req)
		assert.Equal(t, http.StatusCreated, rec.Code)
	})

	t.Run("failure case", func(t *testing.T) {

		req := httptest.NewRequest(http.MethodPost, "/create-product", bytes.NewBuffer(mockbyt))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		mockservice.EXPECT().CreateProduct(gomock.Any()).Return(errors.New("error from database"))

		testServer.CreateProductHandler(rec, req)
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
	})

}
