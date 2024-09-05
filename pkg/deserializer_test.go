package pkg_test

import (
	"reflect"
	"testing"

	"github.com/obakeng-develops/redis-server/pkg"
)

func TestDeserializeSimplesStrings(t *testing.T) {
	tests := []struct {
		description string
		input       string
		expected    string
	}{
		{
			description: "returns OK",
			input:       "+OK\r\n",
			expected:    "OK",
		},
		{
			description: "returns PONG",
			input:       "+PONG\r\n",
			expected:    "PONG",
		},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			result := pkg.DeserializeSimpleStrings(tc.input)

			if result != tc.expected {
				t.Errorf("got %s, want %s", result, tc.expected)
			}
		})
	}
}

func TestDeserializeErrors(t *testing.T) {
	tests := []struct {
		description string
		input       string
		expected    string
	}{
		{
			description: "returns unknown command",
			input:       "-unknown command\r\n",
			expected:    "unknown command",
		},
		{
			description: "returns operation against wrong type",
			input:       "-operation against wrong type\r\n",
			expected:    "operation against wrong type",
		},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			result := pkg.DeserializeErrors(tc.input)

			if result != tc.expected {
				t.Errorf("got %s, want %s", result, tc.expected)
			}
		})
	}
}

func TestDeserializeIntegers(t *testing.T) {
	tests := []struct {
		description string
		input       string
		expected    int
	}{
		{
			description: "returns 100",
			input:       ":100\r\n",
			expected:    100,
		},
		{
			description: "returns 10",
			input:       ":10\r\n",
			expected:    10,
		},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			result := pkg.DeserializeIntegers(tc.input)

			if result != tc.expected {
				t.Errorf("got %v, want %v", result, tc.expected)
			}
		})
	}
}

func TestDeserializeBulkStrings(t *testing.T) {
	tests := []struct {
		description string
		input       string
		expected    string
	}{
		{
			description: "returns hello",
			input:       "$5\r\nhello\r\n",
			expected:    "hello",
		},
		{
			description: "returns 10",
			input:       "$0\r\n\r\n",
			expected:    "",
		},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			result := pkg.DeserializeBulkStrings(tc.input)

			if result != tc.expected {
				t.Errorf("got %v, want %v", result, tc.expected)
			}
		})
	}
}

func TestDeserializeArrays(t *testing.T) {
	tests := []struct {
		description string
		input       string
		expected    interface{}
	}{
		{
			description: "returns 3 integer arrary",
			input:       "*3\r\n:1\r\n:2\r\n:3\r\n",
			expected:    []int{1, 2, 3},
		},
		{
			description: "returns hello world",
			input:       "*2\r\n$5\r\nhello\r\n$5\r\nworld\r\n",
			expected:    []string{"hello", "world"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			result := pkg.DeserializeArrays(tc.input)

			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("got %v, want %v", result, tc.expected)
			}
		})
	}
}
