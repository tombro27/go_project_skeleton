package testex

import (
	"testing"
)

func TestAddition(t *testing.T) {
	result := Add(2, 3)
	expected := 5
	if result != expected {
		t.Errorf("Addition incorrect, got: %d, want: %d.", result, expected)
	}
}

func Add(a, b int) int {
	return a + b
}

func BenchmarkAdd(b *testing.B) {
	// Run the Add function b.N times
	for i := 0; i < b.N; i++ {
		Add(i, i*i)
	}
}
