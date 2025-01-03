package solution2418

import (
	"fmt"
	"reflect"
	"testing"
)

var testCases = []struct {
	InputA []string
	InputB []int
	Output []string
}{
	{
		InputA: []string{"Mary", "John", "Emma"},
		InputB: []int{180, 165, 170},
		Output: []string{"Mary", "Emma", "John"},
	},
	{
		InputA: []string{"Alice", "Bob", "Bob"},
		InputB: []int{155, 185, 150},
		Output: []string{"Bob", "Alice", "Bob"},
	},
	{
		InputA: []string{"IEO", "Sgizfdfrims", "QTASHKQ", "Vk", "RPJOFYZUBFSIYp", "EPCFFt", "VOYGWWNCf", "WSpmqvb"},
		InputB: []int{17233, 32521, 14087, 42738, 46669, 65662, 43204, 8224},
		Output: []string{"EPCFFt", "RPJOFYZUBFSIYp", "VOYGWWNCf", "Vk", "Sgizfdfrims", "IEO", "QTASHKQ", "WSpmqvb"},
	},
}

func Test_sortPeople(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v %v Output: %v\n", tc.InputA, tc.InputB, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := sortPeople(tc.InputA, tc.InputB)
			if !reflect.DeepEqual(output, tc.Output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_sortPeople(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = sortPeople(testCases[2].InputA, testCases[2].InputB)
	}
}
