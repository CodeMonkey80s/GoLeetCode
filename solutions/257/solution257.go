package solution257

import (
	"strconv"
)

// ============================================================================
// 257. Binary Tree Paths
// URL: https://leetcode.com/problems/binary-tree-paths/
// ============================================================================

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var list []string

func binaryTreePaths(root *TreeNode) []string {
	list = []string{}
	if root == nil {
		return []string{}
	}
	addPath(root, "")
	return list
}

func addPath(root *TreeNode, s string) {
	if root.Left == nil && root.Right == nil {
		s += strconv.Itoa(root.Val)
		list = append(list, s)
		return
	}
	if root.Left == nil {
		s += strconv.Itoa(root.Val) + "->"
		addPath(root.Right, s)
		return
	}
	if root.Right == nil {
		s += strconv.Itoa(root.Val) + "->"
		addPath(root.Left, s)
		return
	}
	v := strconv.Itoa(root.Val)
	s += v + "->"
	addPath(root.Left, s)
	addPath(root.Right, s)
}
