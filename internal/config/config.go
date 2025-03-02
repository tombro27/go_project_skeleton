package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// Config structure for entire YAML file
type Config struct {
	Database map[string]DBConfig    `yaml:"database"`
	RabbitMQ map[string]RMQConfig   `yaml:"rabbitmq"`
	Kafka    map[string]KafkaConfig `yaml:"kafka"`
}

// Database Configuration (Multiple DBs in a Map)
type DBConfig struct {
	Username string `yaml:"username"`
	Driver   string `yaml:"driver"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	DbName   string `yaml:"dbname"`
	Password string `yaml:"-"` // Loaded separately from secrets.yml
}

// RabbitMQ Configuration (Multiple Servers in a Map)
type RMQConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	VHost    string `yaml:"vhost"`
	Username string `yaml:"username"`
	Password string `yaml:"-"` // Loaded separately
	URL      string `yaml:"url,omitempty"`
}

// Kafka Configuration
type KafkaConfig struct {
	Broker   string `yaml:"broker"`
	Group    string `yaml:"group"`
	Password string `yaml:"-"` // Loaded separately
}

// LoadConfig loads YAML based on environment
func LoadConfig() (*Config, error) {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "dev" // Default to dev
	}

	configFile := fmt.Sprintf("internal/config/%s.yml", env)

	data, err := os.ReadFile(configFile)
	if err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling YAML: %w", err)
	}

	// Load secrets (passwords)
	if err := LoadSecrets(&config); err != nil {
		return nil, err
	}

	return &config, nil
}

// LoadSecrets merges passwords into the config from secrets.yml
func LoadSecrets(cfg *Config) error {
	secretFile := "internal/config/secrets.yml"
	data, err := os.ReadFile(secretFile)
	if err != nil {
		return fmt.Errorf("error reading secrets file: %w", err)
	}

	var secrets struct {
		Database map[string]struct {
			Password string `yaml:"password"`
		} `yaml:"database"`
		RabbitMQ map[string]struct {
			Password string `yaml:"password"`
		} `yaml:"rabbitmq"`
		Kafka map[string]struct {
			Password string `yaml:"password"`
		} `yaml:"kafka"`
	}

	err = yaml.Unmarshal(data, &secrets)
	if err != nil {
		return fmt.Errorf("error unmarshaling secrets YAML: %w", err)
	}

	// Merge passwords into config
	for key, secret := range secrets.Database {
		if dbConfig, exists := cfg.Database[key]; exists {
			dbConfig.Password = secret.Password
			cfg.Database[key] = dbConfig
		}
	}
	for key, secret := range secrets.RabbitMQ {
		if rmqConfig, exists := cfg.RabbitMQ[key]; exists {
			rmqConfig.Password = secret.Password
			cfg.RabbitMQ[key] = rmqConfig
		}
	}
	for key, secret := range secrets.Kafka {
		if kafkaConfig, exists := cfg.Kafka[key]; exists {
			kafkaConfig.Password = secret.Password
			cfg.Kafka[key] = kafkaConfig
		}
	}

	return nil
}
