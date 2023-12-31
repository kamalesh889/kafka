package api

import (
	"Kafka/model"
	"compress/gzip"
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/segmentio/kafka-go"
)

type Service interface {
	CreateProduct(*ProductRequest) error
	CreateUser(*UserRequest) (uint64, error)
	GetProductFromKafka(bool)
	DownloadAndCompressImage(string, string, bool) (string, error)
}

const (
	imagedownload       = "downloadimage"
	compressedimage     = "compressedimage"
	testimagedownload   = "testdownloadimage"
	testcompressedimage = "testcompressimage"
)

func (s *server) CreateProduct(productDetails *ProductRequest) error {

	// Asuuming for creating prdouct , user should exists in our record

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

func (s *server) GetProductFromKafka(env bool) {

	msg, err := s.consumer.ReadMessage(context.Background())
	if err != nil {
		log.Println(err)
		return
	}

	productid, err := strconv.Atoi(string(msg.Value))
	if err != nil {
		log.Println(err)
		return
	}

	product, err := s.db.GetProduct(uint64(productid))
	if err != nil {
		log.Println(err)
		return
	}

	var compressedimagefiles []string

	// we can add goroutines here to handle all the urls concurrenly

	for i, val := range product.ProductImages {

		imagename := fmt.Sprintf("%s_%d.jpg", product.ProductName, i)

		path, err := s.DownloadAndCompressImage(val, imagename, env)
		if err != nil {
			log.Println(err)
		}

		compressedimagefiles = append(compressedimagefiles, path)

	}

	updatedProduct := &model.Product{
		ProductId:               product.ProductId,
		CompressedProductImages: compressedimagefiles,
		UpdatedAt:               time.Now(),
	}

	err = s.db.UpdateProduct(updatedProduct)
	if err != nil {
		log.Println(err)
		return
	}

}

func (s *server) DownloadAndCompressImage(imageurl string, imagename string, env bool) (string, error) {

	var imageLoc, compImageLoc string

	if env {
		imageLoc = imagedownload
		compImageLoc = compressedimage
	} else {
		imageLoc = testimagedownload
		compImageLoc = testcompressedimage
	}

	req, err := http.NewRequest(http.MethodGet, imageurl, nil)
	if err != nil {
		log.Println("Error in creating the api request", err)
		return "", err
	}

	resp, err := s.httpClient.Do(req)
	if err != nil {
		log.Println("Error in making the api request", err)
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("HTTP request failed with status: %v", resp.Status)
	}

	filePath := filepath.Join(imageLoc, imagename)

	out, err := os.Create(filePath)
	if err != nil {
		log.Println("Error in creating the image file", err)
		return "", err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return "", err
	}

	compressedFilename := fmt.Sprintf("compressed_%s.gz", imagename)

	inputFile, err := os.Open(filePath)
	if err != nil {
		log.Println("Error in opening the image file", err)
		return "", err
	}
	defer inputFile.Close()

	compfilePath := filepath.Join(compImageLoc, compressedFilename)

	outputFile, err := os.Create(compfilePath)
	if err != nil {
		log.Println("Error in creating the compressed image file", err)
		return "", err
	}
	defer outputFile.Close()

	gw := gzip.NewWriter(outputFile)
	defer gw.Close()

	_, err = io.Copy(gw, inputFile)
	if err != nil {
		return "", err
	}

	absPath, err := filepath.Abs(compressedFilename)
	if err != nil {
		fmt.Printf("Error getting absolute path: %v\n", err)
		return "", err
	}

	if !env {
		os.Remove(filePath)
		os.Remove(compfilePath)
	}

	return absPath, nil

}
