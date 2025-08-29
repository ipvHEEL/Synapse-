// send.go
package main

import (
	"log"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Не удалось подключиться к RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Не удалось открыть канал")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"team_events",
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Не удалось объявить очередь")

	body := `{
		"event": "team.created",
		"data": {
			"id": "t-123",
			"name": "Frontend Team"
		}
	}`

	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(body),
		})
	failOnError(err, "Не удалось отправить сообщение")

	log.Printf("Отправлено: %s", body)
}
