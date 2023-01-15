package api

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

// Global variables
var Connection *amqp.Connection
