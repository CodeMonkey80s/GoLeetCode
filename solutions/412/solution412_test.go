package solution412

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	Input  int
	Output []string
}{
	// Mandatory Test Cases
	{
		Input:  3,
		Output: []string{"1", "2", "Fizz"},
	},
	{
		Input:  5,
		Output: []string{"1", "2", "Fizz", "4", "Buzz"},
	},
	{
		Input:  15,
		Output: []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"},
	},
	// Additional my custom cases
}

func Test_fizzBuzz(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := fizzBuzz(tc.Input)
			if len(output) != len(tc.Output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_fizzBuzz(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = fizzBuzz(testCases[0].Input)
	}
}
