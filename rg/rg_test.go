// rg_test.go

package rg

import (
	"testing"
)

func TestGenerateRandomString(t *testing.T) {
	// Test cases with different options
	testCases := []struct {
		options Options
	}{
		{Options{Length: 8, Alpha: true, Numeric: true, Special: false, UseUpper: true, UseLower: true}},
		{Options{Length: 12, Alpha: false, Numeric: true, Special: true, UseUpper: false, UseLower: true}},
		// Add more test cases as needed
	}

	for _, tc := range testCases {
		t.Run("GenerateRandomString", func(t *testing.T) {
			result, err := GenerateRandomString(tc.options)

			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}

			if len(result) != tc.options.Length {
				t.Errorf("Expected string length of %d, but got %d", tc.options.Length, len(result))
			}

			// generate random string with the same options and check if the result is different from the previous one
			result2, err := GenerateRandomString(tc.options)
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
			if result == result2 {
				t.Errorf("Expected different random strings, but got the same: %s", result)
			}
			// generate 1000000 random strings and check if they are different
			for i := 0; i < 1000000; i++ {
				result3, err := GenerateRandomString(tc.options)
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
				if result == result3 {
					t.Errorf("Expected different random strings, but got the same: %s", result)
				}
			}

			// We can add more tests like check if the string contains only alpha, numeric or special characters
		})
	}
}

// Run the tests with "go test" command
