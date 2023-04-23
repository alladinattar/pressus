package search

import (
	"github.com/pressus/config"
	"github.com/pressus/models/presenters"
	"github.com/rabbitmq/amqp091-go"
)

type EngineRepo interface {
	SaveArticle(obj presenters.ArticleObj) error
}

type engineRepo struct {
	conn    *amqp091.Connection
	channel *amqp091.Channel
	env     config.Env
}

func NewEngineRepo(conn *amqp091.Connection, ch *amqp091.Channel, env config.Env) EngineRepo {
	return &engineRepo{conn, ch, env}
}
