package queue

import (
	"github.com/pressus/models/presenters"
	"github.com/rabbitmq/amqp091-go"
)

type QueueRepo interface {
	SaveLinksToArticles([]presenters.ArticleLink) error
	PushArticleToResults(obj *presenters.ArticleObj) error
	GetTasks(tasks chan amqp091.Delivery)
}

type queueRepo struct {
	conn    *amqp091.Connection
	channel *amqp091.Channel
}

func NewQueueRepo(conn *amqp091.Connection, ch *amqp091.Channel) QueueRepo {
	return &queueRepo{conn, ch}
}
