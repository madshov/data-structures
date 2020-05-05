// Package stack includes a common set of operations for a stack data structure,
// i.e. peek, push, pop, as well as a stack traversal operation.
package elementary

import "errors"

// Various errors a stack function can return.
var (
	ErrStackUnderflow = errors.New("stack underflow")
)

// NewStack creates a new instance of a Stack, and returns a pointer to it.
func NewStack() *Stack {
	return &Stack{}
}

// StackElement defines an element of the stack.
type StackElement struct {
	next  *StackElement
	value int
}

func (se *StackElement) Value() int {
	return se.value
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
		value: val,
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
