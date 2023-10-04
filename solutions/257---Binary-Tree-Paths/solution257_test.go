package solution257

import (
	"fmt"
	"reflect"
	"testing"
)

var listNodeRoot = TreeNode{
	Val:   1,
	Left:  &listNodeA1,
	Right: &listNodeA2,
}

var listNodeA1 = TreeNode{
	Val:   2,
	Left:  &listNodeB1,
	Right: nil,
}

var listNodeA2 = TreeNode{
	Val:   3,
	Left:  nil,
	Right: nil,
}

var listNodeB1 = TreeNode{
	Val:   5,
	Left:  nil,
	Right: nil,
}

var testCases = []struct {
	InputList *TreeNode
	Output    []string
}{
	// Mandatory Test Cases
	{
		InputList: &listNodeRoot,
		Output:    []string{"1->2->5", "1->3"},
	},
	// Additional my custom cases
}

func Test_binaryTreePaths(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input [*TreeNode], Output: %v\n", tc.Output)
		t.Run(label, func(t *testing.T) {
			output := binaryTreePaths(&listNodeRoot)
			if !reflect.DeepEqual(tc.Output, output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
