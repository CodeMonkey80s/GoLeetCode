package solution821

import (
	"fmt"
	"reflect"
	"testing"
)

var testCases = []struct {
	InputA string
	InputB byte
	Output []int
}{
	// Mandatory Test Cases
	{
		InputA: "loveleetcode",
		InputB: 'e',
		Output: []int{3, 2, 1, 0, 1, 0, 0, 1, 2, 2, 1, 0},
	},
	{
		InputA: "aaab",
		InputB: 'b',
		Output: []int{3, 2, 1, 0},
	},
	// Additional my custom cases
	{
		InputA: "baaa",
		InputB: 'b',
		Output: []int{0, 1, 2, 3},
	},
}

func Test_shortestToChar(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, %c Output: %v\n", tc.InputA, tc.InputB, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := shortestToChar(tc.InputA, tc.InputB)
			if !reflect.DeepEqual(output, tc.Output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_shortestToChar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = shortestToChar(testCases[0].InputA, testCases[0].InputB)
	}
}
