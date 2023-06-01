package config

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

var (
	RabbitmqChPubl	*amqp.Channel
	RabbitmqChCons	*amqp.Channel
)

type RabbitmqConf struct{}

func (cfg *RabbitmqConf) RabbitmqMakeConn() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Panicf("%s: %s", "Failed to connect to RabbitMQ", err)
	}

	chPubl, err := conn.Channel()
	if err != nil {
		log.Panicf("%s: %s", "Failed to open a RabbitMQ publisher channel", err)
	}
	chCons, err := conn.Channel()
	if err != nil {
		log.Panicf("%s: %s", "Failed to open a RabbitMQ consumer channel", err)
	}

	RabbitmqChPubl = chPubl
	RabbitmqChCons = chCons
}