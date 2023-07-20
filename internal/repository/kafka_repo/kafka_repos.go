package kafka_repo

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"log"
)

type KafkaRepo struct {
	kafka *kafka.Producer
}

func NewKafkaRepo(kafka *kafka.Producer) *KafkaRepo {
	return &KafkaRepo{
		kafka: kafka,
	}
}

func (r *KafkaRepo) PublishTopic(topic, value string) error {
	deliveryChan := make(chan kafka.Event)
	defer close(deliveryChan)
	if err := r.kafka.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(value),
	},
		deliveryChan,
	); err != nil {
		fmt.Println("error here: ", err)
		return err
	}
	report := <-deliveryChan
	kafkaMsg := report.(*kafka.Message)
	if kafkaMsg.TopicPartition.Error != nil {
		fmt.Println("error here")
		return kafkaMsg.TopicPartition.Error
	}
	log.Printf("Delivery succeed\nTopic:%s. Value:%s\n", "books", value)
	return nil
}
