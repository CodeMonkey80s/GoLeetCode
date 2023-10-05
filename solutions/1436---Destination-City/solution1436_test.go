package solution1436

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input  [][]string
	Output string
}{
	{
		Input:  [][]string{{"London", "New York"}, {"New York", "Lima"}, {"Lima", "Sao Paulo"}},
		Output: "Sao Paulo",
	},
	{
		Input:  [][]string{{"B", "C"}, {"D", "B"}, {"C", "A"}},
		Output: "A",
	},
	{
		Input:  [][]string{{"A", "Z"}},
		Output: "Z",
	},
}

func Test_destCity(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := destCity(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_destCity(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = destCity(testCases[0].Input)
	}
}
