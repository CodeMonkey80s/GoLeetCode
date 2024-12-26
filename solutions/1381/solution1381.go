package solution1381

import (
	"slices"
)

type CustomStack struct {
	stack   []int
	maxSize int
}

func Constructor(maxSize int) CustomStack {
	s := CustomStack{
		stack:   make([]int, 0, maxSize),
		maxSize: maxSize,
	}
	return s
}

func (s *CustomStack) Push(x int) {
	if len(s.stack) == s.maxSize {
		return
	}
	s.stack = append(s.stack, x)
}

func (s *CustomStack) Pop() int {
	if len(s.stack) == 0 {
		return -1
	}
	item := s.stack[len(s.stack)-1]
	s.stack = slices.Delete(s.stack, len(s.stack)-1, len(s.stack))
	return item
}

func (s *CustomStack) Increment(k int, val int) {
	n := len(s.stack)
	if k < len(s.stack) {
		n = k
	}
	for i := 0; i < n; i++ {
		s.stack[i] += val
	}
}

/**
 * Your CustomStack object will be instantiated and called as such:
 * obj := Constructor(maxSize);
 * obj.Push(x);
 * param_2 := obj.Pop();
 * obj.Increment(k,val);
 */
