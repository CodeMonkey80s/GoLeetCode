package solution387

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input  string
	Output int
}{
	// Mandatory Test Cases
	{
		Input:  "leetcode",
		Output: 0,
	},
	{
		Input:  "loveleetcode",
		Output: 2,
	},
	{
		Input:  "aabb",
		Output: -1,
	},
	// Additional my custom cases
}

func Test_firstUniqChar(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := firstUniqChar(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_firstUniqChars(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = firstUniqChar(testCases[0].Input)
	}
}

func Benchmark_firstUniqChars_first_approach(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = firstUniqChar_first_approach(testCases[0].Input)
	}
}
