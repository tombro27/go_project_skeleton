package main

import (
	"fmt"
	"log"

	"github.com/tombro27/go_project_skeleton/internal/config"
	"github.com/tombro27/go_project_skeleton/pkg/database"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	mainCon, err := database.GetDB("ora1", cfg.Database)
	if err != nil {
		log.Fatalf("DB Connection failed: %v", err)
	}

	// Using ExecuteNonQuery to execute insert query
	query := "INSERT INTO users (id, name, age) VALUES (:id, :name, :age)"
	params := map[string]interface{}{
		"id":   101,
		"name": "John Doe",
		"age":  30,
	}

	rowsAffected, err := database.ExecuteNonQuery(mainCon, query, params)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Rows Inserted:", rowsAffected)

	// Using ExecuteNonQuery to execute update query
	query = "UPDATE users SET name = :name WHERE id = :id"
	params = map[string]interface{}{
		"id":   101,
		"name": "John Updated",
	}

	rowsAffected, err = database.ExecuteNonQuery(mainCon, query, params)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Rows Updated:", rowsAffected)

	// Using ExecuteSelect to execute select query
	query = "SELECT id, name FROM users WHERE age > :age"
	params = map[string]interface{}{
		"age": 25,
	}

	results, err := database.ExecuteSelect(mainCon, query, params)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Results:", results)

}
