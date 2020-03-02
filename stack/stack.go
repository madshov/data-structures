package stack

import (
	"fmt"
)

// NewStack creates a new instance of a stack
func NewStack() *Stack {
	return &Stack{nil}
}

type Stack struct {
	top *node
}

type node struct {
	next  *node
	value int
}

func (s *Stack) IsEmpty() bool {
	return s.top == nil
}

func (s *Stack) Peek() int {
	if s.top != nil {
		return s.top.value
	}

	return 0
}

func (s *Stack) Push(val int) {
	n := node{nil, val}
	n.next = s.top
	s.top = &n
}

func (s *Stack) Pop() int {
	if s.top != nil {
		val := s.top.value
		s.top = s.top.next
		return val
	}

	return 0
}

func (s *Stack) Print() {
	current := s.top
	if current != nil {
		fmt.Printf("%d ", current.value)
		for current.next != nil {
			current = current.next
			fmt.Printf("%d ", current.value)
		}
	}
}

/*
func main() {
	s := stack{nil}
	s.push(4)
	s.push(5)
	s.push(6)

	s.print()
	fmt.Println()
	fmt.Printf("%v\n", s.pop())
	s.print()
	fmt.Println()
	fmt.Printf("%v\n", s.pop())
	fmt.Printf("%v\n", s.pop())
	s.print()
}
*/
