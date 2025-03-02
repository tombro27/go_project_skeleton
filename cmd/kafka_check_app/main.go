package main

import (
	"log"

	"github.com/tombro27/go_project_skeleton/internal/config"
	"github.com/tombro27/go_project_skeleton/pkg/kafka"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize Kafka Producer
	producer, err := kafka.GetKafkaProducer(cfg.Kafka["serv1"].Broker)
	if err != nil {
		log.Fatal(err)
	}
	defer producer.Close()

	// Publish message
	err = kafka.PublishMessage(producer, "test-topic", "key1", "Valeu1")
	if err != nil {
		log.Fatal(err)
	}

	// Initialize Kafka Consumer
	consumer, err := kafka.GetKafkaConsumer(cfg.Kafka["serv2"].Broker, "goproject-group", []string{"test-topic"})
	if err != nil {
		log.Fatal(err)
	}
	defer consumer.Close()

	// Read messages
	kafka.ConsumeMessages(consumer)
}
