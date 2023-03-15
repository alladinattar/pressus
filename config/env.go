package config

import (
	"github.com/rabbitmq/amqp091-go"
)

type Env struct {
	Config       *Config
	QueueChannel *amqp091.Channel
}

func NewEnv() *Env {

	return &Env{
		Config:       &Config{},
		QueueChannel: &amqp091.Channel{},
	}
}
