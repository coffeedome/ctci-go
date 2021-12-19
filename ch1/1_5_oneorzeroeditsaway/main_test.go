package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOneOrZeroEditsAway(t *testing.T) {
	type testCase struct {
		testName           string
		testInput          []string
		testExpectedOutput bool
		testError          error
	}

	test_cases := []testCase{
		{
			testName:           "One edit away",
			testInput:          []string{"abc", "abcd"},
			testExpectedOutput: true,
			testError:          nil,
		},
		{
			testName:           "Two edits away",
			testInput:          []string{"abc", "abcde"},
			testExpectedOutput: false,
			testError:          nil,
		},
		{
			testName:           "Zero edits away",
			testInput:          []string{"abc", "abc"},
			testExpectedOutput: true,
			testError:          nil,
		},
		{
			testName:           "Invalid empty strings",
			testInput:          []string{"", ""},
			testExpectedOutput: false,
			testError:          errors.New("invalid input"),
		},
	}

	for _, tt := range test_cases {
		t.Run(tt.testName, func(t *testing.T) {
			out, err := OneOrZeroEditsAway(tt.testInput[0], tt.testInput[1])
			assert.Equal(t, tt.testExpectedOutput, out)
			assert.Equal(t, tt.testError, err)
		})
	}
}
