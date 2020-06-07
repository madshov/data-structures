package tree

import "fmt"

// Color represents the color of a node.
type Color int

const (
	red Color = iota
	black
)

// RBTree defines a tree structure with a root node and node count.
type RBTree struct {
	root  *RBNode
	count int
}

// RBNode defines a node of the binary search tree.
type RBNode struct {
	parent *RBNode
	left   *RBNode
	right  *RBNode
	color  Color
	Value  int
}

// NewRBTree creates a new instance of a RBTree - red-black tree. A red-black
// tree is a binary search tree, but where each node has an extra bit of
// information - a color, either red or black. Given some contraints to the
// color of each node, a red-black tree ensures that no path in the tree is more
// than twice as long as any other path in the tree, and therefore the tree is
// considered balanced. The following properties must be satified at all times
// in a red-black tree:
//     1) Every node is either red or black.
//     2) The root must always be black.
//     3) Every leaf (nil node) is black.
//     4) If a node is red, both children must be black.
//     5) For each node, all paths to its decendent leaves must contain the same
//        number of black nodes.
// Example of an instance of a red black tree with 11 nodes (r indicates a red
// node, b a black node):
//
//                      15b
//                   /       \
//                  6r        18b
//                /   \      /   \
//               3b   9b   17r    20r
//              / \   / \
//             2r  4r7r 13r
//
// A red-black tree with n nodes has height at most 2lg(n+1), which means that
// all search operations (search, min, max, predescessor, successor) can be
// implemented in O(lg n) time. Further, insertion and deletions take O(lg n)
// time, making it an efficient data-structure.
func NewRBTree() *RBTree {
	return &RBTree{}
}

// Insert adds a node n to the correct position in the tree, by following the
// binary-search-tree property. The tree is traversed down, until the parent of
// n is found. A simple value comparison between n's value and the parent is
// used to determine if n should be the left or right child of the parent. The
// inserted node will always have color red. Finally insertRecolor is called to
// fix colors up the tree to maintain the red-black properties.
func (t *RBTree) Insert(val int) {
	var m, n, r *RBNode

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

	// create node n with m as parent and color red
	n = &RBNode{
		parent: m,
		left:   nil,
		right:  nil,
		color:  red,
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

	t.insertRecolor(n)
	t.count++
}

// Delete removes a given node n from the tree.
func (t *RBTree) Delete(n *RBNode) {
	if n == nil {
		return
	}

	var (
		m, o *RBNode
		c    Color
	)

	m = n
	c = m.color

	if n.left == nil {
		o = n.right
		// replace n with n's right child
		if n.parent == nil {
			t.root = n.right
		} else {
			// if n is a left child
			if n == n.parent.left {
				n.parent.left = n.right
			} else {
				// n is a right child
				n.parent.right = n.right
			}
		}

		if n.right != nil {
			n.right.parent = n.parent
		}
	} else {
		if n.right == nil {
			o = n.left
			// replace n with n's left child
			if n.parent == nil {
				t.root = n.left
			} else {
				// if n is a left child
				if n == n.parent.left {
					n.parent.left = n.left
				} else {
					// n is a right child
					n.parent.right = n.left
				}
			}

			if n.left != nil {
				n.left.parent = n.parent
			}
		} else {
			m := t.Min(n.right)
			c = m.color
			o = m.right
			if o != nil && m.parent == n {
				o.parent = m
			} else {
				// replace m with m's right child
				if m.parent == nil {
					t.root = m.right
				} else {
					// if m is a left child
					if m == m.parent.left {
						m.parent.left = m.right
					} else {
						// m is a right child
						m.parent.right = m.right
					}
				}

				if m.right != nil {
					m.right.parent = m.parent
					m.right = n.right
					m.right.parent = m
				}
			}
			// replace n with m
			if n.parent == nil {
				t.root = m
			} else {
				// if n is a left child
				if n == n.parent.left {
					n.parent.left = m
				} else {
					// n is a right child
					n.parent.right = m
				}
			}

			m.parent = n.parent

			m.left = n.left
			m.left.parent = m
			m.color = n.color
		}
	}

	if c == black {
		t.deleteRecolor(o)
	}

	t.count--
}

// Search begins its search from the given node n, and traverses downwards
// through the tree. For each node it encounters, it compares the search value
// with n's value. If they equal, the search terminates. If the search value is
// less, the search continues along n's left subtree.
// Symmetrically, if the search value is greater than n's value it continues
// along n's right subtree. The binary-search-tree property ensures that this
// is the correct, and that if the search value exists, it will be found.
func (t *RBTree) Search(n *RBNode, val int) (*RBNode, bool) {
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
// the leftmost node in the subtree.
func (t *RBTree) Min(n *RBNode) *RBNode {
	if n == nil {
		n = t.root
	}

	for n.left != nil {
		n = n.left
	}

	return n
}

// Max finds the node with the maximum value of a given node n's subtree, i.e.
// the rightmost node in the substree.
func (t *RBTree) Max(n *RBNode) *RBNode {
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
// Otherwise, then either n does not have a successor or the successor is the
// lowest ancestor of n whose left child is also an ancestor of n.
func (t *RBTree) Successor(n *RBNode) (*RBNode, bool) {
	if n == nil {
		n = t.root
	}

	if n.right != nil {
		return t.Min(n.right), true
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
func (t *RBTree) Predecessor(n *RBNode) (*RBNode, bool) {
	if n == nil {
		n = t.root
	}

	if n.left != nil {
		return t.Max(n.left), true
	}

	m := n.parent
	for m != nil && n == m.left {
		n = m
		m = n.parent
	}

	return m, m != nil
}

// InOrder prints an inordered list of the values of node n's subtree. This
// results in an ascending ordered list of the values.
func (t *RBTree) InOrder(n *RBNode, f func(*RBNode)) {
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

func (t *RBTree) insertRecolor(n *RBNode) {
	for n.parent != nil && n.parent.color == red {
		if n.parent == n.parent.parent.left {
			m := n.parent.parent.right
			if m != nil && m.color == red {
				n.parent.color = black
				m.color = black
				n.parent.parent.color = red
				n = n.parent.parent
			} else {
				// if n is a right child
				if n == n.parent.right {
					n = n.parent
					t.rotateLeft(n)
				}

				n.parent.color = black
				n.parent.parent.color = red
				t.rotateRight(n.parent.parent)
			}
		} else {
			m := n.parent.parent.left
			if m != nil && m.color == red {
				n.parent.color = black
				m.color = black
				n.parent.parent.color = red
				n = n.parent.parent
			} else {
				// if n is a left child
				if n == n.parent.left {
					n = n.parent
					t.rotateRight(n)
				}

				n.parent.color = black
				n.parent.parent.color = red
				t.rotateLeft(n.parent.parent)
			}
		}
	}

	t.root.color = black
}

func (t *RBTree) deleteRecolor(n *RBNode) {
	for n != t.root && n.color == black {
		// if n is a left child
		if n == n.parent.left {
			m := n.parent.right
			if m.color == red {
				m.color = black
				n.parent.color = red
				t.rotateLeft(n.parent)
				m = n.parent.right
			}

			if m.left.color == black && m.right.color == black {
				m.color = red
				n = n.parent
			} else {
				if m.right.color == black {
					m.left.color = black
					m.color = red
					t.rotateRight(m)
					m = n.parent.right
				}

				m.color = n.parent.color
				n.parent.color = black
				m.right.color = black
				t.rotateLeft(n.parent)
				n = t.root
			}
		} else {
			// n is a right child
			m := n.parent.left
			if m.color == red {
				m.color = black
				n.parent.color = red
				t.rotateRight(n.parent)
				m = n.parent.left
			}

			if m.left.color == black && m.right.color == black {
				m.color = red
				n = n.parent
			} else {
				if m.left.color == black {
					m.right.color = black
					m.color = red
					t.rotateLeft(m)
					m = n.parent.left
				}

				m.color = n.parent.color
				n.parent.color = black
				m.left.color = black
				t.rotateRight(n.parent)
				n = t.root
			}
		}
	}

	n.color = black
}

// rotateLeft changes the tree structure but maintains the binary-search-tree
// property. This function rotates a node n and it's right child m in a counter
// clockwise fashion:
//            n                m
//           / \              / \
//          a   m    --->    n   c
//             / \          / \
//            b   c        a   b
// Rotating left is sometimes necessary for insertions and deletions, in order
// to maintain the red-black properties of the tree.
func (t *RBTree) rotateLeft(n *RBNode) {
	m := n.right
	n.right = m.left
	if m.left != nil {
		m.left.parent = n
	}

	m.parent = n.parent
	// if n is the root
	if n.parent == nil {
		t.root = m
	} else {
		// if n is a left child
		if n == n.parent.left {
			n.parent.left = m
		} else {
			// n is a right child
			n.parent.right = m
		}
	}

	m.left = n
	n.parent = m
}

// rotateRight changes the tree structure but maintains the binary-search-tree
// property. This function rotates a node n and it's left child m in a clockwise
// fashion:
//            n                m
//           / \              / \
//          m   c    --->    a   n
//         / \                  / \
//        a   b                b   c
// Rotating right is sometimes necessary for insertions and deletions, in order
// to maintain the red-black properties of the tree.
func (t *RBTree) rotateRight(n *RBNode) {
	m := n.left
	n.left = m.right
	if m.right != nil {
		m.right.parent = n
	}

	m.parent = n.parent
	// if n is the root
	if n.parent == nil {
		t.root = m
	} else {
		// if n is a left child
		if n == n.parent.left {
			n.parent.left = m
		} else {
			// n is a right child
			n.parent.right = m
		}
	}

	m.right = n
	n.parent = m
}

// Root returns the root of the tree.
func (t *RBTree) Root() *RBNode {
	return t.root
}

// Size returns the total number of the nodes in the tree.
func (t *RBTree) Size() int {
	return t.count
}

// Print displays a visual 2d representation of the tree.
func (t *RBTree) Print(n *RBNode, level int) {
	if n == nil {
		return
	}

	t.Print(n.right, level+1)

	if level != 0 {
		for i := 0; i < level-1; i++ {
			fmt.Print("|    ")
		}
		fmt.Printf("|----%d-%d\n", n.Value, n.color)
	} else {
		fmt.Printf("%d-%d\n", n.Value, n.color)
	}

	t.Print(n.left, level+1)
}
