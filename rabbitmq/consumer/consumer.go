package main

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func main() {
	fmt.Println("Consumer Application")

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}

	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
		panic(err)
	}
	defer ch.Close()

	msgs, err := ch.Consume(
		"TestQueue", // queue
		"",          // consumer
		false,       // auto-ack
		false,       // exclusive
		false,       // no-local
		false,       // no-wait
		nil,         // args
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %v", err)
		panic(err)
	}

	go func() {
		for d := range msgs {
			log.Printf("Received message: %s\n", d.Body)
		}
	}()
	

	forever := make(chan bool)// for blocking,it can be replaced by select
	
	fmt.Println("Successfully connected to RabbitMQ")
	fmt.Println(" [*] - waiting for messages")
	
	<-forever
		

}
