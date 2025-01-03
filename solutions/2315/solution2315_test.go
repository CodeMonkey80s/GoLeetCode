package solution2315

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input  string
	Output int
}{
	{
		Input:  "l|*e*et|c**o|*de|",
		Output: 2,
	},
	{
		Input:  "iamprogrammer",
		Output: 0,
	},
	{
		Input:  "yo|uar|e**|b|e***au|tifu|l",
		Output: 5,
	},
}

func Test_countAsterisks(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := countAsterisks(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %q but we got %q", tc.Output, output)
			}
		})
	}
}
func Benchmark_countAsterisks(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = countAsterisks(testCases[0].Input)
	}
}

func Benchmark_countAsterisks_std(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = countAsterisks_std(testCases[0].Input)
	}
}
