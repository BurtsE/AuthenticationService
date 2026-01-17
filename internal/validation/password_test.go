package validation

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPassword(t *testing.T) {
	testCases := []struct {
		name     string
		password string
		expected bool
	}{
		{
			name:     "Valid password",
			password: "Password123!",
			expected: true,
		},
		{
			name:     "Password too short",
			password: "Pass1!",
			expected: false,
		},
		{
			name:     "No uppercase",
			password: "password123!",
			expected: false,
		},
		{
			name:     "No lowercase",
			password: "PASSWORD123!",
			expected: false,
		},
		{
			name:     "No digit",
			password: "Password!!",
			expected: false,
		},
		{
			name:     "No special character",
			password: "Password123",
			expected: false,
		},
		{
			name:     "Empty password",
			password: "",
			expected: false,
		},
		{
			name:     "Password with spaces",
			password: "Pass word123!",
			expected: true, // Spaces are not special characters, but they don't invalidate other rules
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := ValidPassword(tc.password)
			assert.Equal(t, tc.expected, actual)
		})

	}
}
