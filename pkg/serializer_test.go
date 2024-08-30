package pkg_test

import (
	"testing"

	"github.com/obakeng-develops/redis-server/pkg"
)

func TestSimpleStrings(t *testing.T) {
	tests := []struct {
		description string
		input       string
		expected    string
	}{
		{
			description: "returns hello world",
			input:       "hello world",
			expected:    "+hello world\r\n",
		},
		{
			description: "returns OK",
			input:       "OK",
			expected:    "+OK\r\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			result := pkg.SerializeSimpleStrings(tt.input)

			if result != tt.expected {
				t.Errorf("got %s, want %s", result, tt.expected)
			}
		})
	}
}

func TestErrors(t *testing.T) {
	tests := []struct {
		description string
		input       int
		expected    string
	}{
		{
			description: "returns 10",
			input:       10,
			expected:    ":10\r\n",
		},
		{
			description: "returns 100",
			input:       100,
			expected:    ":100\r\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			result := pkg.SerializeIntegers(tt.input)

			if result != tt.expected {
				t.Errorf("got %s, want %s", result, tt.expected)
			}
		})
	}
}

func TestBulkStrings(t *testing.T) {
	tests := []struct {
		description string
		input       string
		expected    string
	}{
		{
			description: "returns hello",
			input:       "hello",
			expected:    "$5\r\nhello\r\n",
		},
		{
			description: "returns",
			input:       "",
			expected:    "$0\r\n\r\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			result := pkg.SerializeBulkStrings(tt.input)

			if result != tt.expected {
				t.Errorf("got %s, want %s", result, tt.expected)
			}
		})
	}
}

func TestArrays(t *testing.T) {
	tests := []struct {
		description string
		input       interface{}
		expected    string
	}{
		{
			description: "returns empty array",
			input:       []string{},
			expected:    "*0\r\n",
		},
		{
			description: "returns 2 element array",
			input:       []string{"hello", "world"},
			expected:    "*2\r\n$5\r\nhello\r\n$5\r\nworld\r\n",
		},
		{
			description: "returns 3 integer array",
			input:       []int{1, 2, 3},
			expected:    "*3\r\n:1\r\n:2\r\n:3\r\n",
		},
		{
			description: "returns unsupported type",
			input:       map[string]int{"a": 1},
			expected:    "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			result := pkg.SerializeArrays(tt.input)

			if result != tt.expected {
				t.Errorf("got %s, want %s", result, tt.expected)
			}
		})
	}
}
