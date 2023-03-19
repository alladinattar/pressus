package repository

import (
	"context"
	amqp "github.com/rabbitmq/amqp091-go"
	"time"
)

func (d *queueRepo) PushArticleBody(body string) error {
	q, err := d.channel.QueueDeclare(
		"results", // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = d.channel.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "plain/text",
			Body:        []byte(body),
		})
	if err != nil {
		return err
	}
	return nil
}
