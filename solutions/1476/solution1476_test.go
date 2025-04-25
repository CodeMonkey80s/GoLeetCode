package solution1476

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

var testCases = []struct {
	InputCommands []string
	InputValues   [][][]int
	Output        []string
}{
	{
		InputCommands: []string{
			"SubrectangleQueries",
			"getValue",
			"updateSubrectangle",
			"getValue",
			"getValue",
			"updateSubrectangle",
			"getValue",
			"getValue",
		},
		InputValues: [][][]int{
			{{1, 2, 1}, {4, 3, 4}, {3, 2, 1}, {1, 1, 1}},
			{{0, 2}},
			{{0, 0, 3, 2, 5}},
			{{0, 2}},
			{{3, 1}},
			{{3, 0, 3, 2, 10}},
			{{3, 1}},
			{{0, 2}},
		},

		Output: []string{"null", "1", "null", "5", "5", "null", "10", "5"},
	},
}

func Test_SubrectangleQueries(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v %v, Output: %v\n", tc.InputCommands, tc.InputValues, tc.Output)
		t.Run(label, func(t *testing.T) {
			var output []string
			obj := Constructor(tc.InputValues[0])
			output = append(output, "null")
			for i, cmd := range tc.InputCommands {
				switch cmd {
				case "updateSubrectangle":
					obj.UpdateSubrectangle(tc.InputValues[i][0][0], tc.InputValues[i][0][1], tc.InputValues[i][0][2], tc.InputValues[i][0][3], tc.InputValues[i][0][4])
					output = append(output, "null")
				case "getValue":
					v := obj.GetValue(tc.InputValues[i][0][0], tc.InputValues[i][0][1])
					output = append(output, strconv.Itoa(v))
				}
			}
			if !reflect.DeepEqual(output, tc.Output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
