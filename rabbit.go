package rabbit

import (
	sillyKits "github.com/irealing/silly-kits"
	amqp "github.com/rabbitmq/amqp091-go"
	"time"
)

func connectMQ(uri string, retry int) (*amqp.Connection, error) {
	return sillyKits.Retry(func() (*amqp.Connection, error) {
		return amqp.Dial(uri)
	}, retry, time.Second)
}
