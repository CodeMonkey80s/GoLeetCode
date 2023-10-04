package solution28

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Heystack string
	Needle   string
	Output   int
}{
	// Mandatory Test Cases
	{
		Heystack: "sadbutsad",
		Needle:   "sad",
		Output:   0,
	},
	{
		Heystack: "leetcode",
		Needle:   "leeto",
		Output:   -1,
	},
	// Additional my custom cases
	{
		Heystack: "bbaa",
		Needle:   "aab",
		Output:   -1,
	},
	{
		Heystack: "aaa",
		Needle:   "a",
		Output:   0,
	},
	{
		Heystack: "mississippi",
		Needle:   "issip",
		Output:   4,
	},
	{
		Heystack: "aabaaabaaac",
		Needle:   "aabaaac",
		Output:   4,
	},
	{
		Heystack: "mississippi",
		Needle:   "sippj",
		Output:   -1,
	},
	{
		Heystack: "monastero",
		Needle:   "aste",
		Output:   3,
	},
}

func Test_strStr(t *testing.T) {
	for _, tc := range testCases {
		label := fmt.Sprintf("Case: Heystack: %v, Needle: %v, Output: %v\n", tc.Heystack, tc.Needle, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := strStr(tc.Heystack, tc.Needle)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_strStr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = strStr(testCases[0].Heystack, testCases[0].Needle)
	}
}

func Benchmark_strStr_stdlib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = strStr_stdlib(testCases[0].Heystack, testCases[0].Needle)
	}
}
