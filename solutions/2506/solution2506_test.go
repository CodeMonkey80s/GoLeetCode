package solution2506

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input  []string
	Output int
}{
	{
		Input:  []string{"aba", "aabb", "abcd", "bac", "aabc"},
		Output: 2,
	},
	{
		Input:  []string{"aabb", "ab", "ba"},
		Output: 3,
	},
	{
		Input:  []string{"nba", "cba", "dba"},
		Output: 0,
	},
	{
		Input:  []string{"dcedceadceceaeddcedc", "dddcebcedcdbaeaaaeab", "eecbeddbddeadcbbbdbb", "decbcbebbddceacdeadd", "ccbddbaedcadedbcaaae", "dddcaadaceaedcdceedd", "bbeddbcbbccddcaceeea", "bdabacbbdadabbbddaea"},
		Output: 16,
	},
}

func Test_similarPairs(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := similarPairs(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_similarPairs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = similarPairs(testCases[0].Input)
	}
}
