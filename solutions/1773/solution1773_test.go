package solution1773

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	InputA [][]string
	InputB string
	InputC string
	Output int
}{
	{
		InputA: [][]string{{"phone", "blue", "pixel"}, {"computer", "silver", "lenovo"}, {"phone", "gold", "iphone"}},
		InputB: "color",
		InputC: "silver",
		Output: 1,
	},
	{
		InputA: [][]string{{"phone", "blue", "pixel"}, {"computer", "silver", "phone"}, {"phone", "gold", "iphone"}},
		InputB: "type",
		InputC: "phone",
		Output: 2,
	},
}

func Test_mostWordsFound(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v Output: %v\n", tc.InputA, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := countMatches(tc.InputA, tc.InputB, tc.InputC)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_countMatches(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = countMatches(testCases[0].InputA, testCases[0].InputB, testCases[0].InputC)
	}
}
