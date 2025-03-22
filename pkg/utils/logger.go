package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime/debug"
	"time"
)

func init() {
	_ = os.MkdirAll("logs", os.ModePerm)
}

func getLogFilePath(logType string) string {
	date := time.Now().Format("2006-01-02")
	return filepath.Join("logs", fmt.Sprintf("%s_%s.log", logType, date))
}

// App Log Struct
type AppLog struct {
	WorkerName string                 `json:"worker_name"`
	WorkerType string                 `json:"worker_type"`
	StartTime  time.Time              `json:"start_time"`
	EndTime    time.Time              `json:"end_time"`
	Data       map[string]interface{} `json:"data"`
}

// Error Log Struct
type ErrorLog struct {
	WorkerName string    `json:"worker_name"`
	WorkerType string    `json:"worker_type"`
	ErrorTime  time.Time `json:"error_time"`
	ErrorMsg   string    `json:"error_msg"`
	StackTrace string    `json:"stack_trace"`
}

// App Log Function
func LogApplication(log AppLog) error {
	return writeJSONToFile("application", log)
}

// Error Log Function
func LogError(workerName, workerType, errMsg string) error {
	errLog := ErrorLog{
		WorkerName: workerName,
		WorkerType: workerType,
		ErrorTime:  time.Now(),
		ErrorMsg:   errMsg,
		StackTrace: string(debug.Stack()),
	}
	return writeJSONToFile("error", errLog)
}

// Write JSON to File
func writeJSONToFile(logType string, data interface{}) error {
	filePath := getLogFilePath(logType)
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return fmt.Errorf("could not open log file: %v", err)
	}
	defer file.Close()

	jsonData, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("could not marshal log to JSON: %v", err)
	}

	_, err = file.WriteString(string(jsonData) + "\n")
	return err
}
