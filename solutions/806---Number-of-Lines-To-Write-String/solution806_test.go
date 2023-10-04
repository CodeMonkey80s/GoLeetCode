package solution806

import (
	"fmt"
	"reflect"
	"testing"
)

var testCases = []struct {
	InputA []int
	InputB string
	Output []int
}{
	// Mandatory Test Cases
	{
		InputA: []int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10},
		InputB: "abcdefghijklmnopqrstuvwxyz",
		Output: []int{3, 60},
	},
	{
		InputA: []int{4, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10},
		InputB: "bbbcccdddaaa",
		Output: []int{2, 4},
	},
	// Additional my custom cases
	{
		InputA: []int{3, 4, 10, 4, 8, 7, 3, 3, 4, 9, 8, 2, 9, 6, 2, 8, 4, 9, 9, 10, 2, 4, 9, 10, 8, 2},
		InputB: "mqblbtpvicqhbrejb",
		Output: []int{1, 100},
	},
	{
		InputA: []int{7, 5, 4, 7, 10, 7, 9, 4, 8, 9, 6, 5, 4, 2, 3, 10, 9, 9, 3, 7, 5, 2, 9, 4, 8, 9},
		InputB: "zlrovckbgjqofmdzqetflraziyvkvcxzahzuuveypqxmjbwrjvmpdxjuhqyktuwjvmbeswxuznumazgxvitdrzxmqzhaaudztgie",
		Output: []int{7, 69},
	},
}

func Test_numberOfLines(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v, %v, Output: %v\n", tc.InputA, tc.InputB, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := numberOfLines(tc.InputA, tc.InputB)
			if !reflect.DeepEqual(output, tc.Output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_numberOfLines(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = numberOfLines(testCases[0].InputA, testCases[0].InputB)
	}
}
