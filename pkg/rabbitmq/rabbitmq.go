package rabbitmq

import (
	"fmt"
	"log"

	"github.com/rabbitmq/amqp091-go"
	"github.com/tombro27/go_project_skeleton/internal/config"
)

// GetRabbitMQConnection establishes a connection to the RabbitMQ server.
// It creates RabbitMQ connection URL dynamically using the host, port, username, password, and vhost parameters from RMQConfig object.
func GetRabbitMQConnection(rmqcfg config.RMQConfig) (*amqp091.Connection, error) {
	// Construct connection URL
	connURL := fmt.Sprintf("amqp://%s:%s@%s:%d/%s", rmqcfg.Username, rmqcfg.Password, rmqcfg.Host, rmqcfg.Port, rmqcfg.VHost)

	// Connect to RabbitMQ
	conn, err := amqp091.Dial(connURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to RabbitMQ: %w", err)
	}

	log.Println("Connected to RabbitMQ successfully!")
	return conn, nil
}

// GetRabbitMQConnectionByURL establishes a connection to the RabbitMQ server using connection URL.
func GetRabbitMQConnectionByURL(url string) (*amqp091.Connection, error) {
	conn, err := amqp091.Dial(url)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to RabbitMQ: %w", err)
	}
	return conn, nil
}

// PublishMessage publishes a message to a RabbitMQ queue.
func PublishMessage(ch *amqp091.Channel, queueName string, body string) error {
	_, err := ch.QueueDeclare(
		queueName, // Queue name
		true,      // Durable
		false,     // Auto-delete
		false,     // Exclusive
		false,     // No-wait
		nil,       // Arguments
	)
	if err != nil {
		return fmt.Errorf("failed to declare queue: %w", err)
	}

	err = ch.Publish(
		"",        // Exchange
		queueName, // Routing key (queue name)
		false,     // Mandatory
		false,     // Immediate
		amqp091.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		},
	)
	if err != nil {
		return fmt.Errorf("failed to publish message: %w", err)
	}

	log.Printf("Published message to queue %s: %s", queueName, body)
	return nil
}

// ConsumeMessages reads messages from a RabbitMQ queue.
func ConsumeMessages(ch *amqp091.Channel, queueName string) (<-chan amqp091.Delivery, error) {
	msgs, err := ch.Consume(
		queueName, // Queue name
		"",        // Consumer name
		true,      // Auto-ack (set to false if manual acknowledgment is needed)
		false,     // Exclusive
		false,     // No-local
		false,     // No-wait
		nil,       // Arguments
	)
	if err != nil {
		return nil, fmt.Errorf("failed to register consumer: %w", err)
	}

	log.Printf("Listening for messages on queue: %s", queueName)
	return msgs, nil
}
