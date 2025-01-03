package solution1

import (
	"fmt"
	"reflect"
	"testing"
)

var testCases = []struct {
	Nums   []int
	Output []int
	Target int
}{
	// Mandatory Test Cases
	{
		Nums:   []int{2, 7, 11, 15},
		Output: []int{0, 1},
		Target: 9,
	},
	{
		Nums:   []int{3, 2, 4},
		Output: []int{1, 2},
		Target: 6,
	},
	{
		Nums:   []int{3, 3},
		Output: []int{0, 1},
		Target: 6,
	},
	// Additional my custom cases
	{
		Nums:   []int{2, 5, 5, 11},
		Output: []int{1, 2},
		Target: 10,
	},
	{
		Nums:   []int{3, 2, 3},
		Output: []int{0, 2},
		Target: 6,
	},
}

func Test_twoSum_Iterations(t *testing.T) {
	for _, tc := range testCases {
		label := fmt.Sprintf("Case: Target: %d, Nums: %v, Out: %v\n", tc.Target, tc.Nums, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := twoSum_Iterations(tc.Nums, tc.Target)
			if !reflect.DeepEqual(output, tc.Output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Test_twoSum_Map(t *testing.T) {
	for _, tc := range testCases {
		label := fmt.Sprintf("Case: Target: %d, Nums: %v, Out: %v\n", tc.Target, tc.Nums, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := twoSum_Map(tc.Nums, tc.Target)
			if !reflect.DeepEqual(output, tc.Output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_twoSum_Iterations(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = twoSum_Iterations(testCases[0].Nums, testCases[0].Target)
	}
}

func Benchmark_twoSum_Map(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = twoSum_Map(testCases[0].Nums, testCases[0].Target)
	}
}
