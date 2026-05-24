package main

import (
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {

var conn *amqp.Connection
var err error

	for {

		conn, err = amqp.Dial(
			"amqp://guest:guest@concert_rabbitmq:5672/",
		)

		if err == nil {
			break
		}

		log.Println("Waiting for RabbitMQ...")

		time.Sleep(3 * time.Second)
	}

	defer conn.Close()

	ch, err := conn.Channel()

	if err != nil {
		log.Fatal(err)
	}

	defer ch.Close()

	q, err := ch.QueueDeclare(
		"notifications",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Fatal(err)
	}

	msgs, err := ch.Consume(
		q.Name,
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

	log.Println("Notification service started")

	forever := make(chan bool)

	go func() {

		for d := range msgs {

			log.Printf(
				"Received notification: %s",
				d.Body,
			)
		}
	}()

	<-forever
}