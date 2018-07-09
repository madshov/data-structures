package main

import (
	"fmt"
)

type node struct {
	next *node
	value int
}

type list struct {
	head *node
}

func (l *list) append(val int) {
	current := l.head
	for current.next != nil {
		current = current.next
	}

	node := node{nil, val}
	current.next = &node
}

func (l *list) prepend(val int) {
	new := node{l.head, val}
	l.head = &new
}

func (l *list) delete(val int) {
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

func (l *list) print() {
	current := l.head
	fmt.Print(current.value)
	fmt.Print(" ");

	for current.next != nil {
		current = current.next
		fmt.Print(current.value);
		fmt.Print(" ")
	}
}

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