package database

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	_ "github.com/godror/godror" // Oracle driver
	_ "github.com/lib/pq"        // PostgreSQL driver
	"github.com/tombro27/go_project_skeleton/internal/config"
)

// Global DB connection map (to handle multiple DBs)
var (
	dbConnections = make(map[string]*sql.DB)
	mu            sync.Mutex
)

// GetDB returns a database connection for the given DB key
func GetDB(dbKey string, config map[string]config.DBConfig) (*sql.DB, error) {
	mu.Lock()
	defer mu.Unlock()

	// If DB connection already exists, return it
	if db, ok := dbConnections[dbKey]; ok {
		return db, nil
	}

	// Get config and password
	cfg, exists := config[dbKey]
	if !exists {
		return nil, fmt.Errorf("database key %s not found", dbKey)
	}
	password := config[dbKey].Password
	if password == "" {
		return nil, fmt.Errorf("password for database key %s not found", dbKey)
	}

	// Build DSN (Data Source Name)
	var dsn string
	switch cfg.Driver {
	case "postgres":
		dsn = fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
			cfg.Username, password, cfg.Host, cfg.Port, cfg.DbName)
	case "godror":
		dsn = fmt.Sprintf("%s/%s@%s:%d/%s",
			cfg.Username, password, cfg.Host, cfg.Port, cfg.DbName)
	default:
		return nil, fmt.Errorf("unsupported driver: %s", cfg.Driver)
	}

	// Open DB connection
	db, err := sql.Open(cfg.Driver, dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to DB: %v", err)
	}

	// Ping DB to check connection
	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping DB: %v", err)
	}

	// Store connection in map
	dbConnections[dbKey] = db
	log.Printf("Connected to %s database successfully!", dbKey)
	return db, nil
}

func ExecuteSelect(db *sql.DB, query string, params map[string]interface{}) ([]map[string]interface{}, error) {
	// Convert named parameters map to ordered slice
	args := make([]interface{}, 0, len(params))
	for key, value := range params {
		query = replaceNamedParam(query, key)
		args = append(args, value)
	}

	// Prepare statement
	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()

	// Execute query
	rows, err := stmt.Query(args...)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	// Get column names
	columns, err := rows.Columns()
	if err != nil {
		return nil, fmt.Errorf("failed to get columns: %w", err)
	}

	// Prepare result storage
	var results []map[string]interface{}

	// Iterate over rows
	for rows.Next() {
		// Create a slice to store column values
		values := make([]interface{}, len(columns))
		valuePtrs := make([]interface{}, len(columns))
		for i := range values {
			valuePtrs[i] = &values[i]
		}

		// Scan row into values
		if err := rows.Scan(valuePtrs...); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}

		// Store row data in a map
		rowMap := make(map[string]interface{})
		for i, colName := range columns {
			rowMap[colName] = values[i]
		}

		results = append(results, rowMap)
	}

	// Check for errors during iteration
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %w", err)
	}

	return results, nil
}

// ExecuteNonQuery runs INSERT or UPDATE queries with named parameters and returns the number of rows affected.
func ExecuteNonQuery(db *sql.DB, query string, params map[string]interface{}) (int64, error) {
	// Convert named parameters map to ordered slice
	args := make([]interface{}, 0, len(params))
	for key, value := range params {
		query = replaceNamedParam(query, key)
		args = append(args, value)
	}

	// Prepare statement
	stmt, err := db.Prepare(query)
	if err != nil {
		return 0, fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close()

	// Execute query
	result, err := stmt.Exec(args...)
	if err != nil {
		return 0, fmt.Errorf("failed to execute query: %w", err)
	}

	// Get affected rows
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("failed to retrieve affected rows: %w", err)
	}

	return rowsAffected, nil
}

// replaceNamedParam replaces named placeholders like :param with ?
func replaceNamedParam(query, param string) string {
	return query // Modify this based on your DB (e.g., replace `:param` with `?`)
}
