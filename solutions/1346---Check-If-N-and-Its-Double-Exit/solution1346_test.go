package solution1346

import (
	"fmt"
	"reflect"
	"testing"
)

var testCases = []struct {
	Input  []int
	Output bool
}{
	{
		Input:  []int{10, 2, 5, 3},
		Output: true,
	},
	{
		Input:  []int{3, 1, 7, 11},
		Output: false,
	},
}

func Test_checkIfExist(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := checkIfExist(tc.Input)
			if !reflect.DeepEqual(output, tc.Output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func BenchmarkCheckIfExist(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = checkIfExist(testCases[0].Input)
	}
}
