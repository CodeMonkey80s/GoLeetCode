package solution2391

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	InputA []string
	InputB []int
	Output int
}{
	{
		InputA: []string{"G", "P", "GP", "GG"},
		InputB: []int{2, 4, 3},
		Output: 21,
	},
	{
		InputA: []string{"MMM", "PGM", "GP"},
		InputB: []int{3, 10},
		Output: 37,
	},
}

func Test_garbageCollection(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v %v, Output: %v\n", tc.InputA, tc.InputB, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := garbageCollectionV2(tc.InputA, tc.InputB)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_garbageCollectionV2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = garbageCollectionV2(testCases[0].InputA, testCases[0].InputB)
	}
}

func Benchmark_garbageCollectionV1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = garbageCollectionV1(testCases[0].InputA, testCases[0].InputB)
	}
}
