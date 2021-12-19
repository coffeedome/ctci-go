package main

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsSubstring(t *testing.T) {
	type testCase struct {
		testName           string
		testInputString1   string
		testInputString2   string
		testExpectedOutput bool
		testError          error
	}

	testCases := []testCase{
		{
			testName:           "Match case",
			testInputString1:   "hello",
			testInputString2:   "olleh",
			testExpectedOutput: true,
			testError:          nil,
		},
		{
			testName:           "NonMatch case",
			testInputString1:   "foo",
			testInputString2:   "bar",
			testExpectedOutput: false,
			testError:          nil,
		},
		{
			testName:           "Empty string error",
			testInputString1:   "",
			testInputString2:   "",
			testExpectedOutput: false,
			testError:          errors.New("Input string empty is invalid"),
		},
	}

	for _, tt := range testCases {
		t.Run(tt.testName, func(t *testing.T) {
			s1 := tt.testInputString1
			s2 := tt.testInputString2
			out, err := IsSubstring(s1, s2)
			if err != nil {
				fmt.Print(err)
			}
			assert.Equal(t, tt.testExpectedOutput, out)
		})
	}

}
