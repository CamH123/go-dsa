package redblack

type color uint8

const (
	black color = iota
	red
)

type node[V any] struct {
	key    int
	value  V
	color  color
	parent *node[V]
	left   *node[V]
	right  *node[V]
}

type RBTree[V any] struct {
	root     *node[V]
	sentinel *node[V]
}

func New[V any]() *RBTree[V] {
	sentinel := &node[V]{color: black}

	sentinel.parent = sentinel
	sentinel.left = sentinel
	sentinel.right = sentinel

	return &RBTree[V]{
		sentinel: sentinel,
		root:     sentinel,
	}
}

func (t *RBTree[V]) searchNode(target int) *node[V] {
	n := t.root
	for n != t.sentinel && n.key != target {
		if target > n.key {
			n = n.right
		} else {
			n = n.left
		}
	}

	return n
}

func (t *RBTree[V]) minimumNode(n *node[V]) *node[V] {
	for n != t.sentinel && n.left != t.sentinel {
		n = n.left
	}

	return n
}

func (t *RBTree[V]) maximumNode(n *node[V]) *node[V] {
	for n != t.sentinel && n.right != t.sentinel {
		n = n.right
	}

	return n
}

func (t *RBTree[V]) Search(target int) (V, bool) {
	n := t.searchNode(target)
	if n != t.sentinel {
		return n.value, true
	}

	var zero V
	return zero, false
}

func (t *RBTree[V]) Minimum() (int, V, bool) {
	n := t.minimumNode(t.root)
	if n == t.sentinel {
		var zero V
		return 0, zero, false
	}

	return n.key, n.value, true
}

func (t *RBTree[V]) Maximum() (int, V, bool) {
	n := t.maximumNode(t.root)
	if n == t.sentinel {
		var zero V
		return 0, zero, false
	}

	return n.key, n.value, true
}

func (t *RBTree[V]) Predecessor(target int) (int, V, color, bool) {
	// 1. Find the node
	n := t.searchNode(target)

	if n == t.sentinel {
		var zero V
		return 0, zero, black, false
	}

	// Case 1: Rightmost child in left child
	if n.left != t.sentinel {
		n = t.maximumNode(n.left)
		return n.key, n.value, n.color, true
	}

	// Case 2: first parent going left
	parent := n.parent
	for parent != t.sentinel {
		if parent.right == n {
			return parent.key, parent.value, parent.color, true
		}
		n = parent
		parent = parent.parent
	}

	// Case 3: elemet is minimum and there is no predecessor
	var zero V
	return 0, zero, black, false
}

func (t *RBTree[V]) Successor(target int) (int, V, color, bool) {
	// 1. Find the node
	n := t.searchNode(target)

	if n == t.sentinel {
		var zero V
		return 0, zero, black, false
	}

	// Case 1: Leftmost child in right child
	if n.right != t.sentinel {
		n = t.minimumNode(n.right)
		return n.key, n.value, n.color, true
	}

	// Case 2: first parent going right
	parent := n.parent
	for parent != t.sentinel {
		if parent.left == n {
			return parent.key, parent.value, parent.color, true
		}
		n = parent
		parent = parent.parent
	}

	// Case 3: elemet is minimum and there is no predecessor
	var zero V
	return 0, zero, black, false

}

func (t *RBTree[V]) Insert(key int, value V) {
	x := t.root
	y := t.sentinel
	for x != t.sentinel {
		y = x

		if key == x.key {
			x.value = value
			return
		}

		if key > x.key {
			x = x.right
		} else {
			x = x.left
		}
	}
	n := &node[V]{
		parent: y,
		color:  red,
		left:   t.sentinel,
		right:  t.sentinel,
		key:    key,
		value:  value,
	}

	if y == t.sentinel {
		t.root = n
	} else if n.key > y.key {
		y.right = n
	} else {
		y.left = n
	}

	t.insertFixup(n)
}

func (t *RBTree[V]) insertFixup(z *node[V]) {
	for z.parent.color == red {
		if z.parent == z.parent.parent.left {
			y := z.parent.parent.right
			// Case 1: both z.parent and uncle are red
			if y.color == red {
				z.parent.color = black
				y.color = black
				z.parent.parent.color = red
				z = z.parent.parent
			} else {
				if z == z.parent.right {
					z = z.parent
					t.leftRotate(z)
				}
				z.parent.color = black
				z.parent.parent.color = red
				t.rightRotate(z.parent.parent)
			}
		} else {
			y := z.parent.parent.left
			if y.color == red {
				z.parent.color = black
				y.color = black
				z.parent.parent.color = red
				z = z.parent.parent
			} else {
				if z == z.parent.left {
					z = z.parent
					t.rightRotate(z)
				}
				z.parent.color = black
				z.parent.parent.color = red
				t.leftRotate(z.parent.parent)
			}
		}
	}
	t.root.color = black
}

func (t *RBTree[V]) transplant(u *node[V], v *node[V]) {
	if u.parent == t.sentinel {
		t.root = v
	} else if u == u.parent.left {
		u.parent.left = v
	} else {
		u.parent.right = v
	}

	v.parent = u.parent

}

func (t *RBTree[V]) leftRotate(x *node[V]) {
	y := x.right
	x.right = y.left
	if y.left != t.sentinel {
		y.left.parent = x
	}
	y.left = x

	if x.parent == t.sentinel {
		t.root = y
	} else if x == x.parent.right {
		x.parent.right = y
	} else {
		x.parent.left = y
	}
	y.parent = x.parent
	x.parent = y
}

func (t *RBTree[V]) rightRotate(y *node[V]) {
	x := y.left
	y.left = x.right
	if x.right != t.sentinel {
		x.right.parent = y
	}
	x.right = y

	if y.parent == t.sentinel {
		t.root = x
	} else if y.parent.left == y {
		y.parent.left = x
	} else {
		y.parent.right = x
	}
	x.parent = y.parent
	y.parent = x
}

func (t *RBTree[V]) Delete(key int) bool {
	// 1. find node

	z := t.root
	for z != t.sentinel && z.key != key {
		if key > z.key {
			z = z.right
		} else {
			z = z.left
		}
	}

	if z == t.sentinel {
		return false
	}

	y := z
	originalColor := z.color
	x := z.right
	if z.left == t.sentinel {
		x = z.right
		t.transplant(z, z.right)
	} else if z.right == t.sentinel {
		x = z.left
		t.transplant(z, z.left)
	} else {
		y = t.minimumNode(z.right)
		originalColor = y.color
		x = y.right
		if y != z.right {
			t.transplant(y, y.right)
			y.right = z.right
			y.right.parent = y
		} else {
			x.parent = y
		}
		t.transplant(z, y)
		y.left = z.left
		y.left.parent = y
		y.color = z.color

	}
	if originalColor == black {
		t.deleteFixup(x)
	}

	return true

}

func (t *RBTree[V]) deleteFixup(x *node[V]) {
	for x != t.root && x.color == black {
		if x == x.parent.left {
			w := x.parent.right

			if w.color == red {
				w.color = black
				x.parent.color = red
				t.leftRotate(x.parent)
				w = x.parent.right
			}
			if w.left.color == black && w.right.color == black {
				w.color = red
				x = x.parent
			} else {
				if w.right.color == black {
					w.left.color = black
					w.color = red
					t.rightRotate(w)
					w = x.parent.right
				}
				w.color = x.parent.color
				x.parent.color = black
				w.right.color = black
				t.leftRotate(x.parent)
				x = t.root
			}
		} else {
			w := x.parent.left
			if w.color == red {
				w.color = black
				x.parent.color = red
				t.rightRotate(x.parent)
				w = x.parent.left
			}
			if w.right.color == black && w.left.color == black {
				w.color = red
				x = x.parent
			} else {
				if w.left.color == black {
					w.right.color = black
					w.color = red
					t.leftRotate(w)
					w = x.parent.left
				}
				w.color = x.parent.color
				x.parent.color = black
				w.left.color = black
				t.rightRotate(x.parent)
				x = t.root
			}
		}
	}
	x.color = black
}
