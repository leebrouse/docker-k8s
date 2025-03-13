package main

import (
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

	msg, err := ch.Consume(
		"hello",
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	//consume the message
	var forever chan struct{}
	go func() {
		for v := range msg {
			log.Println("Receive a message:", string(v.Body))
			// v.Ack(false)
		}
	}()
	<-forever
}
