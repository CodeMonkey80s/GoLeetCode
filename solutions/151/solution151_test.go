package solution151

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input  string
	Output string
}{

	{
		Input:  "the sky is blue",
		Output: "blue is sky the",
	},
	{
		Input:  "  hello world  ",
		Output: "world hello",
	},
	{
		Input:  "a good   example",
		Output: "example good a",
	},
}

func Test_reverseWords(t *testing.T) {
	for _, tc := range testCases {
		label := fmt.Sprintf("Case: Input: %v Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := reverseWordsV2(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func BenchmarkReverseWordsV2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = reverseWordsV2(testCases[0].Input)
	}
}

func BenchmarkReverseWordsV1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = reverseWordsV1(testCases[0].Input)
	}
}
