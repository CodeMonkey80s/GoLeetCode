package solution3028

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input  []int
	Output int
}{
	{
		Input:  []int{2, 3, -5},
		Output: 1,
	},
	{
		Input:  []int{3, 2, -3, -4},
		Output: 0,
	},
}

func Test_returnToBoundaryCount(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := returnToBoundaryCount(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %q but we got %q", tc.Output, output)
			}
		})
	}
}

func BenchmarkGetEncryptedString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = returnToBoundaryCount(testCases[0].Input)
	}
}
