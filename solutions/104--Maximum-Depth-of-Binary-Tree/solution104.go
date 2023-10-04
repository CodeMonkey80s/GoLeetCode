package solution104

// ============================================================================
// 104. Maximum Depth of Binary Tree
// URL: https://leetcode.com/problems/maximum-depth-of-binary-tree/
// ============================================================================

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	a := maxDepth(root.Left)
	b := maxDepth(root.Right)
	if a > b {
		a++
		return a
	} else {
		b++
		return b
	}
}
