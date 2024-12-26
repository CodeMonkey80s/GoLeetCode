package solution1381

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

var testCases = []struct {
	InputA []string
	InputB [][]int
	Output []string
}{
	{
		InputA: []string{"CustomStack", "push", "push", "pop", "push", "push", "push", "increment", "increment", "pop", "pop", "pop", "pop"},
		InputB: [][]int{{3}, {1}, {2}, {}, {2}, {3}, {4}, {5, 100}, {2, 100}, {}, {}, {}, {}},
		Output: []string{"null", "null", "null", "2", "null", "null", "null", "null", "null", "103", "202", "201", "-1"},
	},
}

func Test_customStack(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v %v, Output: %v\n", tc.InputA, tc.InputB, tc.Output)
		t.Run(label, func(t *testing.T) {
			var output []string
			obj := Constructor(tc.InputB[0][0])
			output = append(output, "null")
			for i, cmd := range tc.InputA {
				switch {
				case cmd == "push":
					obj.Push(tc.InputB[i][0])
					output = append(output, "null")
				case cmd == "pop":
					n := obj.Pop()
					output = append(output, strconv.Itoa(n))
				case cmd == "increment":
					obj.Increment(tc.InputB[i][0], tc.InputB[i][1])
					output = append(output, "null")
				}
			}
			if !reflect.DeepEqual(output, tc.Output) {
				t.Errorf("Expected output to be %#v but we got %#v", tc.Output, output)
			}
		})
	}
}
