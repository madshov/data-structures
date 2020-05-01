package elementary

import (
	"errors"
)

// Various errors a list function can return.
var (
	ErrDeleteSentinel = errors.New("cannot delete sentinel of list")
)

// NewLinkedList creates a new instance of a list
func NewLinkedList() *List {
	s := &ListElement{
		value: 0,
	}
	s.next = s
	s.prev = s

	return &List{
		head: s,
	}
}

type List struct {
	head *ListElement
}

type ListElement struct {
	next  *ListElement
	prev  *ListElement
	value int
}

// IsEmpty checks if the list is empty.
func (l *List) IsEmpty() bool {
	return l.head.next == l.head.prev
}

func (l *List) Insert(val int) {
	e := &ListElement{
		next:  l.head.next,
		prev:  l.head,
		value: val,
	}

	l.head.next.prev = e
	l.head.next = e
}

func (l *List) Search(val int) *ListElement {
	e := l.head.next
	for e != l.head {
		if e.value == val {
			return e
		}
		e = e.next
	}

	return nil
}

func (l *List) Delete(e *ListElement) error {
	if e == l.head {
		return ErrDeleteSentinel
	}

	e.prev.next = e.next
	e.next.prev = e.prev

	return nil
}

// Traverse loops through each element in the list until it reaches the
// sentinel.
func (l *List) Traverse(f func(*ListElement)) {
	e := l.head.next
	for e != l.head {
		f(e)
		e = e.next
	}
}
