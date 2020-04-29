package elementary

// NewDSQueue creates a new instance of a DSQueue, and returns a pointer to it.
func NewDSQueue() *DSQueue {
	s0 := NewStack()
	s1 := NewStack()

	return &DSQueue{s0, s1}
}

// DSQueue holds the two stacks necessary to mimic the queue.
type DSQueue struct {
	s0 *Stack
	s1 *Stack
}

// IsEmpty returns whether queue is empty or not.
func (q *DSQueue) IsEmpty() bool {
	return q.s0.IsEmpty() && q.s1.IsEmpty()
}

// Peek returns the value of the first element in the queue. The element will
// not be dequeued.
func (q *DSQueue) Peek() int {
	if !q.IsEmpty() {
		if q.s1.IsEmpty() {
			for {
				v, _ := q.s0.Pop()
				if v.value == 0 {
					break
				}
				q.s1.Push(v.value)
			}
		}

		return q.s1.Peek().value
	}

	return 0
}

// Enqueue adds an element to the queue.
func (q *DSQueue) Enqueue(val int) {
	q.s0.Push(val)
}

// Dequeue removes an element from the queue and returns its value.
func (q *DSQueue) Dequeue() int {
	if !q.IsEmpty() {
		if q.s1.IsEmpty() {
			for {
				v, _ := q.s0.Pop()
				if v.value == 0 {
					break
				}
				q.s1.Push(v.value)
			}
		}

		return q.s1.Pop().value
	}

	return 0
}
