package kafka

import (
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

// GetKafkaProducer initializes and returns a Kafka producer.
func GetKafkaProducer(broker string) (*kafka.Producer, error) {
	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": broker,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create Kafka producer: %w", err)
	}
	return producer, nil
}

// GetKafkaConsumer initializes and returns a Kafka consumer.
func GetKafkaConsumer(broker, groupID string, topics []string) (*kafka.Consumer, error) {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": broker,
		"group.id":          groupID,
		"auto.offset.reset": "earliest", // Change to "latest" if needed
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create Kafka consumer: %w", err)
	}

	// Subscribe to topics
	err = consumer.SubscribeTopics(topics, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to subscribe to topics: %w", err)
	}

	return consumer, nil
}

// PublishMessage sends a message to the given Kafka topic.
func PublishMessage(producer *kafka.Producer, topic, key, value string) error {
	deliveryChan := make(chan kafka.Event)

	// Produce message
	err := producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Key:            []byte(key),
		Value:          []byte(value),
	}, deliveryChan)
	if err != nil {
		return fmt.Errorf("failed to produce message: %w", err)
	}

	// Wait for delivery report
	e := <-deliveryChan
	m := e.(*kafka.Message)
	if m.TopicPartition.Error != nil {
		return fmt.Errorf("failed to deliver message: %v", m.TopicPartition.Error)
	}

	log.Printf("Message delivered to %v [%d] at offset %v", *m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
	return nil
}

// ConsumeMessages reads messages from the Kafka topic.
func ConsumeMessages(consumer *kafka.Consumer) {
	for {
		msg, err := consumer.ReadMessage(-1) // Timeout: -1 means wait indefinitely
		if err == nil {
			log.Printf("Received message: %s from topic %s\n", string(msg.Value), *msg.TopicPartition.Topic)
		} else {
			log.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}
}
