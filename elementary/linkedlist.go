package elementary

import (
	"errors"
)

// Various errors a list function can return.
var (
	ErrDeleteSentinel = errors.New("cannot delete sentinel of list")
)

// NewLinkedList creates a new instance of a linked list data structure, which
// is just an arrangement of elements in a linear order. The list is doubly
// linked, i.e. each element has a pointer to its next and previous elements,
// thereby providing a simple, but flexible representation of a dynamic set.
// Further, the list contains a dummy element, called the sentinel. The sentinel
// always lies between the head and the tail of the list - its previous pointer
// will point the tail, while the tail's next pointer points to the sentinel,
// thereby providing a circular, doubly linked list. This allows for simplifying
// each list operation's boundary conditions, but adds an extra element, thus
// adding to increased memory usage.
func NewLinkedList() *List {
	// create sentinel value
	s := &ListElement{
		Key:   "",
		Value: 0,
	}
	s.next = s
	s.prev = s

	// insert sentinel into the list
	return &List{
		sent: s,
	}
}

// List defines a queue structure with a head and tail element and a count of
// the elements in it.
type List struct {
	sent *ListElement
}

// ListElement defines an element of the list.
type ListElement struct {
	next  *ListElement
	prev  *ListElement
	Key   string
	Value int
}

// IsEmpty checks if the list is empty, by comparing the sentinel's next and
// previous values.
func (l *List) IsEmpty() bool {
	return l.sent.next == l.sent.prev
}

// Insert adds a new element with a given value to the list, by inserting it
// right after the sentinel.
func (l *List) Insert(key string, val int) {
	el := &ListElement{
		next:  l.sent.next,
		prev:  l.sent,
		Key:   key,
		Value: val,
	}

	// insert element between sentinel and the currently next element.
	l.sent.next.prev = el
	l.sent.next = el
}

// Search searches for an element with a given key by iteratively checking the
// next element of the list.
func (l *List) Search(key string) (*ListElement, bool) {
	el := l.sent.next
	for el != l.sent {
		if el.Key == key {
			return el, true
		}
		el = el.next
	}

	return nil, false
}

// Delete removes a given element from the list.
func (l *List) Delete(el *ListElement) error {
	if el == l.sent {
		return ErrDeleteSentinel
	}

	el.prev.next = el.next
	el.next.prev = el.prev

	return nil
}

// Traverse loops through each element in the list until it reaches the
// sentinel.
func (l *List) Traverse(f func(*ListElement)) {
	el := l.sent.next
	for el != l.sent {
		f(el)
		el = el.next
	}
}
