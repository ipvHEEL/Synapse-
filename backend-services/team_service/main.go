package main

import "github.com/rabbitmq/amqp091-go"

func main() {

	conn, err := amqp091.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return
	}
	ch, err := conn.Channel()
	if err != nil {
		return
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"team_events",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return
	}

	msgs, err := ch.Consume()

}
