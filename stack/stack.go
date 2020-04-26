// Package stack includes a common set of operations for a stack data structure,
// i.e. peek, push, pop, as well as a stack traversal operation.
package stack

// NewStack creates a new instance of a Stack, and returns a pointer to it.
func NewStack() *Stack {
	return &Stack{}
}

// Stack defines a stack structure with a top element and a count of the
// elements in it.
type Stack struct {
	top   *Element
	count int
}

// Element defines an element of the stack.
type Element struct {
	next  *Element
	value int
}

// IsEmpty checks if the stack is empty.
func (s *Stack) IsEmpty() bool {
	return s.top == nil
}

// Peek returns the top element of the stack. The element is not pushed.
func (s *Stack) Peek() *Element {
	return s.top
}

// Push adds an element to the top of the stack.
func (s *Stack) Push(val int) {
	e := &Element{
		next:  s.top,
		value: val,
	}

	s.top = e
	s.count++
}

// Pop removes and returns the top element of the stack, unless the stack
// underflows.
func (s *Stack) Pop() *Element {
	if !s.IsEmpty() {
		e := s.top
		s.top = e.next
		s.count--
		return e
	}

	return nil
}

// Traverse loops through each node in the stack.
func (s *Stack) Traverse(f func(*Element)) {
	e := s.top
	if e != nil {
		f(e)

		for e.next != nil {
			e = e.next
			f(e)
		}
	}
}

// Size returns the total number of the element in the stack.
func (s *Stack) Size() int {
	return s.count
}
