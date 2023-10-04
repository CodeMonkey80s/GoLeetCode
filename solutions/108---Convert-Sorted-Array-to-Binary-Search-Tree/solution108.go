package solution108

import (
	"fmt"
)

// ============================================================================
// 108. Convert Sorted Array to Binary Search Tree
// URL: https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree/
// ============================================================================

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	count := len(nums)

	s := make([]int, count*2)
	for i := 0; i < count; i++ {
		v := nums[i]
		fmt.Printf("i=%d\t%d\n", i, v)
	}
	fmt.Printf("%v\n", s)

	// n := &TreeNode{
	// 	Val: nums[size/2],
	// }
	// root := n
	// a, b := 0, 0
	// if size%2 != 0 {
	// 	for i := 0; i < size/2; i++ {
	// 		a = size/2 - i - 1
	// 		b = size/2 + i + 1
	// 		v1 := nums[a]
	// 		v2 := nums[b]
	// 		n1 := &TreeNode{
	// 			Val: v1,
	// 		}
	// 		n.Left = n1
	// 		n2 := &TreeNode{
	// 			Val: v2,
	// 		}
	// 		n.Right = n2
	// 		fmt.Printf("i=%d, %d %d, %d %d\n", i, a, b, v1, v2)
	// 	}
	// }

	root := &TreeNode{}
	return root
}
