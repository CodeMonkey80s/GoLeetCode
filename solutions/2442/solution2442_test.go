package solution2442

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input  []int
	Output int
}{
	{
		Input:  []int{1, 13, 10, 12, 31},
		Output: 6,
	},
	{
		Input:  []int{2, 2, 2},
		Output: 1,
	},
	{
		Input:  []int{394583, 326065, 700337, 345233, 477213, 520531, 323169, 202303, 520493, 977282, 504271, 597209},
		Output: 24,
	},
}

func Test_countDistinctIntegers(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := countDistinctIntegers(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
