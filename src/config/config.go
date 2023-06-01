package config

import (
	jobs "github.com/williamluisan/golang_pdfgenerator/jobs"
)

type Config struct{}

func (cfg *Config) InitRabbitmq() {
	var rabbitmqJob jobs.RabbitmqJob

	// make connection
	var rabbitmqConf RabbitmqConf
	rabbitmqConf.RabbitmqMakeConn()

	// declare rabbitmq queue publisher
	rabbitmqJob.DeclareQueue(RabbitmqChPubl, "queue2")

	// declare rabbitmq queue consumer
	// rabbitmqJob.DeclareQueue(RabbitmqChCons, "queue2")
}