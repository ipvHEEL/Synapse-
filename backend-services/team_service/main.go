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

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Не удалось подписаться на очередь")

	log.Println(" Ожидаем сообщения о командах. Для выхода нажмите CTRL+C")

	forever := make(chan bool)
	go func() {
		for msg := range msgs {
			log.Printf("📩 Получено: %s", msg.Body)
		}
	}()

	<-forever
}
