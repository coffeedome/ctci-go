package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompressString(t *testing.T) {
	type testCase struct {
		testName           string
		testInput          string
		testExpectedOutput string
		testError          error
	}

	testCases := []testCase{
		{
			testName:           "String with repeat characters",
			testInput:          "aaabcccdaa",
			testExpectedOutput: "a3bc3da2",
			testError:          nil,
		},
		{
			testName:           "String without repeat characters",
			testInput:          "abcd",
			testExpectedOutput: "abcd",
			testError:          nil,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.testName, func(t *testing.T) {
			out := CompressString(tt.testInput)

			assert.Equal(t, tt.testExpectedOutput, out)
		})
	}
}
