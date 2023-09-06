package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	temp "transcoding-service/consumer"

	"github.com/nsqio/go-nsq"
)

func main() {
	config := nsq.NewConfig()
	consumer, err := nsq.NewConsumer("transcoding", "ch01", config)
	if err != nil {
		log.Fatal(err)
	}
	consumer.AddHandler(&temp.TranscodingHandler{})
	err = consumer.ConnectToNSQLookupd("localhost:4161")
	if err != nil {
		log.Fatal(err)
	}
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan
	consumer.Stop()
}
