package solution824

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input  string
	Output string
}{
	// Mandatory Test Cases
	{
		Input:  "I speak Goat Latin",
		Output: "Imaa peaksmaaa oatGmaaaa atinLmaaaaa",
	},
	{
		Input:  "The quick brown fox jumped over the lazy dog",
		Output: "heTmaa uickqmaaa rownbmaaaa oxfmaaaaa umpedjmaaaaaa overmaaaaaaa hetmaaaaaaaa azylmaaaaaaaaa ogdmaaaaaaaaaa",
	},
	// Additional my custom cases
}

func Test_toGoatLatin(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := toGoatLatin(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_toGoatLatin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = toGoatLatin(testCases[0].Input)
	}
}

func Benchmark_toGoatLatin_map(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = toGoatLatin_map(testCases[0].Input)
	}
}
