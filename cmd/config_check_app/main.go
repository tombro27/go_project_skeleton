package main

import (
	"fmt"
	"log"

	"github.com/tombro27/go_project_skeleton/internal/config"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Example: Accessing Oracle Database 1 credentials
	fmt.Println("ORA1 Host:", cfg.Database["ora1"].Host)
	fmt.Println("ORA1 Username:", cfg.Database["ora1"].Username)
	fmt.Println("ORA1 Password:", cfg.Database["ora1"].Password)

	// Example: Accessing RabbitMQ Server 1 credentials
	fmt.Println("RabbitMQ Server 1 Host:", cfg.RabbitMQ["ser1"].Host)
	fmt.Println("RabbitMQ Server 1 Password:", cfg.RabbitMQ["ser1"].Password)
	fmt.Println("RabbitMQ Server 1 URL:", cfg.RabbitMQ["ser1api"].URL)

	// Example: Accessing Kafka
	fmt.Println("Kafka Broker:", cfg.Kafka["serv1"].Broker)
	fmt.Println("Kafka Password:", cfg.Kafka["serv1"].Password)
}
