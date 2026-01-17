package validation

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEmailVerification(t *testing.T) {
	testCases := []struct {
		name     string
		email    string
		expected bool
	}{
		{
			name:     "Valid email",
			email:    "test@example.com",
			expected: true,
		},
		{
			name:     "Invalid email format",
			email:    "invalid-email",
			expected: false,
		},
		{
			name:     "Empty email",
			email:    "",
			expected: false,
		},
		{
			name:     "Email with leading/trailing spaces",
			email:    "  test@example.com  ",
			expected: true,
		},
		{
			name:     "Email with no domain",
			email:    "test@",
			expected: false,
		},
		{
			name:     "Email with no local part",
			email:    "@example.com",
			expected: false,
		},
		{
			name:     "Email with multiple @ symbols",
			email:    "test@example@com",
			expected: false,
		},
		{
			name:     "Email with special characters in local part",
			email:    "test.name+alias@example.com",
			expected: true,
		},
		{
			name:     "Email with IP address as domain",
			email:    "test@[192.168.1.1]",
			expected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := ValidEmail(tc.email)
			assert.Equal(t, tc.expected, actual)
		})q
	}
}
