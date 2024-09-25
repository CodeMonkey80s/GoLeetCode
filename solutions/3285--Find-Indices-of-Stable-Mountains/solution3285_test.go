package solution3285

import (
	"fmt"
	"reflect"
	"testing"
)

var testCases = []struct {
	Input     []int
	Threshold int
	Output    []int
}{

	// Mandatory Test Cases
	{
		Input:     []int{1, 2, 3, 4, 5},
		Threshold: 2,
		Output:    []int{3, 4},
	},
	{

		Input:     []int{10, 1, 10, 1, 10},
		Threshold: 3,
		Output:    []int{1, 3},
	},
	{

		Input:     []int{10, 1, 10, 1, 10},
		Threshold: 10,
		Output:    nil,
	}, // Additional my custom cases
}

func Test_stableMountains(t *testing.T) {
	for _, tc := range testCases {
		label := fmt.Sprintf("Case: Input: %v Output: %v, Threshold: %d\n", tc.Input, tc.Output, tc.Threshold)
		t.Run(label, func(t *testing.T) {
			output := stableMountains(tc.Input, tc.Threshold)
			if !reflect.DeepEqual(output, tc.Output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_StableMountains(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = stableMountains(testCases[0].Input, testCases[0].Threshold)
	}
}
