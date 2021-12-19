package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPermutationOfPalindrome(t *testing.T) {
	type testCase struct {
		testName           string
		testInput          string
		testExpectedOutput bool
		testError          error
	}

	test_cases := []testCase{
		{
			testName:           "Valid string",
			testInput:          "ocat tac",
			testExpectedOutput: true,
			testError:          nil,
		},
		{
			testName:           "Invalid string",
			testInput:          "foo bar",
			testExpectedOutput: false,
			testError:          nil,
		},
		{
			testName:           "Empty string",
			testInput:          "",
			testExpectedOutput: false,
			testError:          errors.New("empty string is invalid"),
		},
	}

	for _, tt := range test_cases {
		t.Run(tt.testName, func(t *testing.T) {
			out, err := IsPermutationOfPalindrome(tt.testInput)
			assert.Equal(t, tt.testExpectedOutput, out)
			assert.Equal(t, tt.testError, err)
		})
	}
}
