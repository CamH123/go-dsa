package bst

type node[V any] struct {
	key    int
	value  V
	parent *node[V]
	left   *node[V]
	right  *node[V]
}

// Search returns node with target key or nil.
func (n *node[V]) search(target int) *node[V] {
	for n != nil && n.key != target {
		if target < n.key {
			n = n.left
		} else {
			n = n.right
		}
	}
	return n
}

// Returns node with smallest key.
func (n *node[V]) minimum() *node[V] {
	for n != nil && n.left != nil {
		n = n.left
	}
	return n
}

// Returns node with largest key.
func (n *node[V]) maximum() *node[V] {
	for n != nil && n.right != nil {
		n = n.right
	}
	return n
}

// Returns previous node from in-order traversal.
func (n *node[V]) predecessor() *node[V] {
	if n == nil {
		return nil
	}

	if n.left != nil {
		return n.left.maximum()
	}

	parent := n.parent
	for parent != nil && n == parent.left {
		n = parent
		parent = parent.parent
	}

	return parent
}

// Returns next node in in-order traversal.
func (n *node[V]) successor() *node[V] {
	if n == nil {
		return nil
	}

	if n.right != nil {
		return n.right.minimum()
	}

	parent := n.parent
	for parent != nil && n == parent.right {
		n = parent
		parent = parent.parent
	}

	return parent
}

// Inserts node with update for duplicates.
func (n *node[V]) insert(key int, value V) {
	parent := n
	for n != nil {
		if n.key == key {
			n.value = value
			return
		}

		parent = n

		if key > n.key {
			n = n.right
		} else {
			n = n.left
		}
	}

	newNode := node[V]{
		key:    key,
		value:  value,
		parent: parent,
	}

	if key > parent.key {
		parent.right = &newNode
	} else {
		parent.left = &newNode
	}
}

type BST[V any] struct {
	root *node[V]
}

// Search returns the value associated with target and reports existence in tree.
func (t *BST[V]) Search(target int) (V, bool) {
	n := t.root.search(target)
	if n == nil {
		var zero V
		return zero, false
	}

	return n.value, true
}

// Minimum returns the smallest key and associated value in a tree.
func (t *BST[V]) Minimum() (int, V, bool) {
	n := t.root.minimum()
	if n == nil {
		var zero V
		return 0, zero, false
	}

	return n.key, n.value, true
}

// Maximum returns the largest key and associated value in a tree.
func (t *BST[V]) Maximum() (int, V, bool) {
	n := t.root.maximum()
	if n == nil {
		var zero V
		return 0, zero, false
	}

	return n.key, n.value, true
}

// Predecessor returns element before target in in-order traversal.
func (t *BST[V]) Predecessor(target int) (int, V, bool) {
	n := t.root.search(target)
	if n == nil {
		var zero V
		return 0, zero, false
	}

	pred := n.predecessor()

	if pred == nil {
		var zero V
		return 0, zero, false
	}

	return pred.key, pred.value, true
}

// Successor returns element after target in in-order traversal.
func (t *BST[V]) Successor(target int) (int, V, bool) {
	n := t.root.search(target)
	if n == nil {
		var zero V
		return 0, zero, false
	}

	next := n.successor()

	if next == nil {
		var zero V
		return 0, zero, false
	}

	return next.key, next.value, true
}

// Insert add or updates tree with node.
func (t *BST[V]) Insert(key int, value V) {
	// Check root isn't nil
	if t.root == nil {
		t.root = &node[V]{
			key:   key,
			value: value,
		}
		return
	} else {
		t.root.insert(key, value)
	}
}

// Moves node v to node u
func (t *BST[V]) transplant(u *node[V], v *node[V]) {
	if u.parent == nil {
		t.root = v
	} else if u == u.parent.left {
		u.parent.left = v
	} else {
		u.parent.right = v
	}
	if v != nil {
		v.parent = u.parent
	}
}

// Delete removes a key and returns a boolean regarding success.
func (t *BST[V]) Delete(key int) bool {
	target := t.root.search(key)

	if target == nil {
		return false
	}

	if target.left == nil {
		t.transplant(target, target.right)
		return true
	}
	if target.right == nil {
		t.transplant(target, target.left)
		return true
	}
	successor := target.successor()
	if successor != target.right {
		t.transplant(successor, successor.right)
		successor.right = target.right
		successor.right.parent = successor
	}
	t.transplant(target, successor)
	successor.left = target.left
	successor.left.parent = successor
	return true
}
