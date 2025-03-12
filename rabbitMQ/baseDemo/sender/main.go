package main

import (
	"context"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s:%s", msg, err)
	}
}

func main() {
	//connect rabbitMQ
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "connecting rabbitmq failed")
	defer conn.Close()

	// create channel
	ch, err := conn.Channel()
	failOnError(err, "create channel failed")
	defer ch.Close()

	//create queue
	q, err := ch.QueueDeclare(
		"hello",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "create queue failed")

	for range 5 {
		//post message
		body := "fuck world"
		ch.PublishWithContext(
			context.Background(),
			"",
			q.Name,
			false,
			false,
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(body),
			},
		)
	}

}
