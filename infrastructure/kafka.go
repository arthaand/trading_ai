package infrastructure

import (
	"strings"

	kafka "andromeda.ottopay.id/ottopoint/ottopoint-kafka"
	s "andromeda.ottopay.id/ottopoint/ottopoint-shared"
	"github.com/Shopify/sarama"
)

var (
	consumer sarama.Consumer
	producer sarama.AsyncProducer
)

func GetConsumer() sarama.Consumer {
	return consumer
}

func GetProducer() sarama.AsyncProducer {
	return producer
}

func InitKafkaConnection() error {
	brokers := strings.Split(s.Config.Kafka.Host, ",")
	var err error

	producer, err = kafka.GetProducer(brokers)
	if err != nil {
		return err
	}
	consumer, err = kafka.GetConsumer(brokers)
	if err != nil {
		return err
	}
	return nil
}
