package solution500

import (
	"fmt"
	"reflect"
	"testing"
)

var testCases = []struct {
	Input  []string
	Output []string
}{
	// Mandatory Test Cases
	{
		Input:  []string{"Hello", "Alaska", "Dad", "Peace"},
		Output: []string{"Alaska", "Dad"},
	},
	{
		Input:  []string{"omk"},
		Output: []string{},
	},
	{
		Input:  []string{"adsdf", "sfd"},
		Output: []string{"adsdf", "sfd"},
	},
	// Additional my custom cases
	{
		Input:  []string{"Az"},
		Output: []string{},
	},
}

func Test_findWords(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := findWords(tc.Input)
			if !reflect.DeepEqual(output, tc.Output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_findWords(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = findWords(testCases[0].Input)
	}
}
