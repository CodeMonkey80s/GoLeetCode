package solution1854

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input  [][]int
	Output int
}{
	{
		Input:  [][]int{{1993, 1999}, {2000, 2010}},
		Output: 1993,
	},
	{
		Input:  [][]int{{1950, 1961}, {1960, 1971}, {1970, 1981}},
		Output: 1960,
	},
	{
		Input:  [][]int{{1987, 2047}, {1952, 2006}, {2021, 2042}, {2047, 2049}, {2036, 2040}, {1994, 2009}},
		Output: 1994,
	},
}

func Test_maximumPopulations(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := maximumPopulations(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_maximumPopulation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = maximumPopulations(testCases[0].Input)
	}
}
