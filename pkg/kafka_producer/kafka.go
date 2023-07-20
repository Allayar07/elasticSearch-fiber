package kafka_producer

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

const bootstrapServers = "bootstrap.servers"

func ProducerQueues() (*kafka.Producer, error) {
	producer, err := kafka.NewProducer(&kafka.ConfigMap{bootstrapServers: "localhost:9092"})
	if err != nil {
		return nil, err
	}
	fmt.Println("successfully connected to kafka")
	return producer, nil
}
