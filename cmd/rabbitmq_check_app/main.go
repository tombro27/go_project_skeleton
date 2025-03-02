package main

import (
	"log"

	"github.com/tombro27/go_project_skeleton/internal/config"
	"github.com/tombro27/go_project_skeleton/pkg/rabbitmq"
)

var cfg *config.Config

func main() {
	var err error
	cfg, err = config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	testPublishMessage("ser1")

	testConsumeMessage("ser2")
}

func testPublishMessage(key string) {
	// Connect to RabbitMQ
	conn, err := rabbitmq.GetRabbitMQConnection(cfg.RabbitMQ[key])
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// Get channel
	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	defer ch.Close()

	// Publish message
	err = rabbitmq.PublishMessage(ch, "test-queue", "Hello, RabbitMQ!")
	if err != nil {
		log.Fatal(err)
	}
}

func testConsumeMessage(key string) {
	// Connect to RabbitMQ
	conn, err := rabbitmq.GetRabbitMQConnection(cfg.RabbitMQ[key])
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// Get channel
	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	defer ch.Close()

	// Consume messages
	msgs, err := rabbitmq.ConsumeMessages(ch, "test-queue")
	if err != nil {
		log.Fatal(err)
	}

	// Process messages
	for msg := range msgs {
		log.Printf("Received message: %s", string(msg.Body))
	}
}
