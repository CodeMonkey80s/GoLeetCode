package solution83

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
	Val:  1,
	Next: &listNodeA3,
}

var listNodeA3 = ListNode{
	Val:  2,
	Next: nil,
}

var testCases = []struct {
	InputList *ListNode
	Output    []int
}{
	// Mandatory Test Cases
	{
		InputList: &listNodeA1,
		Output:    []int{1, 2},
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
			list := deleteDuplicates(tc.InputList)
			output := getListValues(list)
			if !reflect.DeepEqual(tc.Output, output) {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
