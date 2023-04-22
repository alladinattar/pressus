package queue

import (
	"context"
	"encoding/json"
	"github.com/pressus/models/presenters"
	amqp "github.com/rabbitmq/amqp091-go"
	log "github.com/sirupsen/logrus"
	"time"
)

func (d *queueRepo) SaveLinksToArticles(articlesTitles []presenters.ArticleLink) error {
	q, err := d.channel.QueueDeclare(
		"tasks", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	for _, article := range articlesTitles {
		body, err := json.Marshal(article)
		if err != nil {
			return err
		}
		err = d.channel.PublishWithContext(ctx,
			"",     // exchange
			q.Name, // routing key
			false,  // mandatory
			false,  // immediate
			amqp.Publishing{
				ContentType: "application/json",
				Body:        body,
			})
		if err != nil {
			return err
		}
		log.Printf(" [x] Sent %s\n", body)
	}
	return nil
}
