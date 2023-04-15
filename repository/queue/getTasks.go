package queue

import (
	"github.com/rabbitmq/amqp091-go"
	log "github.com/sirupsen/logrus"
)

func (r *queueRepo) GetTasks(tasks chan amqp091.Delivery) {
	q, err := r.channel.QueueDeclare(
		TASKS_QUEUE, // name
		false,       // durable
		false,       // delete when unused
		false,       // exclusive
		false,       // no-wait
		nil,         // arguments
	)
	if err != nil {
		log.Error("Failed Queue Declare: ", err)
	}

	msgs, err := r.channel.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		log.Error("Failed cunsume: ", err)
	}

	var forever chan struct{}

	go func() {
		for d := range msgs {
			//task := &presenters.ArticleObj{}
			//err := json.Unmarshal(d.Body, &task)
			//if err != nil {
			//	log.Error("Failed unmarshall task: ", err.Error())
			//}
			tasks <- d
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
