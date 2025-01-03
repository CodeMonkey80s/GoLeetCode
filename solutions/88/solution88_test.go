package solution88

import (
	"fmt"
	"reflect"
	"testing"
)

var testCases = []struct {
	Nums1  []int
	Num1   int
	Nums2  []int
	Num2   int
	Output []int
}{
	// Mandatory Test Cases
	{
		Nums1:  []int{1, 2, 3, 0, 0, 0},
		Num1:   3,
		Nums2:  []int{2, 5, 6},
		Num2:   3,
		Output: []int{1, 2, 2, 3, 5, 6},
	},
	{
		Nums1:  []int{1},
		Num1:   1,
		Nums2:  []int{},
		Num2:   0,
		Output: []int{1},
	},
	{
		Nums1:  []int{0},
		Num1:   0,
		Nums2:  []int{1},
		Num2:   1,
		Output: []int{1},
	},
	// Additional my custom cases
}

func Test_merge(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Nums: %v, %d, %v, %d, Output: %v\n", tc.Nums1, tc.Num1, tc.Nums2, tc.Num2, tc.Output)
		t.Run(label, func(t *testing.T) {
			merge(tc.Nums1, tc.Num1, tc.Nums2, tc.Num2)
			if !reflect.DeepEqual(tc.Nums1, tc.Output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, tc.Nums1)
			}
		})
	}
}

func Benchmark_merge(b *testing.B) {
	for i := 0; i < b.N; i++ {
		merge(testCases[0].Nums1, testCases[0].Num1, testCases[0].Nums2, testCases[0].Num2)
	}
}
