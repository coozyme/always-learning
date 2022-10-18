package rabitmq

import (
	"fmt"
	"log"
	"os"
	"strings"

	pkgEvent "always-learning/golang/try-on/rabbitmq/event"

	"github.com/streadway/amqp"
)

type rabbitmq struct {
	connection *amqp.Connection
	events     pkgEvent.Events
}

func New(event pkgEvent.Events) (*rabbitmq, error) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")

	if err != nil {
		log.Println(err, "Failed to connect to RabbitMQ")
		return nil, err
	}

	return &rabbitmq{
		connection: conn,
	}, nil
}

func (r *rabbitmq) Run() (err error) {

	r.Listen()

	return nil
}

func (r *rabbitmq) Publish(exchange, routingKey string, body string) {

	ch, err := r.connection.Channel()

	if err != nil {
		log.Println("Failed to open a channel", err)
		return
	}

	defer ch.Close()

	log.Printf("gotxc Connection, getting Channel")

	log.Printf("got Channelxx, declaring Exchange (%q)", exchange)
	if err = ch.ExchangeDeclare(
		exchange, // name
		"topic",  // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // noWait
		nil,      // arguments
	); err != nil {
		fmt.Errorf("Error exchange Declare: %s", err)
		return
	}

	if err != nil {
		log.Println("Err ExchangeDeclare", err)
		return
	}

	log.Printf("declared Exchange, publishing %dB body (%q)", len(body), body)

	// q, err := ch.QueueDeclare(
	// 	exchange,
	// 	false,
	// 	false,
	// 	false,
	// 	false,
	// 	nil,
	// )
	// We can print out the status of our Queue here
	// this will information like the amount of messages on
	// the queue
	// fmt.Println(q.)
	// Handle any errors if we were unable to create the queue
	if err != nil {
		fmt.Println("err QueueDeclare", err)
		return
	}
	// valueBody := pkg.BodyFrom(body)

	err = ch.Publish(
		exchange,   // exchange
		routingKey, // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	if err != nil {
		log.Println("Failed to publish a message", err)
		return
	}

	log.Printf(" [x] Sent %s", body)
	return

}

// func (r *rabbitmq) Listen(exchange, exchangeType, queueName, key string) error {
func (r *rabbitmq) Listen() (err error) {
	channel, err := r.connection.Channel()

	if err != nil {
		log.Println(err, "Failed to open a channel")
		return
	}

	defer channel.Close()

	var forever chan struct{}

	// err = channel.ExchangeDeclare(
	// 	"delivery", // name
	// 	"topic",    // type
	// 	true,       // durable
	// 	false,      // auto-deleted
	// 	false,      // internal
	// 	false,      // no-wait
	// 	nil,        // arguments
	// )

	// if err != nil {
	// 	log.Println("Err Failed to declare a queue :", err)
	// 	return
	// }

	// q, err := channel.QueueDeclare(
	// 	"",    // name
	// 	false, // durable
	// 	false, // delete when unused
	// 	true,  // exclusive
	// 	false, // no-wait
	// 	nil,   // arguments
	// )
	// if err != nil {
	// 	log.Println("Err Failed to declare a queue :", err)
	// }

	// if len(os.Args) < 2 {
	// 	log.Printf("Usage: %s [binding_key]...", os.Args[0])
	// 	os.Exit(0)
	// }

	// for _, s := range os.Args[1:] {
	// 	err = channel.QueueBind(
	// 		q.Name,     // queue name
	// 		s,          // routing key
	// 		"delivery", // exchange
	// 		false,
	// 		nil)
	// 	if err != nil {
	// 		log.Println("Err QueueBind :", err)
	// 		return
	// 	}
	// }

	// msgs, err := channel.Consume(
	// 	q.Name, // queue
	// 	"",     // consumer
	// 	true,   // auto ack
	// 	false,  // exclusive
	// 	false,  // no local
	// 	false,  // no wait
	// 	nil,    // args
	// )
	// if err != nil {
	// 	log.Println("Err consume :", err)
	// 	return
	// }

	for _, e := range r.events {
		log.Println("e", e)
		err = channel.ExchangeDeclare("delivery", "topic", false, false, false, false, nil)

		if err != nil {
			return err
		}

		q, err := channel.QueueDeclare("delivery", true, false, false, false, nil)
		if err != nil {
			return err
		}

		if len(os.Args) < 2 {
			log.Printf("Usage: %s [binding_key]...", os.Args[0])
			os.Exit(0)
		}

		for _, s := range os.Args[1:] {
			err = channel.QueueBind(
				q.Name,     // queue name
				s,          // routing key
				"delivery", // exchange
				false,
				nil)
			log.Println("Failed to bind a queue", err)
		}

		msgs, err := channel.Consume(q.Name, "", true, false, false, false, nil)
		if err != nil {
			return err
		}

		// go r.handling(e.Name(), msgs)
		go r.Handling(e.Name(), e.Exchange(), msgs)
	}

	// Handle messages
	// go r.Handling("ORDER_DELIVERED", "payment.paid", msgs)

	// }

	log.Println("The RabbitMQ is listening.")

	<-forever
	return
}

func (r *rabbitmq) Handling(eventName, routingKey string, msgs <-chan amqp.Delivery) {

	for d := range msgs {
		// orderUUID, err := uuid.Parse(string(d.Body))
		// if err != nil {
		// 	log.Println(err)

		// 	return
		// }

		// switch eventName {
		// case "order.paid":
		// 	log.Println(r.paymentEventHandler.ReceiveOrderHasBeenPaid(event.NewOrderPaid(orderUUID)))
		// case "order.delivered":
		// 	log.Println(r.deliveryEventHandler.ReceiveOrderHasBeenDelivered(event.NewOrderDelivered(orderUUID)))
		// }

		if strings.Contains(eventName, "order.paid") {
			log.Println("Handling Order Paid", d)
		} else if strings.Contains(eventName, "order.delivered") {
			log.Println("Handling Order Paid", d)
		}
	}
}
