package search

import (
	"github.com/pressus/models/presenters"
	"github.com/rabbitmq/amqp091-go"
)

type EngineRepo interface {
	SaveArticle(obj presenters.ArticleObj)
}

type engineRepo struct {
	conn    *amqp091.Connection
	channel *amqp091.Channel
}

func NewEngineRepo(conn *amqp091.Connection, ch *amqp091.Channel) EngineRepo {
	return &engineRepo{conn, ch}
}
