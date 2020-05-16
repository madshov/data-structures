package elementary

import "errors"

// Various errors a stack function can return.
var (
	ErrStackUnderflow = errors.New("stack underflow")
)

// NewStack creates a new instance of a Stack data structure. It's a basic stack
// that implements a LIFO policy, with the basic operations peek, pop and push.
// Stack contains a top pointer to the top element of the stack. When an element
// is pushed, it takes its place at the top of the stack. When an element is
// popped, it is always from the top of the stack. Stack has no upper bound on
// the number of elements, so it cannot overflow. Attempts to pop from an empty
// stack will cause the tack to underflow. All three operations are done in O(1)
// time.
func NewStack() *Stack {
	return &Stack{}
}

// StackElement defines an element of the stack.
type StackElement struct {
	next  *StackElement
	Value int
}

// Stack defines a stack structure with a top element and a count of the
// elements in it.
type Stack struct {
	top   *StackElement
	count int
}

// IsEmpty checks if the stack is empty.
func (s *Stack) IsEmpty() bool {
	return s.top == nil
}

// Peek returns the top element of the stack. The element is not pushed.
func (s *Stack) Peek() *StackElement {
	return s.top
}

// Push adds an element to the top of the stack.
func (s *Stack) Push(val int) {
	e := &StackElement{
		next:  s.top,
		Value: val,
	}

	s.top = e
	s.count++
}

// Pop removes and returns the top element of the stack, unless the stack
// underflows.
func (s *Stack) Pop() (*StackElement, error) {
	if s.IsEmpty() {
		return nil, ErrStackUnderflow
	}

	e := s.top
	s.top = e.next
	s.count--
	return e, nil
}

// Traverse loops through each element in the stack.
func (s *Stack) Traverse(f func(*StackElement)) {
	e := s.top
	if e != nil {
		f(e)

		for e.next != nil {
			e = e.next
			f(e)
		}
	}
}

// Size returns the total number elements in the stack.
func (s *Stack) Size() int {
	return s.count
}
