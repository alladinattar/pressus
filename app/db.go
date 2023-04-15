package app

import (
	"fmt"
	"github.com/pressus/config"
	amqp "github.com/rabbitmq/amqp091-go"
)

func QueueConnection(cfg *config.Env) (*amqp.Connection, *amqp.Channel, error) {
	connString := fmt.Sprintf("amqp://%s:%s@%s:5672/", cfg.Config.Queue.Username,
		cfg.Config.Queue.Password, cfg.Config.Queue.Ip)
	conn, err := amqp.Dial(connString)
	if err != nil {
		return nil, nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, nil, err
	}

	return conn, ch, nil
}
