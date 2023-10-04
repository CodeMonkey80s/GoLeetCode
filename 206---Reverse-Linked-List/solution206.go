package solution206

// ============================================================================
// 206. Reverse Linked List
// URL: https://leetcode.com/problems/reverse-linked-list/
// ============================================================================

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var next *ListNode
	var prev *ListNode
	var cur *ListNode = head
	for cur != nil {
		next = cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	head = prev
	return head
}
