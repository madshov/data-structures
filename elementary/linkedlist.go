package elementary

import (
	"errors"
)

// Various errors a list function can return.
var (
	ErrListUnderflow    = errors.New("list underflow")
	ErrNoElementWithVal = errors.New("no element with value exists in list")
)

// NewLinkedList creates a new instance of a list
func NewLinkedList() *List {
	return &List{
		head: nil,
	}
}

type List struct {
	head *Element
}

// IsEmpty checks if the list is empty.
func (l *List) IsEmpty() bool {
	return l.head == nil
}

func (l *List) Append(val int) {
	e := &Element{
		next:  nil,
		value: val,
	}

	h := l.head
	for h.next != nil {
		h = h.next
	}

	h.next = e
}

func (l *List) Prepend(val int) {
	e := &Element{
		next:  l.head,
		value: val,
	}

	l.head = e
}

func (l *List) Delete(val int) (*Element, error) {
	if l.IsEmpty() {
		return nil, ErrQueueUnderflow
	}

	if l.head.value == val {
		l.head = l.head.next
	} else {
		h := l.head
		for h.next != nil {
			e := h.next
			if e.value == val {
				h.next = e.next
				return e, nil
			}

			h = h.next
		}
	}

	return nil, ErrNoElementWithVal
}

// Traverse loops through each element in the list.
func (l *List) Traverse(f func(*Element)) {
	e := l.head
	if e != nil {
		f(e)

		for e.next != nil {
			e = e.next
			f(e)
		}
	}
}
