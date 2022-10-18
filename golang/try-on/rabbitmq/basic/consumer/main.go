package main

import (
	"always-learning/golang/try-on/rabbitmq/basic/rabbitmq"
	"log"
)

func main() {
	rabbitmq, err := rabbitmq.New()

	if err != nil {
		log.Fatal(err)
	}
	go func() {
		log.Print(rabbitmq.Run())
	}()
	rabbitmq.Listen()
}
