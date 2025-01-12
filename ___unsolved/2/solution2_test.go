package solution2

import (
	"fmt"
	"testing"
)

var (
	nodeA1 = ListNode{
		Val:  2,
		Next: &nodeA2,
	}
	nodeA2 = ListNode{
		Val:  4,
		Next: &nodeA3,
	}
	nodeA3 = ListNode{
		Val:  3,
		Next: nil,
	}
)

var (
	nodeB1 = ListNode{
		Val:  5,
		Next: &nodeB2,
	}
	nodeB2 = ListNode{
		Val:  6,
		Next: &nodeB3,
	}
	nodeB3 = ListNode{
		Val:  4,
		Next: nil,
	}
)

var (
	nodeC1 = ListNode{
		Val:  8,
		Next: &nodeC2,
	}
	nodeC2 = ListNode{
		Val:  0,
		Next: &nodeC3,
	}
	nodeC3 = ListNode{
		Val:  7,
		Next: nil,
	}
)

var testCases = []struct {
	InputA *ListNode
	InputB *ListNode
	Output *ListNode
}{
	{
		InputA: &nodeA1,
		InputB: &nodeB1,
		Output: &nodeC1,
	},
}

func Test_addTwoNumbers(t *testing.T) {
	var label string
	for _, tc := range testCases {
		label = fmt.Sprintf("Case: Input: %v %v, Output: %v\n", tc.InputA, tc.InputB, tc.Output)
		t.Run(label, func(t *testing.T) {
			output := addTwoNumbers(tc.InputA, tc.InputB)
			if output != tc.Output {
				t.Errorf("Expected output to be %v but we got %v", tc.Output, output)
			}
		})
	}
}
