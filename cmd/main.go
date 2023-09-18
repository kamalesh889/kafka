package main

import (
	"Kafka/api"
	"net/http"

	"fmt"
	"log"
)

func main() {
	log.Println("Kafka Producer Consumer service")

	db, err := api.InitializeDB()
	if err != nil {
		log.Panicln("Error in Connecting to Database:", err)
	}

	server, err := api.NewServer(db)
	if err != nil {
		log.Panicln("Error in creating server:", err)
	}

	// starting a separate thread for the consumer
	// we can create a differnt main file for the same

	go func() {
		server.GetProductFromKafka()
	}()

	mux := api.Router(server)
	http.Handle("/", mux)
	http.ListenAndServe(fmt.Sprintf(":%s", "8080"), mux)
}
