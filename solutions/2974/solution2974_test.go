package solution2974

import (
	"fmt"
	"reflect"
	"testing"
)

var testCases = []struct {
	Input  []int
	Output []int
}{
	{
		Input:  []int{5, 4, 2, 3},
		Output: []int{3, 2, 5, 4},
	},
	{
		Input:  []int{2, 5},
		Output: []int{5, 2},
	},
}

func Test_numberGame(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := numberGame(tc.Input)
			if !reflect.DeepEqual(output, tc.Output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func BenchmarkNumberGame(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = numberGame(testCases[0].Input)
	}
}
