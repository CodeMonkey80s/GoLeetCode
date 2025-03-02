package solution3408

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

var testCases = []struct {
	InputA []string
	InputB [][][]int
	Output []string
}{
	{
		InputA: []string{"TaskManager", "add", "edit", "execTop", "rmv", "add", "execTop"},
		InputB: [][][]int{
			{{1, 101, 10}, {2, 102, 20}, {3, 103, 15}},
			{{4, 104, 5}},
			{{102, 8}},
			{{}},
			{{101}},
			{{5, 105, 15}},
			{{}},
		},
		Output: []string{"null", "null", "null", "3", "null", "null", "5"},
	},
}

func Test_taskManager(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v %v, Output: %v\n", tc.InputA, tc.InputB, tc.Output)
		t.Run(label, func(t *testing.T) {
			var output []string
			obj := Constructor(tc.InputB[0])
			output = append(output, "null")
			for i, cmd := range tc.InputA {
				switch {
				case cmd == "add":
					obj.Add(tc.InputB[i][0][0], tc.InputB[i][0][1], tc.InputB[i][0][2])
					output = append(output, "null")
				case cmd == "edit":
					obj.Edit(tc.InputB[i][0][0], tc.InputB[i][0][1])
					output = append(output, "null")
				case cmd == "rmv":
					obj.Rmv(tc.InputB[i][0][0])
					output = append(output, "null")
				case cmd == "execTop":
					userId := obj.ExecTop()
					output = append(output, strconv.Itoa(userId))
				}
			}
			if !reflect.DeepEqual(output, tc.Output) {
				t.Errorf("Expected output to be %#v but we got %#v", tc.Output, output)
			}
		})
	}
}
