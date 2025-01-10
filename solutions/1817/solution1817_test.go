package solution1817

import (
	"fmt"
	"reflect"
	"testing"
)

var testCases = []struct {
	InputA [][]int
	InputB int
	Output []int
}{
	{
		InputA: [][]int{{0, 5}, {1, 2}, {0, 2}, {0, 5}, {1, 3}},
		InputB: 5,
		Output: []int{0, 2, 0, 0, 0},
	},
	{
		InputA: [][]int{{1, 1}, {2, 2}, {2, 3}},
		InputB: 4,
		Output: []int{1, 1, 0, 0},
	},
	{
		InputA: [][]int{{283268890, 14532}, {283268891, 14530}, {283268889, 14530}, {283268892, 14530}, {283268890, 14531}},
		InputB: 2,
		Output: []int{3, 1},
	},
}

func Test_findingUsersActiveMinutes(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v %v, Output: %v\n", tc.InputA, tc.InputB, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := findingUsersActiveMinutesV1(tc.InputA, tc.InputB)
			if !reflect.DeepEqual(output, tc.Output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
