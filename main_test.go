package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddWithBonuses(t *testing.T) {
	testCases := []struct {
		name        string
		inputString string
		expected    int
		wantError   bool
	}{
		{
			name:        "should return the summation when input is only string",
			inputString: "4,5,6",
			expected:    15,
		},
		{
			name:        "should return the summation",
			inputString: "1\n,2,3",
			expected:    6,
		},
		{
			name:        "should return the summation when input has a single custom delimiter",
			inputString: "//;\n1;3;4",
			expected:    8,
		},
		{
			name:        "should return the summation when input has an arbitrary length delimiter",
			inputString: "//***\n1***2***3",
			expected:    6,
		},
		{
			name:        "should return the summation when input has multiple delimiters",
			inputString: "//$,@\n1$2@3",
			expected:    6,
		},
		{
			name:        "should return the summation when input has multiple delimiters of arbitrary length",
			inputString: "//@abc,$eq\n1@abc3$eq4",
			expected:    8,
		},
		{
			name:        "should return error if a number is negative",
			inputString: "//***\n1***-2***3",
			wantError:   true,
		},
		{
			name:        "should skip larger than 1000",
			inputString: "//***\n1***1005***3",
			expected:    4,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := Add(tc.inputString)

			if err != nil {
				assert.Error(t, err)
			} else {
				assert.Equal(t, tc.expected, actual, tc.name)
				assert.Nil(t, err)
			}
		})
	}
}
