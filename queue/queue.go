// Package queue includes a common set of operations for a queue data structure,
// i.e. peek, enqueue, dequeue, as well as a queue traversal operation.
package queue

// NewQueue creates a new instance of a Queue, and returns a pointer to it.
func NewQueue() *Queue {
	return &Queue{}
}

// Queue defines a queue structure with a head and tail element and a count of
// the elements in it.
type Queue struct {
	head  *Element
	tail  *Element
	count int
}

// Element defines an element of the queue.
type Element struct {
	next  *Element
	value int
}

// IsEmpty checks if the queue is empty.
func (q *Queue) IsEmpty() bool {
	return q.head == nil
}

// Peek returns the head element of the queue. The element is not be dequeued.
func (q *Queue) Peek() *Element {
	return q.head
}

// Enqueue adds an element to the tail of the queue.
func (q *Queue) Enqueue(val int) {
	e := Element{
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
func (q *Queue) Dequeue() *Element {
	if !q.IsEmpty() {
		e := q.head
		q.head = q.head.next

		if q.head == nil {
			q.tail = nil
		}
		q.count--
		return e
	}

	return nil
}

// Traverse loops through each node in the queue.
func (q *Queue) Traverse(f func(*Element)) {
	e := q.head
	if e != nil {
		f(e)

		for e.next != nil {
			e = e.next
			f(e)
		}
	}
}

// Size returns the total number of the element in the queue.
func (q *Queue) Size() int {
	return q.count
}
