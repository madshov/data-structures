package main

import (
	"fmt"
)

type node struct {
	left  *node
	right *node
	value int
}

func (n *node) insert(val int) {
	if val <= n.value {
		if n.left == nil {
			node := node{nil,nil,val}
			n.left = &node
		} else {
			n.left.insert(val)
		}
	} else {
		if n.right == nil {
			node := node{nil,nil,val}
			n.right = &node
		} else {
			n.right.insert(val)
		}
	}
}

func (n *node) contains(val int) bool {
	if n.value == val {
		return true
	}
		
	if n.value > val {
		if n.left != nil {
			return n.left.contains(val)
		}
	} else {
		if n.right != nil {
			return n.right.contains(val)
		}
	}

	return false
}

func (n *node) inorder() {
	if n.left != nil {
		n.left.inorder()
	}
	fmt.Print(n.value)
	fmt.Print(" ")
	if n.right != nil {
		n.right.inorder()
	}
}

func (n *node) preorder() {
	fmt.Print(n.value)
	fmt.Print(" ")
	if n.left != nil {
		n.left.preorder()
	}
	if n.right != nil {
		n.right.preorder()
	}
}

func (n *node) postorder() {
	if n.left != nil {
		n.left.postorder()
	}
	if n.right != nil {
		n.right.postorder()
	}
	fmt.Print(n.value)
	fmt.Print(" ")
}

func main() {
	root := node{nil,nil,10}
	root.insert(3)
	root.insert(14)
	root.insert(5)
	root.insert(7)
	root.insert(1)
	root.insert(59)
	root.insert(34)
	root.insert(23)
	root.insert(19)
	//n1 := node{&n,nil,5}
	//fmt.Printf("%v", root)
	//fmt.Printf("%v", root.left)
	//fmt.Printf("%v", root.right)
	//fmt.Printf("%v", root.left.right)
	root.inorder()
	fmt.Println()
	root.preorder()
	fmt.Println()
	root.postorder()
	//fmt.Printf("%v", n1)
}