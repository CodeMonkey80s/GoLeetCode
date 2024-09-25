package solution3270

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	InputA int
	InputB int
	InputC int
	Output int
}{

	{
		InputA: 1,
		InputB: 10,
		InputC: 1000,
		Output: 0,
	},
	{
		InputA: 1,
		InputB: 2,
		InputC: 3,
		Output: 1,
	},
	{
		InputA: 987,
		InputB: 879,
		InputC: 798,
		Output: 777,
	},
}

func Test_generateKey(t *testing.T) {
	for _, tc := range testCases {
		label := fmt.Sprintf("Case: Input: %v %v %v Output: %v\n", tc.InputA, tc.InputB, tc.InputC, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := generateKeyV2(tc.InputA, tc.InputB, tc.InputC)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func BenchmarkGenerateKeyV2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = generateKeyV2(testCases[0].InputA, testCases[0].InputB, testCases[0].InputC)
	}
}

func BenchmarkGenerateKeyV1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = generateKeyV1(testCases[0].InputA, testCases[0].InputB, testCases[0].InputC)
	}
}
