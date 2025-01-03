package solution101

import (
	"fmt"
	"testing"
)

var treeNodeRoot = TreeNode{
	Val:   1,
	Left:  &treeNodeA1,
	Right: &treeNodeA2,
}

var treeNodeA1 = TreeNode{
	Val:   2,
	Left:  &treeNodeB1,
	Right: &treeNodeB2,
}

var treeNodeA2 = TreeNode{
	Val:   2,
	Left:  &treeNodeB3,
	Right: &treeNodeB4,
}

var treeNodeB1 = TreeNode{
	Val:   3,
	Left:  nil,
	Right: nil,
}

var treeNodeB2 = TreeNode{
	Val:   4,
	Left:  nil,
	Right: nil,
}

var treeNodeB3 = TreeNode{
	Val:   4,
	Left:  nil,
	Right: nil,
}

var treeNodeB4 = TreeNode{
	Val:   3,
	Left:  nil,
	Right: nil,
}

var testCases = []struct {
	Input  *TreeNode
	Output bool
}{
	// Mandatory Test Cases
	{
		Input:  &treeNodeRoot,
		Output: true,
	},
	// Additional my custom cases
}

func Test_isSymmetric(t *testing.T) {
	for _, tc := range testCases {
		label := fmt.Sprintf("Case: Input: %v, Output: %v\n", tc.Input.Val, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := isSymmetric(tc.Input)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}

func Benchmark_maxDepth(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = isSymmetric(testCases[0].Input)
	}
}
