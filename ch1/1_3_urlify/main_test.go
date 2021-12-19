package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUrlify(t *testing.T) {
	type testCase struct {
		testName           string
		testInput          string
		testExpectedOutput string
		testExpectedError  error
	}

	test_cases := []testCase{
		{
			testName:           "String with spaces",
			testInput:          " Foo  Bar ",
			testExpectedOutput: "%20Foo%20%20Bar%20",
			testExpectedError:  nil,
		},
		{
			testName:           "String without spaces",
			testInput:          "FooBar",
			testExpectedOutput: "FooBar",
			testExpectedError:  nil,
		},
		{
			testName:           "Empty String",
			testInput:          "",
			testExpectedOutput: "",
			testExpectedError:  nil,
		},
	}

	for _, tt := range test_cases {
		t.Run(tt.testName, func(t *testing.T) {
			assert.Equal(t, Urlify(tt.testInput), tt.testExpectedOutput)
		})
	}
}
