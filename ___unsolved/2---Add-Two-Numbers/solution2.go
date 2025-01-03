package solution2

import (
	"fmt"
	"log"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	number1 := 0
	depth := 1
	current := l1
	for {
		number1 += current.Val * depth
		fmt.Printf("Number: %d\n", current.Val)
		fmt.Printf("Total: %d\n", number1)
		depth *= 10
		current = current.Next
		if current == nil {
			break
		}
	}

	log.Println("A")

	number2 := 0
	depth = 1
	current = l2
	for {
		number2 += current.Val * depth
		fmt.Printf("Number: %d\n", current.Val)
		fmt.Printf("Total: %d\n", number2)
		depth *= 10
		current = current.Next
		if current == nil {
			break
		}
	}

	log.Println("B")

	var list []*ListNode

	sum := 0
	total := number1 + number2
	// num := total
	for total > 0 {
		sum = sum + int(total%10)
		total = int(total / 10)
		list = append(list, &ListNode{
			Val:  total,
			Next: nil,
		})
	}

	log.Println("C")

	return nil
}
