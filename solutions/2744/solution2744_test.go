package solution2744

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input  []string
	Output int
}{
	{
		Input:  []string{"cd", "ac", "dc", "ca", "zz"},
		Output: 2,
	},
	{
		Input:  []string{"ab", "ba", "cc"},
		Output: 1,
	},
	{
		Input:  []string{"aa", "ab"},
		Output: 0,
	},
	{
		Input:  []string{"az", "zz"},
		Output: 0,
	},
	{
		Input:  []string{"xi", "nw", "qp", "to", "oo", "xp", "ix", "wn", "pq"},
		Output: 3,
	},
	{
		Input:  []string{"ff", "by", "dy", "rk", "nm", "yb"},
		Output: 1,
	},
}

func Test_maximumNumberOfStringPairs(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := maximumNumberOfStringPairs(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_maximumNumberOfStringPairs_first_attempt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = maximumNumberOfStringPairs_first_attempt(testCases[0].Input)
	}

}

func Benchmark_maximumNumberOfStringPairs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = maximumNumberOfStringPairs(testCases[0].Input)
	}
}
