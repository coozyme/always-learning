package main

import (
	pkgEvent "always-learning/golang/try-on/rabbitmq/event"
	rabitmq "always-learning/golang/try-on/rabbitmq/taskk/rabbitmq"
	"log"
)

func main() {
	ms, err := rabitmq.New("amqp://guest:guest@localhost:5672/", pkgEvent.Events{
		pkgEvent.New("ORDER_PAID", "payment.paid", "order"),
		pkgEvent.New("ORDER_DELIVERED", "delivery.delivered", "delivery"),
	})

	if err != nil {
		log.Panic(err)
	}

	// ms.Publish("order2", "a2", "payment.paid2", `{"paymentID": "PAY001", "orderUUID": "ORD8712", "invoiceNo": "INV9761"}`)
	// ms.Publish("order3", "a3", "payment.paid3", `{"paymentID": "PAY001", "orderUUID": "ORD8712", "invoiceNo": "INV9761"}`)

	ms.Publish("order1", "a1", "payment.paid1", `{"paymentID": "PAY001", "orderUUID": "ORD8712", "invoiceNo": "INV9761"}`)
	go func() {
		_, err = ms.Run()
		log.Println("err run", err)
	}()

}
