package main

import (
	pkgEvent "always-learning/golang/try-on/rabbitmq/event"
	rabitmq "always-learning/golang/try-on/rabbitmq/logs_topic/rabbitmq"
	"log"
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
	ms.Publish("delivery", "payment.paid1", `{"paymentID": "PAY001", "orderUUID": "ORD8712", "invoiceNo": "INV9761"}`)
	// ms.Publish("delivery", "payment.paid2", `{"paymentID": "PAY001", "orderUUID": "ORD8712", "invoiceNo": "INV9761"}`)
	// ms.Publish("delivery", "payment.paid3", `{"paymentID": "PAY001", "orderUUID": "ORD8712", "invoiceNo": "INV9761"}`)
}
