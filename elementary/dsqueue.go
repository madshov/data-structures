package elementary

// NewDSQueue creates a new instance of a queue structure based on two stacks
// called stack0 and stack1. Like Queue, DSQueue implements a LIFO policy, with
// the basic operations peek, enqueue and dequeue. DSQueue contains two stacks
// called S0 and s1. s0 will be used as the main stack, while s1 will be used to
// queue elements until s0 is empty. Elements may therefore, be popped from s0
// and pushed to s1 within the same operation. The LIFO policy will always be
// maintained. The enqueue operation is done in O(1) time, while both peek and
// dequeue are done in O(n), so there is no avantage in using DSQueue compared
// to using Queue.
func NewDSQueue() *DSQueue {
	s0 := NewStack()
	s1 := NewStack()

	return &DSQueue{
		s0: s0,
		s1: s1,
	}
}

// DSQueue defines a queue structure with with two stackes and a count of the
// elements in it.
type DSQueue struct {
	s0    *Stack
	s1    *Stack
	count int
}

// IsEmpty checks if the queue is empty.
func (q *DSQueue) IsEmpty() bool {
	return q.s0.IsEmpty() && q.s1.IsEmpty()
}

// Peek returns the head element of the queue. The element is not dequeued. The
// head element is the top element of s1. If empty, all elements of s0, must be
// pushed to s1 first.
func (q *DSQueue) Peek() *StackElement {
	if q.s1.IsEmpty() {
		n, _ := q.s0.Pop()
		for n != nil {
			q.s1.Push(n.Value)
			n, _ = q.s0.Pop()
		}
	}

	return q.s1.Peek()
}

// Enqueue adds an element to the tail of the queue, by pushing it on to s0.
func (q *DSQueue) Enqueue(val int) {
	q.s0.Push(val)
	q.count++
}

// Dequeue removes and returns the head element of the queue, unless the queue
// underflows. The head element is the top element of s1. If empty, all elements
// of s0, must be pused to stack1 first.
func (q *DSQueue) Dequeue() (*StackElement, error) {
	if q.s1.IsEmpty() {
		e, _ := q.s0.Pop()
		for e != nil {
			q.s1.Push(e.Value)
			e, _ = q.s0.Pop()
		}
	}

	e, err := q.s1.Pop()
	if err != nil {
		return nil, err
	}

	q.count--
	return e, nil
}

// Traverse loops through each element in the queue. Elements will be moved
// around between s0 and s1, but the initial state will reoptained at the end.
func (q *DSQueue) Traverse(f func(*StackElement)) {
	var c int

	// pop and call function on each element
	// in s1 and push to s0
	for {
		e, _ := q.s1.Pop()
		if e == nil {
			break
		} else {
			f(e)
			q.s0.Push(e.Value)
			c++
		}
	}

	// restore s1 to initial state
	for c > 0 {
		e, _ := q.s0.Pop()
		q.s1.Push(e.Value)
		c--
	}

	// pop all elements from s0 and push
	// to s1
	for {
		e, _ := q.s0.Pop()
		if e == nil {
			break
		} else {
			q.s1.Push(e.Value)
			c++
		}
	}

	// pop and call function on each pushed
	// element in s1 and push to stack to
	// restore to initial state
	for c > 0 {
		e, _ := q.s1.Pop()
		f(e)
		q.s0.Push(e.Value)
		c--
	}
}

// Size returns the total number of elements in the queue.
func (q *DSQueue) Size() int {
	return q.count
}
