package pkg_test

import (
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
