package main

import (
	"fmt"
	"time"

	"github.com/tombro27/go_project_skeleton/internal/config"
	"github.com/tombro27/go_project_skeleton/pkg/database"
	"github.com/tombro27/go_project_skeleton/pkg/utils"
)

var workerName = "DBConCheck"

func main() {
	start := time.Now()
	logData := make(map[string]interface{})
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
		utils.LogApplication(utils.AppLog{
			WorkerName: "UserConsumer",
			WorkerType: "Kafka",
			StartTime:  start,
			EndTime:    time.Now(),
			Data:       logData,
		})
	}()

	cfg, err := config.LoadConfig()
	if err != nil {
		utils.LogError(workerName, "testApp", err.Error())
		panic("Failed to load config: " + err.Error())
	}

	mainCon, err := database.GetDB("ora1", cfg.Database)
	if err != nil {
		utils.LogError(workerName, "testApp", err.Error())
		panic("DB Connection failed: " + err.Error())
	}
	defer mainCon.Close()

	// Using ExecuteNonQuery to execute insert query
	query := "INSERT INTO users (id, name, age) VALUES (:id, :name, :age)"
	params := map[string]interface{}{
		"id":   101,
		"name": "John Doe",
		"age":  30,
	}
	rowsAffected, err := database.ExecuteNonQuery(mainCon, query, params)
	if err != nil {
		utils.LogError(workerName, "testApp", err.Error())
		panic("Error in insertion: " + err.Error())
	}
	logData["RowsInserted"] = rowsAffected
	fmt.Println("Rows Inserted:", rowsAffected)

	// Using ExecuteNonQuery to execute update query
	query = "UPDATE users SET name = :name WHERE id = :id"
	params = map[string]interface{}{
		"id":   101,
		"name": "John Updated",
	}

	rowsAffected, err = database.ExecuteNonQuery(mainCon, query, params)
	if err != nil {
		utils.LogError(workerName, "testApp", err.Error())
		panic("Error in update: " + err.Error())
	}
	fmt.Println("Rows Updated:", rowsAffected)
	logData["RowsUpdated"] = rowsAffected

	// Using ExecuteSelect to execute select query
	query = "SELECT id, name FROM users WHERE age > :age"
	params = map[string]interface{}{
		"age": 25,
	}

	results, err := database.ExecuteSelect(mainCon, query, params)
	if err != nil {
		utils.LogError(workerName, "testApp", err.Error())
		panic("Error in fetching data from DB: " + err.Error())
	}

	fmt.Println("Results:", results)
	logData["RowsFtched"] = len(results)

}
