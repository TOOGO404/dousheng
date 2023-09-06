package main

import (
	"log"

	"github.com/nsqio/go-nsq"
)

const topicName = "transcoding"

func main() {
	config := nsq.NewConfig()
	producer, err := nsq.NewProducer("localhost:4161", config)
	if err != nil {
		log.Fatal(err)
	}
	messageBody := []byte("hello")

	err = producer.Publish(topicName, messageBody)
	if err != nil {
		log.Fatal(err)
	}
	producer.Stop()
}
