package solution83

// ============================================================================
// 83. Remove Duplicates from Sorted List
// URL: https://leetcode.com/problems/remove-duplicates-from-sorted-list/
// ============================================================================

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	var cur = head
	for cur != nil {
		tmp := cur
		for tmp != nil && tmp.Val == cur.Val {
			tmp = tmp.Next
		}
		cur.Next = tmp
		cur = cur.Next
	}
	return head
}
