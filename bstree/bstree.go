// Package bstree includes a common set of operations for a binary search tree,
// i.e. insert, delete, various search operations as well as tree walk
// operations. A binary search tree is defined by a set nodes that satifies the
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
//                      15
//                    /    \
//                   6      18
//                  / \    /  \
//                 3   7  17  20
//                / \   \
//               2   4  13
//                     /
//                    9
//
// A binary search tree's structure is dependent on the order of the inserted
// nodes, but a nonempty tree with x nodes must have a height of at least
// ⌊lg x⌋ and at most x. The expected runtime of most of the basic operations is
// equivalent to the height of the tree, therefore making it Θ(x). On average
// a lower runtime of Θ(lg x) would be expected though, especially if the tree
// is randomly built.
// This package contains two similar functions for each search and tree walk
// operation - one which uses a recursive strategy, and another using an
// iterative strategy. Both types can be used interchangeably. The recursive
// functions tend to be more readable and requires less lines on average, but
// most iterative functions are slightly more efficient in terms of memory and
// runtime.
package bstree

// Tree defines a tree structure using a recursive implementation strategy.
type Tree struct {
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

// NewBSTree creates a new instance of a Tree.
func NewBSTree(par int) *Tree {
	return &Tree{}
}

// Insert adds a node n to the correct position in the tree, by following the
// binary-search-tree property. The path is traced down the tree, until the
// parent of n is found. A simple value comparison is between n's value and the
// parent is used to determine if n will be the left or right child of the
// parent.
func (t *Tree) Insert(val int) {
	r := t.root
	n := &Node{nil, nil, nil, val}
	var m *Node

	// trace down the tree to find the correct parent
	for r != nil {
		m = r
		if val < r.value {
			r = r.left
		} else {
			r = r.right
		}
	}

	n.parent = m

	if m == nil {
		t.root = n
	} else {
		if n.value < m.value {
			m.left = n
		} else {
			m.right = n
		}
	}

	t.count++
}

// Delete removes a node n with a given value in the tree.
func (t *Tree) Delete(n *Node) {
	if n == nil {
		return
	}

	if n.left == nil {
		// transplant n with n's right child
		m := n.right
		if n.parent == nil {
			t.root = m
		} else {
			// if n is the left child
			if n == n.parent.left {
				n.parent.left = m
			} else {
				n.parent.right = m
			}
		}

		if m != nil {
			m.parent = n.parent
		}
		return
	}

	if n.right == nil {
		// transplant n with n's left child
		m := n.left
		if n.parent == nil {
			t.root = m
		} else {
			// if n is the left child
			if n == n.parent.left {
				n.parent.left = m
			} else {
				n.parent.right = m
			}
		}

		if m != nil {
			m.parent = n.parent
		}
		return
	}

	m := t.MinIt(n)
	if m.parent != n {
		// transplant m with m's right child
		o := m.right
		if n.parent == nil {
			t.root = o
		} else {
			// if n is the left child
			if n == n.parent.left {
				n.parent.left = o
			} else {
				n.parent.right = o
			}
		}

		if o != nil {
			o.parent = n.parent
		}

		o = n.right
		o.parent = m
	}

	// transplant n with m
	if n.parent == nil {
		t.root = m
	} else {
		// if n is the left child
		if n == n.parent.left {
			n.parent.left = m
		} else {
			n.parent.right = m
		}
	}

	if m != nil {
		m.parent = n.parent
	}

	m.left = n.left
	m.left.parent = m

	t.count--
}

// Search begins its search from the root node, and traces a simple path
// downwards in the tree. For each node it encounters, it compares the search
// value with n's value. If they equal, the search terminates. If the search
// value is less than, the search continues along n's left subtree.
// Symmetrically, if the search value is greater than n's value it continues
// along n's right subtree. The binary-search-tree property ensures that this
// is the correct, and that if the search value exists, it will be found.
// Search is recursive and uses a helper function to maintain the recursive
// callstack on either the left or right subtree.
func (t *Tree) Search(val int) *Node {
	r := t.root
	if r == nil {
		return nil
	}

	return t.sch(r, val)
}

func (t *Tree) sch(n *Node, val int) *Node {
	if val == n.value {
		return n
	}

	if val < n.value {
		if n.left != nil {
			return t.sch(n.left, val)
		}
	}

	if val > n.value {
		if n.right != nil {
			return t.sch(n.right, val)
		}
	}

	return nil
}

// SearchIt is simliar to Search but is iterative and uses a for loop to
// maintain the callstack on the left or right subtree instead of a helper
// function.
func (t *Tree) SearchIt(val int) *Node {
	n := t.root

	for n != nil && val != n.value {
		if val < n.value {
			n = n.left
		} else {
			n = n.right
		}
	}

	return n
}

// Min finds the node with the minimum value of a given node n's subtree, i.e.
// the leftmost node in the subtree. Min is recursive, and calls itself on the
// left subtree.
func (t *Tree) Min(n *Node) *Node {
	if n == nil {
		n = t.root
	}

	if n.left != nil {
		return t.Min(n.left)
	}

	return n
}

// MinIt is similar to Min, but is iterative and uses a for loop on the left
// subtree.
func (t *Tree) MinIt(n *Node) *Node {
	if n == nil {
		n = t.root
	}

	for n.left != nil {
		n = n.left
	}

	return n
}

// Max finds the node with the maximum value of a given node n's subtree, i.e.
// the rightmost node in the substree. Max is recursive, and calls itself on the
// right subtree.
func (t *Tree) Max(n *Node) *Node {
	if n == nil {
		n = t.root
	}

	if n.right != nil {
		return t.Max(n.right)
	}

	return n
}

// MaxIt is similar to Max, but is iterative and uses a for loop on the right
// subtree.
func (t *Tree) MaxIt(n *Node) *Node {
	if n == nil {
		n = t.root
	}

	for n.right != nil {
		n = n.right
	}

	return n
}

// Successor finds the next node for a given node n. The function is split into
// two cases. If the right subtree of node n is not empty, then the successor
// of n is just the leftmost node in n's right substree.
// If the right subtree of node n is empty, then either n does not have a
// successor or the successor is the lowest ancestor of n whose left child is
// also an ancestor of n.
func (t *Tree) Successor(n *Node) *Node {
	if n == nil {
		n = t.root
	}

	if n.right != nil {
		return t.Min(n.right)
	}

	if n.parent != nil {
		return t.suc(n, n.parent)
	}

	return nil
}

// suc is a helper function that traces up the tree from a given node n, and
// checks if n is the right child of its parent. When this is false, i.e. n is a
// left child, the function returns n's parent.
func (t *Tree) suc(n, par *Node) *Node {
	if par.right != nil && n == par.right {
		if par.parent != nil {
			return t.suc(n.parent, par.parent)
		}

		return nil
	}

	return par
}

func (t *Tree) SuccessorIt(n *Node) *Node {
	if n == nil {
		n = t.root
	}

	if n.right != nil {
		return t.MinIt(n.right)
	}

	m := n.parent
	for m != nil && n == m.right {
		n = m
		m = n.parent
	}

	return m
}

// Predecessor finds the previous node for a given node n. The function is split
// into two cases. If the left subtree of node n is not empty, then the
// predecessor of n is just the rightmost node in n's left subtree.
// If the left substree of node n is empty, then either n does not have a
// predecessor or the predecessor is the lowest ancestor of n whose right child
// is also an ancestor of n.
func (t *Tree) Predecessor(n *Node) *Node {
	if n == nil {
		n = t.root
	}

	if n.left != nil {
		return t.Max(n.left)
	}

	if n.parent != nil {
		return t.pre(n, n.parent)
	}

	return nil
}

// pre is a helper function that traces up the tree from a given node n, and
// checks if n is the left child of its parent. When this is false, i.e. n is a
// right child, the function returns n's parent.
func (t *Tree) pre(m, par *Node) *Node {
	if par.left != nil && m == par.left {
		if par.parent != nil {
			return t.pre(m.parent, par.parent)
		}

		return nil
	}

	return par
}

func (t *Tree) PredecessorIt(n *Node) *Node {
	if n == nil {
		n = t.root
	}

	if n.left != nil {
		return t.MaxIt(n.left)
	}

	m := n.parent
	for m != nil && n == m.left {
		n = m
		m = n.parent
	}

	return m
}

// InOrder prints an inordered list of the values of node n's subtree. This
// results in an ascending ordered list of the values. InOrder is recursive, and
// calls itself on the left and right subtree, but calling the a function given
// as input between each of them.
func (t *Tree) InOrder(n *Node, f func(*Node)) {
	if n == nil {
		n = t.root
	}

	if n.left != nil {
		t.InOrder(n.left, f)
	}

	f(n)

	if n.right != nil {
		t.InOrder(n.right, f)
	}
}

func (t *Tree) InOrderIt(f func(*Node)) {
	n := t.root
	ns := make([]*Node, t.Size())
	top := 0

	for n != nil {
		for n != nil {
			if n.right != nil {
				ns[top] = n.right
				top++
			}

			ns[top] = n
			top++
			n = n.left
		}

		top--
		n = ns[top]

		for top > 0 && n.right == nil {
			f(n)

			top--
			n = ns[top]
		}

		f(n)

		if top > 0 {
			top--
			n = ns[top]
		} else {
			n = nil
		}
	}
}

// PreOrder prints a preordered list of the values of node n's subtree. PreOrder
// is recursive, and calls itself on the left and right subtree, after calling
// a function given as  input.
func (t *Tree) PreOrder(n *Node, f func(*Node)) {
	if n == nil {
		n = t.root
	}

	f(n)

	if n.left != nil {
		t.PreOrder(n.left, f)
	}

	if n.right != nil {
		t.PreOrder(n.right, f)
	}
}

func (t *Tree) PreOrderIt(f func(*Node)) {
	n := t.root
	ns := make([]*Node, t.Size())
	top := 0
	ns[top] = n
	top++

	for top > 0 {
		top--
		n = ns[top]

		f(n)

		if n.right != nil {
			ns[top] = n.right
			top++
		}

		if n.left != nil {
			ns[top] = n.left
			top++
		}
	}
}

// PostOrder prints a postordered list of the values of node n's subtree.
// PostOrder is recursive, and calls itself on the left and right subtree,
// before calling a function given as input.
func (t *Tree) PostOrder(n *Node, f func(*Node)) {
	if n == nil {
		n = t.root
	}

	if n.left != nil {
		t.PostOrder(n.left, f)
	}

	if n.right != nil {
		t.PostOrder(n.right, f)
	}

	f(n)
}

func (t *Tree) PostOrderIt(f func(*Node)) {
	n := t.root
	ns := make([]*Node, t.Size())
	top := 0
	flag := true

	for n != nil {
		for n != nil && flag {
			if n.right != nil {
				ns[top] = n.right
				top++
			}

			ns[top] = n
			top++
			n = n.left
		}

		top--
		n = ns[top]

		if top > 0 {
			if n.right != nil && ns[top-1].value == n.right.value {
				top--
				ns[top] = n
				top++

				n = n.right
				flag = true
			} else {
				f(n)
				flag = false
			}
		} else {
			f(n)
			n = nil
		}
	}
}

// Size returns the total number of the nodes in the entire tree.
func (t *Tree) Size() int {
	return t.count
}
