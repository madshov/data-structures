package queue

import (
	"fmt"
)

// NewQueue creates a new instance of a Queue, and returns a pointer to it.
func NewQueue() *Queue {
	return &Queue{nil, nil}
}

// Queue holds a node pointers to first and last elements of the queue
// respectively.
type Queue struct {
	head *node
	tail *node
}

type node struct {
	next  *node
	value int
}

// IsEmpty returns whether queue is empty or not.
func (q *Queue) IsEmpty() bool {
	return q.head == nil
}

// Peek returns the value of the first element in the queue. The element will
// not be dequeued.
func (q *Queue) Peek() int {
	if q.head != nil {
		return q.head.value
	}

	return 0
}

// Enqueue adds an element to the queue.
func (q *Queue) Enqueue(val int) {
	node := node{nil, val}
	if q.tail != nil {
		q.tail.next = &node
	}
	q.tail = &node

	if q.head == nil {
		q.head = &node
	}
}

// Dequeue removes an element from the queue and returns its value.
func (q *Queue) Dequeue() int {
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

// Print prints a visual representation of the queue to stdout.
func (q *Queue) Print() {
	current := q.head
	if current != nil {
		fmt.Printf("%d ", current.value)
		for current.next != nil {
			current = current.next
			fmt.Printf("%d ", current.value)
		}
	}
}
