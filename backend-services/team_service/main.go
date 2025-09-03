package main

import (
	"encoding/json"
	"log"

	"github.com/streadway/amqp"
)

func failonError(err error, msg string) {
	if err != nil {
		log.Fatal("%s: %s", msg, err)
	}
}

func sendForQueue(TeamName string) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failonError(err, "Не удалось запустить Rabbit")
	defer conn.Close()

	ch, err := conn.Channel()
	failonError(err, "Не удалось запусить канал")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"team_events",
		false,
		false,
		false,
		false,
		nil,
	)
	failonError(err, "Не удалось звпустить очередь")

	event := map[string]interface{}{
		"event": "team_created",
		"data": map[string]string{
			"id":   "1",
			"name": TeamName,
		},
	}

	body, err := json.Marshal(event)

	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})
	failonError(err, "Не удалось отправить сообщение в очередь")

	log.Printf("Отправлено: %s", body)
}

func main() {
	// server := gin.Default()
	// server.POST("/team", sendForQueue)
	sendForQueue("frontend team")
}
