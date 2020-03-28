// Package bstree includes a common set of functions on a binary search tree. A
// binary search tree is defined by a set nodes that satifies the
// binary-search-tree propery:
//     Let n be a node in a binary search tree. If m is a node in the left
//     subtree of n, then m.value <= n.value. If node m is in the right subtree
//     of n, then m.value > n.value.
// A node consists of a pointer to a parent node, along with two pointers to the
// node left child and right child. Only the root node will have a nil parent
// node, whereas any node can have either a nil left child or a nil right child.
// Leaf nodes will have both. Finally, each node holds a intervalue value or
// key.
// Example of an instance of a binary search tree with 11 nodes:
//
//						15
//					  /    \
//					 6      18
//					/ \    /  \
//				   3   7  17  20
//				  / \   \
//				 2   4  13
//				       /
//				      9
package bstree

import (
	"fmt"
)

// TreeRec defines a tree structure for recursive implementation.
type TreeRec struct {
	root  *Node
	count int
}

// Node defines a node of the binary search tree.
type Node struct {
	parent *Node
	left   *Node
	right  *Node
	value  int
}

// NewBSTreeRec creates a new instance of a binary search tree with a recursive
// approach to the implementation.
func NewBSTreeRec() *TreeRec {
	return &TreeRec{}
}

// Insert adds a node n to the correct position in the tree, by following the
// binary-search-tree property. The path is traced down the tree, until the
// parent of n is found. A simple value comparison is between n's value and the
// parent is used to determine if n will be the left or right child of the
// parent.
func (t *TreeRec) Insert(val int) {
	r := t.root
	if r == nil {
		t.root = &Node{nil, nil, nil, val}
	} else {
		par := t.ins(r, val)
		n := &Node{par, nil, nil, val}
		if val <= par.value {
			par.left = n
		} else {
			par.right = n
		}
	}

	t.count++
}

// ins is a helper function that traces down the tree and returns the first node
// with a nil child that satisfies the binary-search-tree property.
func (t *TreeRec) ins(n *Node, val int) *Node {
	if val <= n.value {
		if n.left != nil {
			return t.ins(n.left, val)
		}

		return n
	}

	if n.right != nil {
		return t.ins(n.right, val)
	}

	return n
}

// Delete removes a node n with a given value in the tree. This function is
// split into 3 cases, n has no children, n has one child or n has two children.
// If n has no children, n will just be deleted. Else if n has one child, the
// child is copied to n, and the child is deleted. Else n must has two children,
// and n must have a successor which is copied to n and then deleted.
func (t *TreeRec) Delete(val int) {
	r := t.root
	if r != nil {
		t.del(r, val)
		t.count--
	}
}

// del is a helper function that traces down the tree and finds the first node
// with the given value. Depending on the number of children, the respective
// operation is done.
func (t *TreeRec) del(n *Node, val int) *Node {
	if val < n.value {
		n.left = t.del(n.left, val)
		return n
	}

	if val > n.value {
		n.right = t.del(n.right, val)
		return n
	}

	if n.left == nil {
		return n.right
	}

	if n.right == nil {
		return n.left
	}

	m := t.Successor(n)
	n.value = m.value
	n.right = t.del(n.right, n.value)
	return n
}

// Search begins its search from node n, and traces a simple path downwards in
// the tree. For each node it encounters, it compares the search value with n's
// value. If they equal, the search terminates. If the search value is less
// than, the search continues along n's left subtree. Symmetrically, if the
// search value is greater than n's value it continues along n's right subtree.
// The binary-search-tree property ensures that this is the correct, and that
// if the search value exists, it will be found.
func (t *TreeRec) Search(n *Node, val int) *Node {
	if val == n.value {
		return n
	}

	if val < n.value {
		if n.left != nil {
			return t.Search(n.left, val)
		}
	}

	if val > n.value {
		if n.right != nil {
			return t.Search(n.right, val)
		}
	}

	return nil
}

// Min finds the node with the minimum value of a given node n's subtree, i.e.
// the leftmost node in the subtree.
func (t *TreeRec) Min(n *Node) *Node {
	if n != nil {
		if n.left != nil {
			return t.Min(n.left)
		}
	}

	return n
}

// Max finds the node with the maximum value of a given node n's subtree, i.e.
// the rightmost node in the substree.
func (t *TreeRec) Max(n *Node) *Node {
	if n != nil {
		if n.right != nil {
			return t.Max(n.right)
		}
	}

	return n
}

// Successor finds the next node for a given node n. The function is split into
// two cases. If the right subtree of node n is not empty, then the successor
// of n is just the leftmost node in n's right substree.
// If the right subtree of node n is empty, then either n does not have a
// successor or the successor is the lowest ancestor of n whose left child is
// also an ancestor of n.
func (t *TreeRec) Successor(n *Node) *Node {
	if n != nil {
		if n.right != nil {
			return t.Min(n.right)
		}

		if n.parent != nil {
			return t.suc(n, n.parent)
		}
	}

	return nil
}

// Predecessor finds the previous node for a given node n. The function is split
// into two cases. If the left subtree of node n is not empty, then the
// predecessor of n is just the rightmost node in n's left subtree.
// If the left substree of node n is empty, then either n does not have a
// predecessor or the predecessor is the lowest ancestor of n whose right child
// is also an ancestor of n.
func (t *TreeRec) Predecessor(n *Node) *Node {
	if n != nil {
		if n.left != nil {
			return t.Max(n.left)
		}

		if n.parent != nil {
			return t.pre(n, n.parent)
		}
	}

	return nil
}

// suc is a helper function that traces up the tree from a given node n, and
// checks if n is the right child of its parent. When this is false, i.e. n is a
// left child, the function returns n's parent.
func (t *TreeRec) suc(n, par *Node) *Node {
	if n.value == par.right.value {
		return t.suc(n.parent, par.parent)
	}

	return par
}

// pre is a helper function that traces up the tree from a given node n, and
// checks if n is the left child of its parent. When this is false, i.e. n is a
// right child, the function returns n's parent.
func (t *TreeRec) pre(m, par *Node) *Node {
	if m.value == par.left.value {
		return t.pre(m.parent, par.parent)
	}

	return par
}

// InOrder prints an inordered list of the values of node n's subtree.
// This results in an ascending ordered list of the values.
func (t *TreeRec) InOrder(n *Node) {
	if n != nil {
		if n.left != nil {
			t.InOrder(n.left)
		}

		fmt.Printf("%d ", n.value)

		if n.right != nil {
			t.InOrder(n.right)
		}
	}
}

// PreOrder prints a preordered list of the values of node n's subtree.
func (t *TreeRec) PreOrder(n *Node) {
	if n != nil {
		fmt.Printf("%d ", n.value)

		if n.left != nil {
			t.PreOrder(n.left)
		}

		if n.right != nil {
			t.PreOrder(n.right)
		}
	}
}

// PostOrder prints a postordered list of the values of node n's subtree.
func (t *TreeRec) PostOrder(n *Node) {
	if n != nil {
		if n.left != nil {
			t.PostOrder(n.left)
		}

		if n.right != nil {
			t.PostOrder(n.right)
		}

		fmt.Printf("%d ", n.value)
	}
}

// Size returns the total number of the nodes in the entire tree.
func (t *TreeRec) Size() int {
	return t.count
}
