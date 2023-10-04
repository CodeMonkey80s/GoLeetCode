package solution101

// ============================================================================
// 101. Symmetric Tree
// URL: https://leetcode.com/problems/symmetric-tree/
// ============================================================================

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isMirror(root1 *TreeNode, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}
	if root1 != nil && root2 != nil {
		if root1.Val == root2.Val {
			return isMirror(root1.Left, root2.Right) && isMirror(root1.Right, root2.Left)
		}
	}
	return false
}

func isSymmetric(root *TreeNode) bool {
	return isMirror(root, root)
}
