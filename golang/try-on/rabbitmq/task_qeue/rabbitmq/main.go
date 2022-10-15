package rabitmq

import (
	"fmt"
	"log"
	"strings"
	"time"

	pkgEvent "always-learning/golang/try-on/rabbitmq/event"

	"github.com/streadway/amqp"
)

type rabbitmq struct {
	Connection *amqp.Connection
	Channel    *amqp.Channel
	events     pkgEvent.Events
	// deliveryEventHandler handler.DeliveryEventHandler
	// paymentEventHandler  handler.PaymentEventHandler
}

func New(url string, event pkgEvent.Events) (*rabbitmq, error) {
	connection, err := amqp.Dial(url)

	if err != nil {
		log.Println(err, "Failed to connect to RabbitMQ")
		return nil, err
	}

	defer connection.Close()

	// ch, err := conn.Channel()

	// defer ch.Close()

	channel, err := connection.Channel()

	if err != nil {
		log.Println(err, "Failed to open a channel")
		return nil, err
	}

	defer channel.Close()

	log.Printf("got Connection, getting Channel")

	return &rabbitmq{
		Connection: connection,
		Channel:    channel,
	}, nil
}

func (r *rabbitmq) Run() (err error) {

	// r.Channel, err = r.Connection.Channel()
	// if err != nil {
	// 	log.Println(err, "Failed to open a channel")
	// 	return err
	// }
	// defer r.Channel.Close()

	// log.Printf("got Connection, getting Channel")
	fmt.Println("running whell")
	r.Listen()

	return nil
}

func (r *rabbitmq) Publish(exchange, exchangeType, routingKey string, body string) error {

	// This function dials, connects, declares, publishes, and tears down,
	// all in one go. In a real service, you probably want to maintain a
	// long-lived connection as state, and publish against that.

	// log.Printf("dialing %q", amqpURI)

	// connection, err := amqp.Dial(amqpURI)
	// if err != nil {
	// return fmt.Errorf("Dial: %s", err)
	// }
	// defer connection.Close()

	// Create an exchange
	// Create an exchange
	log.Printf("got Channel, declaring %q Exchange (%q)", exchangeType, exchange)
	if err := r.Channel.ExchangeDeclare(
		exchange,     // name
		exchangeType, // type
		false,        // durable
		false,        // auto-deleted
		false,        // internal
		false,        // noWait
		nil,          // arguments
	); err != nil {
		return fmt.Errorf("Exchange Declare: %s", err)
	}

	log.Printf("declared Exchange, publishing %dB body (%q)", len(body), body)

	pri := []int{5, 1, 9}
	count := 0
	// This loop will procude item with priority 5, 1, 9, 5, 1, 9 with a time gap of 1 second
	for i := 0; ; i++ {
		routingKey := "sales.call"
		priority := pri[count]
		count++
		if count == 3 {
			count = 0
		}
		expire := "60000"
		log.Printf("publishing %dB body (%q), counter %d, expire %s, priority %d", len(body), body, i, expire, priority)
		if err := r.Channel.Publish(
			exchange,   // publish to an exchange
			routingKey, // routing to 0 or more queues
			false,      // mandatory
			false,      // immediate
			amqp.Publishing{
				Headers:         amqp.Table{},
				ContentType:     "text/plain",
				ContentEncoding: "",
				Body:            []byte(routingKey + " " + body + " " + fmt.Sprintf("priority=%d", priority) + " " + "expire=" + expire),
				DeliveryMode:    amqp.Persistent, // 1=non-persistent, 2=persistent
				Priority:        uint8(priority), // 0-9
				Expiration:      expire,
				// a bunch of application/implementation-specific fields
			},
		); err != nil {
			return fmt.Errorf("Exchange Publish: %s", err)
		}
		time.Sleep(time.Second * 1)
	}

}

// func (r *rabbitmq) Listen(exchange, exchangeType, queueName, key string) error {
func (r *rabbitmq) Listen() error {

	var forever chan struct{}
	// key = "sales"
	// queueName = key + "_queue"
	// log.Printf("declared Exchange, declaring Queue %q", queueName)
	// Declare a priority queue

	for _, e := range r.events {

		queue, err := r.Channel.QueueDeclare(
			"queueName",                      // name of the queue
			true,                             // durable
			false,                            // delete when unused
			false,                            // exclusive
			false,                            // noWait
			amqp.Table{"x-max-priority": 10}, // arguments
		)

		if err != nil {
			return nil
		}

		// key += ".#"
		log.Printf("declared Queue (%q %d messages, %d consumers)",
			queue.Name, queue.Messages, queue.Consumers)

		// Create a queue bind
		if err = r.Channel.QueueBind(
			queue.Name,     // name of the queue
			e.RoutingKey(), // bindingKey
			e.Exchange(),   // sourceExchange
			false,          // noWait
			nil,            // arguments
		); err != nil {
			return nil
		}

		// Set consumer prefetch count to 1, priority queue will not without this.
		err = r.Channel.Qos(1, 0, false)
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

		// Handle messages
		go r.Handling(e.Name(), e.RoutingKey(), deliveries)

	}

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
