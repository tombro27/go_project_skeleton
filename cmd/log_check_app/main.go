package main

import (
	"time"

	"github.com/tombro27/go_project_skeleton/pkg/utils"
)

func main() {
	// Application log
	start := time.Now()

	utils.LogApplication(utils.AppLog{
		WorkerName: "LogTest",
		WorkerType: "TestAPP",
		StartTime:  start,
		EndTime:    time.Now(),
		Data: map[string]interface{}{
			"records_processed": 123,
			"status":            "success",
		},
	})

	// Error log
	utils.LogError("LogTest", "TestAPP", "failed to connect to broker")

}
