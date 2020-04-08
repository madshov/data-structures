package bstree

func (t *Tree) InsertIt(val int) {
	t.Lock()
	defer t.Unlock()

	r := t.root
	n := &Node{nil, nil, nil, val}
	var m *Node

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

// DeleteIt removes a node n with a given value in the tree. This function is
// split into 3 cases, n has no children, n has one child or n has two children.
// If n has no children, n will just be deleted. Else if n has one child, the
// child is copied to n, and the child is deleted. Else n must have two
// children, and n must have a successor which is copied to n and then deleted.
func (t *Tree) DeleteIt(val int) {
	t.Lock()
	defer t.Unlock()
}

// InsertIt begins its search from node n, and traces a simple path downwards in
// the tree. For each node it encounters, it compares the search value with n's
// value. If they equal, the search terminates. If the search value is less
// than, the search continues along n's left subtree. Symmetrically, if the
// search value is greater than n's value it continues along n's right subtree.
// The binary-search-tree property ensures that this is the correct, and that
// if the search value exists, it will be found.
func (t *Tree) SearchIt(n *Node, val int) *Node {
	t.Lock()
	defer t.Unlock()

	if n == nil {
		n = t.root
	}

	for n != nil && val != n.value {
		if val < n.value {
			n = n.left
		} else {
			n = n.right
		}
	}

	return n
}

// MinIt finds the node with the minimum value of a given node n's subtree, i.e.
// the leftmost node in the subtree.
func (t *Tree) MinIt(n *Node) *Node {
	t.Lock()
	defer t.Unlock()

	if n == nil {
		n = t.root
	}

	for n.left != nil {
		n = n.left
	}

	return n
}

// MaxIt finds the node with the maximum value of a given node n's subtree, i.e.
// the rightmost node in the substree.
func (t *Tree) MaxIt(n *Node) *Node {
	t.Lock()
	defer t.Unlock()

	if n == nil {
		n = t.root
	}

	for n.right != nil {
		n = n.right
	}

	return n
}

// SuccessorIt finds the next node for a given node n. The function is split into
// two cases. If the right subtree of node n is not empty, then the successor
// of n is just the leftmost node in n's right substree.
// If the right subtree of node n is empty, then either n does not have a
// successor or the successor is the lowest ancestor of n whose left child is
// also an ancestor of n.
func (t *Tree) SuccessorIt(n *Node) *Node {
	t.Lock()
	defer t.Unlock()

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

// PredecessorIt finds the previous node for a given node n. The function is split
// into two cases. If the left subtree of node n is not empty, then the
// predecessor of n is just the rightmost node in n's left subtree.
// If the left substree of node n is empty, then either n does not have a
// predecessor or the predecessor is the lowest ancestor of n whose right child
// is also an ancestor of n.
func (t *Tree) PredecessorIt(n *Node) *Node {
	t.Lock()
	defer t.Unlock()

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

// InOrderIt prints an inordered list of the values of node n's subtree.
// This results in an ascending ordered list of the values.
func (t *Tree) InOrderIt(f func(*Node)) {
	t.Lock()
	defer t.Unlock()

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

// PreOrderIt prints a preordered list of the values of node n's subtree.
func (t *Tree) PreOrderIt(f func(*Node)) {
	t.Lock()
	defer t.Unlock()

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

// PostOrderIt prints a postordered list of the values of node n's subtree.
func (t *Tree) PostOrderIt(f func(*Node)) {
	t.Lock()
	defer t.Unlock()

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
