package solution2079

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	InputA []int
	InputB int
	Output int
}{

	{
		InputA: []int{2, 2, 3, 3},
		InputB: 5,
		Output: 14,
	},
	{
		InputA: []int{1, 1, 1, 4, 2, 3},
		InputB: 4,
		Output: 30,
	},
	{
		InputA: []int{7, 7, 7, 7, 7, 7, 7},
		InputB: 8,
		Output: 49,
	},
}

func Test_wateringPlants(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v %v, Output: %v\n", tc.InputA, tc.InputB, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := wateringPlantsV2(tc.InputA, tc.InputB)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_wateringPlantsV1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = wateringPlantsV1(testCases[0].InputA, testCases[0].InputB)
	}
}

func Benchmark_wateringPlantsV2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = wateringPlantsV2(testCases[0].InputA, testCases[0].InputB)
	}
}
