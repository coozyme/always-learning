package main

import (
	"fmt"
	"log"

	pkgEvent "always-learning/golang/try-on/rabbitmq/event"
	rabitmq "always-learning/golang/try-on/rabbitmq/task_qeue/rabbitmq"
)

func main() {
	ms, err := rabitmq.New("amqp://guest:guest@localhost:5672/", pkgEvent.Events{
		pkgEvent.New("ORDER_PAID", "payment.paid", "order"),
		pkgEvent.New("ORDER_DELIVERED", "delivery.delivered", "delivery"),
	})

	if err != nil {
		log.Panic(err)
	}

	go func() {
		err := ms.Run()
		// ms.Publish("order", "", "payment.paid", `{"paymentID": "PAY001", "orderUUID": "ORD8712", "invoiceNo": "INV9761"}`)
		// ms.Publish("order", "", "payment.paid", `{"paymentID": "PAY001", "orderUUID": "ORD8712", "invoiceNo": "INV9761"}`)
		// ms.Publish("order", "", "payment.paid", `{"paymentID": "PAY001", "orderUUID": "ORD8712", "invoiceNo": "INV9761"}`)
		log.Panic(ms.Run())

		ms.Listen()
		if err != nil {
			fmt.Println("err run: ", err)
		}
	}()
}
