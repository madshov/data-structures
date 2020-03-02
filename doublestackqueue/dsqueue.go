package dsqueue

import (
	"github.com/madshov/data-structures/stack"
)

// NewDSQueue creates a new instance of a dsqueue
func NewDSQueue() *DSQueue {
	s0 := stack.NewStack()
	s1 := stack.NewStack()

	return &DSQueue{s0, s1}
}

type DSQueue struct {
	s0 *stack.Stack
	s1 *stack.Stack
}

func (q *DSQueue) IsEmpty() bool {
	return q.s0.IsEmpty() && q.s1.IsEmpty()
}

func (q *DSQueue) Peek() int {
	if !q.IsEmpty() {
		if q.s1.IsEmpty() {
			for {
				v := q.s0.Pop()
				if v == 0 {
					break
				}
				q.s1.Push(v)
			}
		}

		return q.s1.Peek()
	}

	return 0
}

func (q *DSQueue) Enqueue(val int) {
	q.s0.Push(val)
}

func (q *DSQueue) Dequeue() int {
	if !q.IsEmpty() {
		if q.s1.IsEmpty() {
			for {
				v := q.s0.Pop()
				if v == 0 {
					break
				}
				q.s1.Push(v)
			}
		}

		return q.s1.Pop()
	}

	return 0
}
