package solution104

import (
	"fmt"
	"testing"
)

var treeNodeRoot = TreeNode{
	Val:   3,
	Left:  &treeNodeA1,
	Right: &treeNodeB1,
}

var treeNodeA1 = TreeNode{
	Val:   3,
	Left:  nil,
	Right: nil,
}

var treeNodeB1 = TreeNode{
	Val:   20,
	Left:  &treeNodeA2,
	Right: &treeNodeB2,
}

var treeNodeA2 = TreeNode{
	Val:   15,
	Left:  nil,
	Right: nil,
}

var treeNodeB2 = TreeNode{
	Val:   7,
	Left:  nil,
	Right: nil,
}

var testCases = []struct {
	Input  *TreeNode
	Output int
}{
	// Mandatory Test Cases
	{
		Input:  &treeNodeRoot,
		Output: 3,
	},
	// Additional my custom cases
}

func Test_maxDepth(t *testing.T) {
	for _, tc := range testCases {
		label := fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := maxDepth(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_maxDepth(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = maxDepth(testCases[0].Input)
	}
}
