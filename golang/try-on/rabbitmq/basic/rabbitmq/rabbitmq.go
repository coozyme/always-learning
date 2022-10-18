package rabbitmq

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/streadway/amqp"
)

type rabbitmq struct {
	Connection *amqp.Connection
}

func New() (*rabbitmq, error) {

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Println("Failed Initializing Broker Connection", err)
		return nil, err
	}

	// Let's start by opening a channel to our RabbitMQ instance
	// over the connection we have already established

	return &rabbitmq{
		Connection: conn,
		// Channel:    ch,
	}, nil
}

func (r *rabbitmq) Run() error {

	r.Listen()

	return nil
}

func (r *rabbitmq) Publish(exchange, routingKey string, body string) {

	ch, err := r.Connection.Channel()
	if err != nil {
		fmt.Println("err channel", err)
		return
	}
	defer ch.Close()

	err = ch.ExchangeDeclare(
		exchange, // name
		"topic",  // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)

	log.Println(err, "Failed to declare an exchange")

	q, err := ch.QueueDeclare(
		exchange,
		false,
		false,
		false,
		false,
		nil,
	)
	// We can print out the status of our Queue here
	// this will information like the amount of messages on
	// the queue
	fmt.Println(q)
	// Handle any errors if we were unable to create the queue
	if err != nil {
		fmt.Println("err QueueDeclare", err)
		return
	}

	// attempt to publish a message to the queue!
	err = ch.Publish(
		exchange,
		routingKey,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		},
	)

	if err != nil {
		fmt.Println("Err Publish", err)
		return
	}
	fmt.Println("Successfully Published Message to Queue")
}

func (r *rabbitmq) Listen() {
	ch, err := r.Connection.Channel()
	if err != nil {
		fmt.Println("err channel", err)
		return
	}
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"ORDER_DELIVERED", // name
		"topic",           // type
		true,              // durable
		false,             // auto-deleted
		false,             // internal
		false,             // no-wait
		nil,               // arguments
	)
	if err != nil {
		log.Println(err, "Failed to declare an exchange")
	}

	q, err := ch.QueueDeclare(
		"ORDER_DELIVERED", // name
		false,             // durable
		false,             // delete when unused
		true,              // exclusive
		false,             // no-wait
		nil,               // arguments
	)

	if err != nil {
		log.Println(err, "Failed to declare a queue")
	}

	if len(os.Args) < 2 {
		log.Printf("Usage: %s [binding_key]...", os.Args[0])
		os.Exit(0)
	}

	for _, s := range os.Args[1:] {
		log.Printf("Binding queue %s to exchange %s with routing key %s",
			q.Name, "logs_topic", s)

		err = ch.QueueBind(
			q.Name,       // queue name
			s,            // routing key
			"logs_topic", // exchange
			false,
			nil)

		if err != nil {
			log.Println(err, "Failed to bind a queue")
		}
	}

	msgs, err := ch.Consume(
		q.Name,            // queue
		"ORDER_DELIVERED", // consumer
		true,              // auto ack
		false,             // exclusive
		false,             // no local
		false,             // no wait
		nil,               // args
	)

	if err != nil {
		log.Println(err, "Failed to register a consumer")
	}

	var forever chan struct{}

	go r.handling("order.paid", "s", msgs)

	log.Println(" [*] Waiting for logs. To exit press CTRL+C")

	<-forever

}

func (r *rabbitmq) handling(eventName, routingKey string, msgs <-chan amqp.Delivery) {

	for d := range msgs {
		log.Println("msg", d.Body)
		if strings.Contains(eventName, "order.paid") {
			log.Println("Handling Order Paid", d.Body)
		} else if strings.Contains(eventName, "order.delivered") {
			log.Println("Handling Order Paid", d.Body)
		}
	}
}
