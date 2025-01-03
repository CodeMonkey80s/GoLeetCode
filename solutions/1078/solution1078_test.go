package solution1078

import (
	"fmt"
	"reflect"
	"testing"
)

var testCases = []struct {
	InputA string
	InputB string
	InputC string
	Output []string
}{
	// Mandatory Test Cases
	{
		InputA: "alice is a good girl she is a good student",
		InputB: "a",
		InputC: "good",
		Output: []string{"girl", "student"},
	},
	{
		InputA: "we will we will rock you",
		InputB: "we",
		InputC: "will",
		Output: []string{"we", "rock"},
	},
	// Additional my custom cases
	{
		InputA: "ypkk lnlqhmaohv lnlqhmaohv lnlqhmaohv ypkk ypkk ypkk ypkk ypkk ypkk lnlqhmaohv lnlqhmaohv lnlqhmaohv lnlqhmaohv ypkk ypkk ypkk lnlqhmaohv lnlqhmaohv ypkk",
		InputB: "lnlqhmaohv",
		InputC: "ypkk",
		Output: []string{"ypkk", "ypkk"},
	},
}

func Test_find(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, %v, %v Output: %v\n", tc.InputA, tc.InputB, tc.InputC, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := findOcurrences(tc.InputA, tc.InputB, tc.InputC)
			if !reflect.DeepEqual(output, tc.Output) {
				t.Errorf("Expected output to be %q but we got %q", tc.Output, output)
			}
		})
	}
}

func Benchmark_find(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = findOcurrences(testCases[0].InputA, testCases[0].InputB, testCases[0].InputC)
	}
}
