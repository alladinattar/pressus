package queue

import (
	"github.com/pressus/models/presenters"
	"github.com/rabbitmq/amqp091-go"
)

type QueueRepo interface {
	SaveLinksToArticles([]presenters.ArticleObj) error
	PushArticleBody(string) error
	GetTasks(tasks chan presenters.ArticleObj)
}

type queueRepo struct {
	conn    *amqp091.Connection
	channel *amqp091.Channel
}

func NewQueueRepo(conn *amqp091.Connection, ch *amqp091.Channel) QueueRepo {
	return &queueRepo{conn, ch}
}
