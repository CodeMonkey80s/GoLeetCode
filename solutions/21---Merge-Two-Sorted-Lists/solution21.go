package solution21

// ============================================================================
// 21. Merge Two Sorted Lists
// URL: https://leetcode.com/problems/merge-two-sorted-lists/
// ============================================================================

type ListNode struct {
	Val  int
	Next *ListNode
}

// mergeTwoLists merges two given lists into one
//
// This recursive approach is described here:
// https://www.geeksforgeeks.org/merge-two-sorted-linked-lists/
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var head *ListNode
	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	}
	if list1.Val <= list2.Val {
		head = list1
		head.Next = mergeTwoLists(list1.Next, list2)
	} else {
		head = list2
		head.Next = mergeTwoLists(list1, list2.Next)
	}
	return head
}
