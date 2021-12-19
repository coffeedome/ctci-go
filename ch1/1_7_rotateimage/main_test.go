package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRotateImageBy90Degrees(t *testing.T) {
	type testCase struct {
		testName           string
		testInput          [][]int
		testExpectedOutput [][]int
		testError          error
	}

	testCases := []testCase{
		{
			testName:           "Valid Square matrix",
			testInput:          [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			testExpectedOutput: [][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}},
			testError:          nil,
		},
		{
			testName:           "Valid non-square matrix",
			testInput:          [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, {10, 11, 12}},
			testExpectedOutput: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, {10, 11, 12}},
			testError:          errors.New("matrix is not square"),
		},
	}

	for _, tt := range testCases {
		t.Run(tt.testName, func(t *testing.T) {
			out, err := RotateImageBy90Degrees(tt.testInput)
			assert.Equal(t, tt.testExpectedOutput, out)
			assert.Equal(t, tt.testError, err)
		})
	}

}
