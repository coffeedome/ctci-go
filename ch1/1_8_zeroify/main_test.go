package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestZeroify(t *testing.T) {
	type testCase struct {
		testName           string
		testInput          [][]int
		testExpectedOutput [][]int
		testError          error
	}

	testCases := []testCase{
		{
			testName: "Valid matrix",
			testInput: [][]int{
				{1, 0, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			testExpectedOutput: [][]int{
				{0, 0, 0},
				{4, 0, 6},
				{7, 0, 9},
			},
			testError: nil,
		},
		{
			testName: "No zeros",
			testInput: [][]int{
				{0, 0, 0},
				{0, 0, 0},
				{0, 0, 0},
			},
			testExpectedOutput: [][]int{
				{0, 0, 0},
				{0, 0, 0},
				{0, 0, 0},
			},
			testError: nil,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.testName, func(t *testing.T) {
			out := Zeroify(tt.testInput)
			assert.Equal(t, tt.testExpectedOutput, out)
		})
	}
}
