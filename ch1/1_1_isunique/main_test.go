package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsUniqueWithMap(t *testing.T) {

	type testCase struct {
		caseName       string
		input          string
		expectedOutput bool
	}

	test_cases := []testCase{
		{
			caseName:       "Unique characters",
			input:          "dog",
			expectedOutput: true,
		},
		{
			caseName:       "Nonunique characters",
			input:          "loon",
			expectedOutput: false,
		},
		{
			caseName:       "Empty string",
			input:          "",
			expectedOutput: true,
		},
	}

	for _, tt := range test_cases {
		t.Run(tt.caseName, func(t *testing.T) {
			out := IsUniqueWithMap(tt.input)
			assert.Equal(t, tt.expectedOutput, out)
		})
	}
}
