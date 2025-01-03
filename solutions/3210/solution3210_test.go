package solution3210

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	InputA string
	InputB int
	Output string
}{
	{
		InputA: "dart",
		InputB: 3,
		Output: "tdar",
	},
	{
		InputA: "aaa",
		InputB: 1,
		Output: "aaa",
	},
}

func Test_getEncryptedString(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v %v Output: %v\n", tc.InputA, tc.InputB, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := getEncryptedString(tc.InputA, tc.InputB)
			if output != tc.Output {
				t.Errorf("Expected output to be %q but we got %q", tc.Output, output)
			}
		})
	}
}

func BenchmarkGetEncryptedString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = getEncryptedString(testCases[0].InputA, testCases[0].InputB)
	}
}
