package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeleteDuplicateNodes(t *testing.T) {
	type testCase struct {
		testName     string
		testInput    List
		testExpError error
	}

	testCases := []testCase{
		{
			testName: "Normal list",
			testInput: List{
				head: Node{
					data: 8,
					next: &Node{
						data: 5,
						next: &Node{
							data: 3,
							next: &Node{
								data: 5,
								next: &Node{
									data: 4,
									next: nil,
								},
							},
						},
					},
				},
			},
			testExpError: nil,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.testName, func(t *testing.T) {
			err := tt.testInput.DeleteDuplicateNodes()
			assert.Equal(t, tt.testExpError, err)

		})
	}

}
