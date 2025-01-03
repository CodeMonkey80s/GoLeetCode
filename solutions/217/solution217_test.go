package solution217

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Nums   []int
	Output bool
}{
	// Mandatory Test Cases
	{
		Nums:   []int{1, 2, 3, 1},
		Output: true,
	},
	{
		Nums:   []int{1, 2, 3, 4},
		Output: false,
	},
	{
		Nums:   []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2},
		Output: true,
	},
	// Additional my custom cases

}

func Test_isPalindrome(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Nums: %v, Output: %v\n", tc.Nums, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := containsDuplicate(tc.Nums)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_containsDuplicate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = containsDuplicate(testCases[0].Nums)
	}
}
