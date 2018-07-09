package main

import (
	"fmt"
)

type queue struct {
	head *node
	tail *node
}

type node struct {
	next *node
	value int
}

func (q *queue) isempty() bool {
	return q.head == nil
}

func (q *queue) peek() int {
	if q.head != nil {
		return q.head.value
	}
	
	return 0
}

func (q *queue) add(val int) {
	node := node{nil, val}
	if q.tail != nil {
		q.tail.next = &node
	}
	q.tail = &node

	if q.head == nil {
		q.head = &node
	}
}

func (q *queue) remove() int {
	if q.head != nil {
		val := q.head.value
		q.head = q.head.next
		if q.head == nil {
			q.tail = nil
		}
		return val
	}

	return 0
}

func (q *queue) print() {
	current := q.head
	if current != nil {
		fmt.Print(current.value)
		fmt.Printf(" ")
		for current.next != nil {
			current = current.next
			fmt.Print(current.value)
			fmt.Printf(" ")
		}
	}
}

func main() {
	q := queue{nil, nil}
	q.add(4)
	q.add(5)
	q.add(6)
	q.add(7)
	fmt.Printf("%v", q.remove())
	fmt.Printf("%v", q.remove())
	fmt.Printf("%v", q.head)
	fmt.Printf("%v\n", q.tail)
	q.print()
}