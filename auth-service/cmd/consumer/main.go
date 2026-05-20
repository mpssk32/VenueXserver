package main

import (
	"auth-service/internal/config"
	"log"
)

func main() {

	config.ConnectRabbitMQ()

	msgs, err := config.RabbitChannel.Consume(
		"notifications",
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Consumer started")

	forever := make(chan bool)

	go func() {

		for msg := range msgs {
			log.Printf(
				"Notification received: %s",
				msg.Body,
			)
		}
	}()

	<-forever
}