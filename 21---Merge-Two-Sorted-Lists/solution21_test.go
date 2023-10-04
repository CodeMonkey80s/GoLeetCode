package solution21

import (
	"fmt"
	"reflect"
	"testing"
)

var listNodeA1 = ListNode{
	Val:  1,
	Next: &listNodeA2,
}

var listNodeA2 = ListNode{
	Val:  2,
	Next: &listNodeA3,
}

var listNodeA3 = ListNode{
	Val:  4,
	Next: nil,
}

var listNodeB1 = ListNode{
	Val:  1,
	Next: &listNodeB2,
}

var listNodeB2 = ListNode{
	Val:  3,
	Next: &listNodeB3,
}

var listNodeB3 = ListNode{
	Val:  4,
	Next: nil,
}

var testCases = []struct {
	InputList1 *ListNode
	InputList2 *ListNode
	Output     []int
}{
	// Mandatory Test Cases
	{
		InputList1: &listNodeA1,
		InputList2: &listNodeB1,
		Output:     []int{1, 1, 2, 3, 4, 4},
	},
	// Additional my custom cases
}

func getListValues(list *ListNode) []int {
	var s []int
	for list.Next != nil {
		s = append(s, list.Val)
		list = list.Next
	}
	s = append(s, list.Val)
	return s
}

func Test_mergeTwoLists(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Two Input Lists, Output: %v\n", tc.Output)
		t.Run(label, func(t *testing.T) {
			list := mergeTwoLists(tc.InputList1, tc.InputList2)
			output := getListValues(list)
			if !reflect.DeepEqual(tc.Output, output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
