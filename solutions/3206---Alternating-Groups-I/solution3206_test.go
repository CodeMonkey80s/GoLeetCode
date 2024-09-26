package solution3206

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input  []int
	Output int
}{
	{
		Input:  []int{0, 1, 0, 0, 1},
		Output: 3,
	},
	{
		Input:  []int{1, 1, 1},
		Output: 0,
	},
}

func Test_numberOfAlternatingGroups(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := numberOfAlternatingGroups(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func BenchmarkNumberOfAlternatingGroups(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = numberOfAlternatingGroups(testCases[0].Input)
	}
}
