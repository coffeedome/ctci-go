package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPermutation(t *testing.T) {

	type testCase struct {
		testName       string
		firstInput     string
		secondInput    string
		expectedOutput bool
		expectedErr    error
	}

	test_cases := []testCase{
		{
			testName:       "Valid permutation",
			firstInput:     "abc",
			secondInput:    "bac",
			expectedOutput: true,
			expectedErr:    nil,
		},
		{
			testName:       "Invalid permutation",
			firstInput:     "abc",
			secondInput:    "def",
			expectedOutput: false,
			expectedErr:    nil,
		},
		{
			testName:       "Empty strings",
			firstInput:     "",
			secondInput:    "",
			expectedOutput: false,
			expectedErr:    errors.New("empty string"),
		},
	}

	for _, tt := range test_cases {
		t.Run(tt.testName, func(t *testing.T) {
			out, err := IsPermutation(tt.firstInput, tt.secondInput)
			assert.Equal(t, tt.expectedOutput, out)
			assert.Equal(t, tt.expectedErr, err)
		})
	}
}
