package utils

import (
	"errors"
	"fmt"
	"reflect"
)

// InterfaceToString converts an interface{} to a string.
func InterfaceToString(input interface{}) (string, error) {
	// Handle nil input
	if input == nil {
		return "", errors.New("input is nil")
	}

	// Type assertion for string
	if str, ok := input.(string); ok {
		return str, nil
	}

	// Handle other types using reflection
	value := reflect.ValueOf(input)
	if value.Kind() == reflect.String {
		return value.String(), nil
	}

	// Convert basic types like int, float, bool to string
	switch v := input.(type) {
	case fmt.Stringer:
		return v.String(), nil
	case error:
		return v.Error(), nil
	case int, int8, int16, int32, int64:
		return fmt.Sprintf("%d", v), nil
	case uint, uint8, uint16, uint32, uint64:
		return fmt.Sprintf("%d", v), nil
	case float32, float64:
		return fmt.Sprintf("%f", v), nil
	case bool:
		return fmt.Sprintf("%t", v), nil
	}

	return "", errors.New("unsupported type")
}
