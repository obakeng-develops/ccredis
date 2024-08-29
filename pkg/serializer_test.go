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
			"returns hello world",
			"hello world",
			"+hello world\r\n",
		},
		{
			"returns OK",
			"OK",
			"+OK\r\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			result := pkg.SerializeSimpleStrings(tt.input)

			if result != tt.expected {
				t.Errorf("got %s, want %s", tt.input, tt.expected)
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
			"returns 10",
			10,
			":10\r\n",
		},
		{
			"returns 100",
			100,
			":100\r\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			result := pkg.SerializeIntegers(tt.input)

			if result != tt.expected {
				t.Errorf("got %d, want %s", tt.input, tt.expected)
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
			"returns nil",
			"",
			"$-1\r\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			result := pkg.SerializeBulkStrings(tt.input)

			if result != tt.expected {
				t.Errorf("got %s, want %s", tt.input, tt.expected)
			}
		})
	}
}
