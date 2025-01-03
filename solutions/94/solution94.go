package solution94

// ============================================================================
// 94. Binary Tree Inorder Traversal
// URL: https://leetcode.com/problems/binary-tree-inorder-traversal/
// ============================================================================

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	v := make([]int, 0)
	if root == nil {
		return v
	}
	a := inorderTraversal(root.Left)
	v = append(v, a...)
	v = append(v, root.Val)
	b := inorderTraversal(root.Right)
	v = append(v, b...)
	return v
}
