package bintree

import (
	"fmt"
)

type node struct {
	left  *node
	right *node
	value int
}

// NewBinTree creates a new instance of a root node
func NewBinTree(rootVal int) *node {
	return &node{nil, nil, rootVal}
}

func (n *node) Insert(val int) {
	if val <= n.value {
		if n.left == nil {
			node := node{nil, nil, val}
			n.left = &node
		} else {
			n.left.Insert(val)
		}
	} else {
		if n.right == nil {
			node := node{nil, nil, val}
			n.right = &node
		} else {
			n.right.Insert(val)
		}
	}
}

func (n *node) Contains(val int) bool {
	if n.value == val {
		return true
	}

	if n.value > val {
		if n.left != nil {
			return n.left.Contains(val)
		}
	} else {
		if n.right != nil {
			return n.right.Contains(val)
		}
	}

	return false
}

func (n *node) InOrder() {
	if n.left != nil {
		n.left.InOrder()
	}
	fmt.Printf("%d ", n.value)
	if n.right != nil {
		n.right.InOrder()
	}
}

func (n *node) PreOrder() {
	fmt.Printf("%d ", n.value)
	if n.left != nil {
		n.left.PreOrder()
	}
	if n.right != nil {
		n.right.PreOrder()
	}
}

func (n *node) PostOrder() {
	if n.left != nil {
		n.left.PostOrder()
	}
	if n.right != nil {
		n.right.PostOrder()
	}
	fmt.Printf("%d ", n.value)
}

/*
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
*/
