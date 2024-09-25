package solution219

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	InputA []int
	InputB int
	Output bool
}{
	{
		InputA: []int{1, 2, 3, 1},
		InputB: 3,
		Output: true,
	},
	{
		InputA: []int{1, 0, 1, 1},
		InputB: 1,
		Output: true,
	},
	{
		InputA: []int{1, 2, 3, 1, 2, 3},
		InputB: 2,
		Output: false,
	},
	{
		InputA: []int{1, 2, 3, 4, 5, 6, 7, 9, 9},
		InputB: 3,
		Output: true,
	},
}

func Test_containsNearbyDuplicate(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v %v Output: %v\n", tc.InputA, tc.InputB, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := containsNearbyDuplicate(tc.InputA, tc.InputB)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_containsNearbyDuplicate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = containsNearbyDuplicate(testCases[0].InputA, testCases[0].InputB)
	}
}
