package solution3280

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
		Input:  "2080-02-29",
		Output: "100000100000-10-11101",
	},
	{
		Input:  "1900-01-01",
		Output: "11101101100-1-1",
	},
	// Additional my custom cases
}

func Test_convertDateToBinary(t *testing.T) {
	for _, tc := range testCases {
		label := fmt.Sprintf("Case: Input: %v Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := convertDateToBinaryV2(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_convertDateToBinaryV2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = convertDateToBinaryV2(testCases[0].Input)
	}
}

func Benchmark_convertDateToBinaryV1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = convertDateToBinaryV1(testCases[0].Input)
	}
}
