package rabbitmq

import (
	"context"

	amqp "github.com/rabbitmq/amqp091-go"
)

// criando um canal no RabbitMQ
func OpenChannel() (*amqp.Channel, error) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}

	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}

	return ch, nil
}

func Consume(ch *amqp.Channel, out chan<- amqp.Delivery, queue string) error {
	msgs, err := ch.Consume(
		queue,
		"go-consumer",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}
	for msg := range msgs {
		out <- msg
	}
	return nil
}

// ch.Publish is deprecated: Use PublishWithContext instead
func Publish(ch *amqp.Channel, body string, exchange string) error {
	err := ch.Publish(
		exchange, // Exchange name
		"",       // Routing key
		false,    // Mandatory
		false,    // Immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		},
	)
	if err != nil {
		return err
	}
	return nil
}

// método não deprecated e utilizar dessa forma
func PublishWithContext(ctx context.Context, ch *amqp.Channel, body string, exchange string) error {
	err := ch.PublishWithContext(ctx,
		exchange, // Exchange name
		"",       // Routing key
		false,    // Mandatory
		false,    // Immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	if err != nil {
		return err
	}
	return nil
}
