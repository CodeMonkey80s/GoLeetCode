package solution1662

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	InputA []string
	InputB []string
	Output bool
}{
	// Mandatory Test Cases
	{
		InputA: []string{"ab", "c"},
		InputB: []string{"a", "bc"},
		Output: true,
	},
	{
		InputA: []string{"a", "cb"},
		InputB: []string{"ab", "c"},
		Output: false,
	},
	{
		InputA: []string{"abc", "d", "defg"},
		InputB: []string{"abcddefg"},
		Output: true,
	},
	// Additional my custom cases
}

func Test_arrayStringsAreEqual(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v %v Output: %v\n", tc.InputA, tc.InputB, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := arrayStringsAreEqual(tc.InputA, tc.InputB)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_arrayStringsAreEqual(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = arrayStringsAreEqual(testCases[0].InputA, testCases[0].InputB)
	}
}
