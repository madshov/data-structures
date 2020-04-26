package elementary

import (
	"fmt"
)

// NewLinkedList creates a new instance of a list
func NewLinkedList(val int) *list {
	return &list{
		&node{nil, val},
	}
}

type node struct {
	next  *node
	value int
}

type list struct {
	head *node
}

func (l *list) Append(val int) {
	current := l.head
	for current.next != nil {
		current = current.next
	}

	node := node{nil, val}
	current.next = &node
}

func (l *list) Prepend(val int) {
	new := node{l.head, val}
	l.head = &new
}

func (l *list) Delete(val int) {
	current := l.head
	if current.value == val {
		l.head = current.next
	} else {
		for current.next != nil {
			if current.next.value == val {
				current.next = current.next.next
				break
			}
			current = current.next
		}
	}
}

func (l *list) Print() {
	current := l.head
	fmt.Printf("%d ", current.value)
	for current.next != nil {
		current = current.next
		fmt.Printf("%d ", current.value)
	}
}

/*
func main() {
	head := node{nil,1}
	l := list{&head}
	l.append(2)
	l.append(4)
	l.append(10)
	l.append(3)
	l.append(17)
	l.prepend(5)
	l.prepend(4)
	l.print()
	l.delete(5)
	fmt.Println()
	l.print()
	l.delete(4)
	fmt.Println()
	l.print()
	l.delete(10)
	fmt.Println()
	l.print()
}
*/
