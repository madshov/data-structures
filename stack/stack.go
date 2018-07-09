package main

import (
	"fmt"
)

type stack struct {
	top *node
}

type node struct {
	next *node
	value int
}

func (s *stack) isempty() bool {
	return s.top == nil
}

func (s *stack) peek() int {
	if s.top != nil {
		return s.top.value
	}

	return 0
}

func (s *stack) push(val int) {
	n := node{nil, val}	
	n.next = s.top
	s.top = &n
}

func (s *stack) pop() int {
	if s.top != nil {
		val := s.top.value
		s.top = s.top.next
		return val
	}
	
	return 0
}

func (s *stack) print() {
	current := s.top
	if current != nil {
		fmt.Print(current.value)
		fmt.Print(" ")
		for current.next != nil {
			current = current.next
			fmt.Print(current.value)
			fmt.Print(" ")
		}
	}
}

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