package main

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func main() {

	fmt.Println("Starting RabbitMQ Tutorial")

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	ch,err:=conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
		panic(err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"TestQueue", // name
		false,       // durable
		false,       // delete when unused
		false,       // exclusive
		false,       // no-wait
		nil,         // arguments
	)

	if err != nil {
		log.Fatalf("Failed to declare a queue: %v", err)
		panic(err)
	}

	fmt.Println(q)

	err=ch.Publish(
		"",     // exchange
		"TestQueue", // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte("Hello World!"),
		},
	)

	if err != nil {
		log.Fatalf("Failed to publish a message: %v", err)
		panic(err)
	}

	fmt.Println("Message published")


}
