package main

import (
	"log"

	pkgEvent "always-learning/golang/try-on/rabbitmq/event"
	rabitmq "always-learning/golang/try-on/rabbitmq/logs_topic/rabbitmq"
)

func main() {
	ms, err := rabitmq.New(pkgEvent.Events{
		pkgEvent.New("ORDER_PAID", "payment.paid", "order"),
		pkgEvent.New("ORDER_DELIVERED", "delivery.delivered", "delivery"),
	})

	if err != nil {
		log.Panic(err)
	}

	go func() {
		log.Panic(ms.Run())
	}()
	log.Println(ms.Listen())
}
