package solution344

import (
	"fmt"
	"reflect"
	"testing"
)

var testCases = []struct {
	Input  []byte
	Output []byte
}{
	{
		Input:  []byte{'a', 'b', 'c'},
		Output: []byte{'c', 'b', 'a'},
	},
}

func Test_reverse(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			reverseString(tc.Input)
			if !reflect.DeepEqual(tc.Input, tc.Output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, tc.Input)
			}
		})
	}
}

func Benchmark_reverse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reverseString(testCases[0].Input)
	}
}
