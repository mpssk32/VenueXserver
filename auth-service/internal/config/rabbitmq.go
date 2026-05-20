package config

import (
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

var RabbitMQ *amqp.Connection
var RabbitChannel *amqp.Channel

func ConnectRabbitMQ() {

	var conn *amqp.Connection
	var err error

	for i := 0; i < 10; i++ {

		conn, err = amqp.Dial(
			"amqp://guest:guest@rabbitmq:5672/",
		)

		if err == nil {
			break
		}

		log.Println("Waiting for RabbitMQ...")

		time.Sleep(3 * time.Second)
	}

	if err != nil {
		log.Fatal(err)
	}

	channel, err := conn.Channel()

	if err != nil {
		log.Fatal(err)
	}

	_, err = channel.QueueDeclare(
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

	RabbitMQ = conn
	RabbitChannel = channel

	log.Println("RabbitMQ connected")
}