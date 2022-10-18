package rabitmq

import (
	"fmt"
	"log"
	"os"
	"strings"

	pkgEvent "always-learning/golang/try-on/rabbitmq/event"
	"always-learning/golang/try-on/rabbitmq/pkg"

	"github.com/streadway/amqp"
)

type rabbitmq struct {
	Connection *amqp.Connection
	Channel    *amqp.Channel
	events     pkgEvent.Events
}

func New(url string, event pkgEvent.Events) (*rabbitmq, error) {
	connection, err := amqp.Dial(url)

	if err != nil {
		log.Println(err, "Failed to connect to RabbitMQ")
		return nil, err
	}

	defer connection.Close()

	log.Printf("got Connection, getting Channel")

	return &rabbitmq{
		Connection: connection,
	}, nil
}

func (r *rabbitmq) Run() (*rabbitmq, error) {

	fmt.Println("running whell")
	channel, err := r.Connection.Channel()

	if err != nil {
		log.Println(err, "Failed to open a channel")
		return nil, err
	}

	defer channel.Close()

	r.Listen()
	return &rabbitmq{
		Channel: channel,
	}, nil
}

func (r *rabbitmq) Publish(exchange, exchangeType, routingKey string, body string) (err error) {

	log.Printf("xgot Channel, declaring %q Exchange (%q)", exchangeType, exchange)
	// if err = r.Channel.ExchangeDeclare(
	// 	exchange,     // name
	// 	exchangeType, // type
	// 	true,         // durable
	// 	false,        // auto-deleted
	// 	false,        // internal
	// 	false,        // noWait
	// 	nil,          // arguments
	// ); err != nil {
	// 	return fmt.Errorf("Error exchange Declare: %s", err)
	// }

	log.Printf("declared Exchange, publishing %d body (%q)", len(body), body)

	err = r.Channel.Publish(
		exchange,                  // exchange
		pkg.SeverityFrom(os.Args), // routing key
		false,                     // mandatory
		false,                     // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	fmt.Println("Failed to publish a message", err)

	log.Printf(" [x] Sent %s", body)

	return nil

}

// func (r *rabbitmq) Listen(exchange, exchangeType, queueName, key string) error {
func (r *rabbitmq) Listen() error {

	var forever chan struct{}
	// key = "sales"
	// queueName = key + "_queue"
	// log.Printf("declared Exchange, declaring Queue %q", queueName)
	// Declare a priority queue

	fmt.Println("listen")
	// for _, e := range r.events {

	queue, err := r.Channel.QueueDeclare(
		"delivery.delivered",             // name of the queue
		true,                             // durable
		false,                            // delete when unused
		false,                            // exclusive
		false,                            // noWait
		amqp.Table{"x-max-priority": 10}, // arguments
	)

	if err != nil {
		return nil
	}

	// // key += ".#"
	// log.Printf("declared Queue (%q %d messages, %d consumers)",
	// 	queue.Name, queue.Messages, queue.Consumers)

	// // Create a queue bind
	if err = r.Channel.QueueBind(
		queue.Name,           // name of the queue
		"delivery.delivered", // bindingKey
		"delivery",           // sourceExchange
		false,                // noWait
		nil,                  // arguments
	); err != nil {
		return nil
	}

	// // Set consumer prefetch count to 1, priority queue will not without this.
	// err = r.Channel.Qos(1, 0, false)
	if err != nil {
		log.Fatalf("basic.qos: %v", err)
	}

	// log.Printf("Queue bound to Exchange, starting Consume (consumer tag %q)", c.tag)
	deliveries, err := r.Channel.Consume(
		queue.Name, // name
		"",         // consumerTag,
		false,      // noAck
		false,      // exclusive
		false,      // noLocal
		false,      // noWait
		nil,        // arguments
	)
	if err != nil {
		fmt.Println("Queue Consume: ", err)
		return nil
	}
	fmt.Println("asd")
	// Handle messages
	go r.Handling("ORDER_DELIVERED", "delivery.delivered", deliveries)

	// }

	log.Println("The RabbitMQ is listening.")

	<-forever

	return nil
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
