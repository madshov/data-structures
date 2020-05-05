// Package queue includes a common set of operations for a queue data structure,
// i.e. peek, enqueue, dequeue, as well as a queue traversal operation.
package elementary

import "errors"

// Various errors a queue function can return.
var (
	ErrQueueUnderflow = errors.New("queue underflow")
)

// NewQueue creates a new instance of a Queue, and returns a pointer to it.
func NewQueue() *Queue {
	return &Queue{}
}

// Queue defines a queue structure with a head and tail element and a count of
// the elements in it.
type Queue struct {
	head  *QueueElement
	tail  *QueueElement
	count int
}

// QueueElement defines an element of the queue.
type QueueElement struct {
	next  *QueueElement
	value int
}

func (qe *QueueElement) Value() int {
	return qe.value
}

// IsEmpty checks if the queue is empty.
func (q *Queue) IsEmpty() bool {
	return q.head == nil
}

// Peek returns the head element of the queue. The element is not dequeued.
func (q *Queue) Peek() *QueueElement {
	return q.head
}

// Enqueue adds an element to the tail of the queue.
func (q *Queue) Enqueue(val int) {
	e := QueueElement{
		next:  nil,
		value: val,
	}

	if q.tail != nil {
		q.tail.next = &e
	} else {
		q.tail = &e
	}

	if q.head == nil {
		q.head = &e
	}

	q.count++
}

// Dequeue removes and returns the head element of the queue, unless the queue
// underflows.
func (q *Queue) Dequeue() (*QueueElement, error) {
	if q.IsEmpty() {
		return nil, ErrQueueUnderflow
	}

	e := q.head
	q.head = q.head.next

	if q.head == nil {
		q.tail = nil
	}
	q.count--
	return e, nil
}

// Traverse loops through each node in the queue.
func (q *Queue) Traverse(f func(*QueueElement)) {
	e := q.head
	if e != nil {
		f(e)

		for e.next != nil {
			e = e.next
			f(e)
		}
	}
}

// Size returns the total number of elements in the queue.
func (q *Queue) Size() int {
	return q.count
}
