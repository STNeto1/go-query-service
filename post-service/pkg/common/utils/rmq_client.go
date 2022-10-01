package utils

import amqp "github.com/rabbitmq/amqp091-go"

func CreateRmqConnection(addr string) *amqp.Connection {
	conn, err := amqp.Dial(addr)
	FailOnError(err, "Failed to connect to RabbitMQ")

	return conn
}

func CreateRmqChannel(conn *amqp.Connection) *amqp.Channel {
	ch, err := conn.Channel()
	FailOnError(err, "Failed to open a channel")

	return ch
}

func CreateRmqQueue(ch *amqp.Channel, name string) amqp.Queue {
	q, err := ch.QueueDeclare(
		name,  // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	FailOnError(err, "Failed to open a channel")

	return q
}
