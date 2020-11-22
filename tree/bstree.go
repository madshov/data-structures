package tree

import (
	"fmt"
)

// BSTree defines a tree structure with a root node and node count.
type BSTree struct {
	root  *BSNode
	count int
}

// BSNode defines a node of the binary search tree.
type BSNode struct {
	parent *BSNode
	left   *BSNode
	right  *BSNode
	Value  int
}

// NewBSTree creates a new instance of a BSTree - a binary search tree. The
// available set operations includes insert, delete, various search operations
// as well as tree traversal operations. A binary search tree is defined by a
// set nodes that satifies the binary-search-tree propery:
//     Let n be a node in a binary search tree. If m is a node in the left
//     subtree of n, then m.value <= n.value. If node m is in the right subtree
//     of n, then m.value > n.value.
// A node consists of a pointer to a parent node, along with two pointers to the
// node left child and right child. Only the root node will have a nil parent
// node, whereas any node can have either a nil left child or a nil right child.
// Leaf nodes will have both. Finally, each node holds a integer value (or key).
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
func NewBSTree() *BSTree {
	return &BSTree{}
}

// Insert adds a node n to the correct position in the tree, by following the
// binary-search-tree property. The tree is traversed down, until the parent of
// n is found. A simple value comparison between n's value and the parent is
// used to determine if n will be the left or right child of the parent.
func (t *BSTree) Insert(val int) {
	var m, n, r *BSNode

	r = t.root
	// traverse down the tree to find the correct parent
	for r != nil {
		m = r
		if val < r.Value {
			r = r.left
		} else {
			r = r.right
		}
	}

	// create node n with m as parent
	n = &BSNode{
		parent: m,
		left:   nil,
		right:  nil,
		Value:  val,
	}

	if m == nil {
		t.root = n
	} else {
		if n.Value < m.Value {
			m.left = n
		} else {
			m.right = n
		}
	}

	t.count++
}

// Delete removes a given node n from the tree. The overall strategy considers
// four different cases: n has no left child, n has just one child, n's
// successor m is the right child of n, or m lies within n right subtree, but is
// not n's right child.
// Case 1: If n has no left child, n is replaced by its right child, which can
// be nil.
// Case 2: If n has just one child, which is the left child, n is replaced by
// its left child.
// Case 3: If m is the right child of m, n is replaced by m.
// Case 4: If m is not case 3, but instead lies within n's right subtree, m is
// replaced by it own right child, then n is replaced by m.
func (t *BSTree) Delete(n *BSNode) {
	if n == nil {
		return
	}

	var m *BSNode

	// case 1, n has no left child
	if n.left == nil {
		m = n.right
		// replace n with n's right child
		if n.parent == nil {
			t.root = m
		} else {
			// if n is the left child, update the left child of n's parent,
			// else update right child of n's parent
			if n == n.parent.left {
				n.parent.left = m
			} else {
				n.parent.right = m
			}
		}

		// update the parent of m to n's parent
		if m != nil {
			m.parent = n.parent
		}

		return
	}

	// case 2, n has one child
	if n.right == nil {
		m := n.left
		// replace n with n's left child
		if n.parent == nil {
			t.root = m
		} else {
			// if n is the left child, update the left child of n's parent,
			// else update right child of n's parent
			if n == n.parent.left {
				n.parent.left = m
			} else {
				n.parent.right = m
			}
		}

		// update the parent of m to n's parent
		if m != nil {
			m.parent = n.parent
		}

		return
	}

	// n's successor will be the min node of the right subtree
	m = t.Min(n.right)

	// case 3, n's successor is (right) child of n
	if m.parent != n {
		o := m.right
		// replace m with m's right child
		if n.parent == nil {
			t.root = o
		} else {
			// if n is the left child, update the left child of n's parent,
			// else update right child of n's parent
			if n == n.parent.left {
				n.parent.left = o
			} else {
				n.parent.right = o
			}
		}

		// update the parent of o to n's parent
		if o != nil {
			o.parent = n.parent
		}

		o = n.right
		o.parent = m
	}

	// case 4, n's successor is within n's (right) subtree
	// replace n with m
	if n.parent == nil {
		t.root = m
	} else {
		// if n is the left child, update the left child of n's parent,
		// else update right child of n's parent
		if n == n.parent.left {
			n.parent.left = m
		} else {
			n.parent.right = m
		}
	}

	// update the parent of m to n's parent
	if m != nil {
		m.parent = n.parent
	}

	m.left = n.left
	m.left.parent = m

	t.count--
}

// Search begins its search from the given node n, and traverses downwards
// through the tree. For each node it encounters, it compares the search value
// with n's value. If they equal, the search terminates. If the search value is
// less, the search continues along n's left subtree.
// Symmetrically, if the search value is greater than n's value it continues
// along n's right subtree. The binary-search-tree property ensures that this
// is the correct, and that if the search value exists, it will be found.
// Search is recursive and uses a helper function to maintain the recursive
// callstack on either the left or right subtree.
func (t *BSTree) Search(n *BSNode, val int) (*BSNode, bool) {
	if n == nil {
		return nil, false
	}

	if val == n.Value {
		return n, true
	}

	if val < n.Value {
		return t.Search(n.left, val)
	}

	return t.Search(n.right, val)
}

// SearchIt is similar to Search, but uses an iterative approach, that utilizes
// a for loop to traverse down the tree.
func (t *BSTree) SearchIt(n *BSNode, val int) (*BSNode, bool) {
	for n != nil && val != n.Value {
		if val < n.Value {
			n = n.left
		} else {
			n = n.right
		}
	}

	return n, n != nil
}

// Min finds the node with the minimum value of a given node n's subtree, i.e.
// the leftmost node in the subtree. Min is recursive, and calls itself on the
// left subtree.
func (t *BSTree) Min(n *BSNode) *BSNode {
	if n != nil && n.left != nil {
		return t.Min(n.left)
	}

	return n
}

// MinIt is similar to Min, but uses an iterative approach, that utilizes a for
// loop to traverse the left subtree.
func (t *BSTree) MinIt(n *BSNode) *BSNode {
	for n != nil && n.left != nil {
		n = n.left
	}

	return n
}

// Max finds the node with the maximum value of a given node n's subtree, i.e.
// the rightmost node in the substree. Max is recursive, and calls itself on the
// right subtree.
func (t *BSTree) Max(n *BSNode) *BSNode {
	if n != nil && n.right != nil {
		return t.Max(n.right)
	}

	return n
}

// MaxIt is similar to Max, but uses an iterative approach, that utilizes a for
// loop to traverse the right subtree.
func (t *BSTree) MaxIt(n *BSNode) *BSNode {
	for n != nil && n.right != nil {
		n = n.right
	}

	return n
}

// Successor finds the next node for a given node n. The function is split into
// two cases. If the right subtree of node n is not empty, then the successor
// of n is just the leftmost node in n's right substree.
// Otherwise, then either n does not have a successor or the successor is the
// lowest ancestor of n whose left child is also an ancestor of n.
func (t *BSTree) Successor(n *BSNode) (*BSNode, bool) {
	if n == nil {
		return nil, false
	}

	if n.right != nil {
		return t.Min(n.right), true
	}

	var m *BSNode
	if n.parent != nil {
		m = t.suc(n, n.parent)
	}

	return m, m != nil
}

// suc is a helper function that traverses up the tree from a given node n, and
// checks if n is the right child of its parent. When this is false, i.e. n is a
// left child, the function returns n's parent.
func (t *BSTree) suc(n, par *BSNode) *BSNode {
	if par.right != nil && n == par.right {
		if par.parent != nil {
			return t.suc(n.parent, par.parent)
		}

		return nil
	}

	return par
}

// SuccessorIt is similar to Successor, but uses an iterative approach that
// utilizes a for loop to traverse up the tree.
func (t *BSTree) SuccessorIt(n *BSNode) (*BSNode, bool) {
	if n == nil {
		return nil, false
	}

	if n.right != nil {
		return t.MinIt(n.right), true
	}

	m := n.parent
	for m != nil && n == m.right {
		n = m
		m = n.parent
	}

	return m, m != nil
}

// Predecessor finds the previous node for a given node n. The function is split
// into two cases. If the left subtree of node n is not empty, then the
// predecessor of n is just the rightmost node in n's left subtree.
// Otherwise, then either n does not have a predecessor or the predecessor is
// the lowest ancestor of n whose right child is also an ancestor of n.
func (t *BSTree) Predecessor(n *BSNode) (*BSNode, bool) {
	if n == nil {
		return nil, false
	}

	if n.left != nil {
		return t.Max(n.left), true
	}

	var m *BSNode
	if n.parent != nil {
		m = t.pre(n, n.parent)
	}

	return m, m != nil
}

// pre is a helper function that traverses up the tree from a given node n, and
// checks if n is the left child of its parent. When this is false, i.e. n is a
// right child, the function returns n's parent.
func (t *BSTree) pre(m, par *BSNode) *BSNode {
	if par.left != nil && m == par.left {
		if par.parent != nil {
			return t.pre(m.parent, par.parent)
		}

		return nil
	}

	return par
}

// PredecessorIt is similar to Predecessor, but uses an iterative approach
// that utilizes a for loop to traverse up the tree.
func (t *BSTree) PredecessorIt(n *BSNode) (*BSNode, bool) {
	if n == nil {
		return nil, false
	}

	if n.left != nil {
		return t.MaxIt(n.left), true
	}

	m := n.parent
	for m != nil && n == m.left {
		n = m
		m = n.parent
	}

	return m, m != nil
}

// InOrder prints an inordered list of the values of node n's subtree. This
// results in an ascending ordered list of the values. InOrder is recursive, and
// calls itself on the left and right subtree, but calling the a function given
// as input between each of them.
func (t *BSTree) InOrder(n *BSNode, f func(*BSNode)) {
	if n == nil {
		return
	}

	if n.left != nil {
		t.InOrder(n.left, f)
	}

	f(n)

	if n.right != nil {
		t.InOrder(n.right, f)
	}
}

// InOrderIt is similar to InOrder, but uses an interative approach that
// utilizes for loops and a stack to hold nodes.
func (t *BSTree) InOrderIt(f func(*BSNode)) {
	n := t.root
	ns := make([]*BSNode, t.Size())
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
func (t *BSTree) PreOrder(n *BSNode, f func(*BSNode)) {
	if n == nil {
		return
	}

	f(n)

	if n.left != nil {
		t.PreOrder(n.left, f)
	}

	if n.right != nil {
		t.PreOrder(n.right, f)
	}
}

// PreOrderIt is similar to PreOrder, but uses an interative approach that
// utilizes for loops and a stack to hold nodes.
func (t *BSTree) PreOrderIt(f func(*BSNode)) {
	n := t.root
	ns := make([]*BSNode, t.Size())
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
func (t *BSTree) PostOrder(n *BSNode, f func(*BSNode)) {
	if n == nil {
		return
	}

	if n.left != nil {
		t.PostOrder(n.left, f)
	}

	if n.right != nil {
		t.PostOrder(n.right, f)
	}

	f(n)
}

// PostOrderIt is similar to PostOrder, but uses an interative approach that
// utilizes for loops and a stack to hold nodes.
func (t *BSTree) PostOrderIt(f func(*BSNode)) {
	n := t.root
	ns := make([]*BSNode, t.Size())
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
			if n.right != nil && ns[top-1].Value == n.right.Value {
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

// Root returns the root of the tree.
func (t *BSTree) Root() *BSNode {
	return t.root
}

// Size returns the total number of the nodes in the tree.
func (t *BSTree) Size() int {
	return t.count
}

// Print displays a visual 2d representation of the tree.
func (t *BSTree) Print(n *BSNode, level int) {
	if n == nil {
		return
	}

	t.Print(n.right, level+1)

	if level != 0 {
		for i := 0; i < level-1; i++ {
			fmt.Print("|    ")
		}
		fmt.Printf("|----%d\n", n.Value)
	} else {
		fmt.Printf("%d\n", n.Value)
	}

	t.Print(n.left, level+1)
}

// GetHeight returns the maxumum height, i.e. the number of edges on the longest
// path of the binay tree. The root node has height 0.
func (t *BSTree) GetHeight(n *BSNode) int {
	if n == nil {
		return 0
	}

	leftHeight, rightHeight := 0, 0

	if n.left != nil {
		leftHeight = 1 + t.GetHeight(n.left)
	}

	if n.right != nil {
		rightHeight = 1 + t.GetHeight(n.right)
	}

	if leftHeight > rightHeight {
		return leftHeight
	}

	return rightHeight
}

// LeftNodeCount returns the number of left nodes in the binary tree.
func (t *BSTree) LeftNodeCount(n *BSNode) int {
	if n == nil {
		return 0
	}

	count := 0
	if n.left != nil {
		count += 1 + t.LeftNodeCount(n.left)
	}

	if n.right != nil {
		count += t.LeftNodeCount(n.right)
	}

	return count
}

// RightNodeCount returns the number of right nodes in the binary tree.
func (t *BSTree) RightNodeCount(n *BSNode) int {
	if n == nil {
		return 0
	}

	count := 0
	if n.right != nil {
		count += 1 + t.RightNodeCount(n.right)
	}

	if n.left != nil {
		count += t.RightNodeCount(n.left)
	}

	return count
}

// FullNodeCount returns the number of full node, i.e. nodes with two children,
// in the binary tree.
func (t *BSTree) FullNodeCount(n *BSNode) int {
	if n == nil {
		return 0
	}

	count := 0
	if n.left != nil && n.right != nil {
		count++
	}

	count += t.FullNodeCount(n.left) + t.FullNodeCount(n.right)

	return count
}

// HalfNodeCount returns the number of half nodes, i.e. nodes with one child, in
// the binary tree.
func (t *BSTree) HalfNodeCount(n *BSNode) int {
	if n == nil {
		return 0
	}

	count := 0
	if (n.left == nil && n.right != nil) ||
		(n.left != nil && n.right == nil) {
		count++
	}

	count += t.HalfNodeCount(n.left) + t.HalfNodeCount(n.right)

	return count
}

// NonLeafCount returns the number of non-leaf nodes, i.e. nodes with at least
// one child, in the binary tree.
func (t *BSTree) NonLeafCount(n *BSNode) int {
	if n == nil {
		return 0
	}

	count := 0
	if n.left != nil || n.right != nil {
		count++
	}

	count += t.NonLeafCount(n.left) + t.NonLeafCount(n.right)

	return count
}
