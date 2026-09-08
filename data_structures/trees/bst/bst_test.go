package bst

import "testing"

func buildTestTree() *BST[string] {
	root := &node[string]{key: 10, value: "ten"}
	left := &node[string]{key: 5, value: "five"}
	right := &node[string]{key: 15, value: "fifteen"}
	leftLeft := &node[string]{key: 2, value: "two"}
	leftRight := &node[string]{key: 7, value: "seven"}
	rightLeft := &node[string]{key: 12, value: "twelve"}
	rightRight := &node[string]{key: 20, value: "twenty"}

	root.left = left
	root.right = right

	left.parent = root
	right.parent = root

	left.left = leftLeft
	left.right = leftRight
	right.left = rightLeft
	right.right = rightRight

	leftLeft.parent = left
	leftRight.parent = left
	rightLeft.parent = right
	rightRight.parent = right

	tree := BST[string]{root: root}
	return &tree
}

func TestSearch(t *testing.T) {

	tree := buildTestTree()

	tests := []struct {
		name   string
		target int
		want   string
		wantOK bool
	}{
		{"search root", 10, "ten", true},
		{"search leaf", 2, "two", true},
		{"search internal node", 15, "fifteen", true},
		{"search missing node", 99, "", false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, ok := tree.Search(test.target)

			if got != test.want || ok != test.wantOK {
				t.Errorf(
					"Search(%d) = (%q, %v), want (%q, %v)",
					test.target,
					got,
					ok,
					test.want,
					test.wantOK,
				)
			}
		})
	}
}

func TestMinimum(t *testing.T) {
	tree := buildTestTree()

	key, value, ok := tree.Minimum()

	if key != 2 || value != "two" || !ok {
		t.Errorf(
			"Minimum() = (%d, %q, %v), want (%d, %q, %v)",
			key,
			value,
			ok,
			2,
			"two",
			true,
		)
	}
}

func TestMaximum(t *testing.T) {
	tree := buildTestTree()

	key, value, ok := tree.Maximum()

	if key != 20 || value != "twenty" || !ok {
		t.Errorf(
			"Maximum() = (%d, %q, %v), want (%d, %q, %v)",
			key,
			value,
			ok,
			20,
			"twenty",
			true,
		)
	}
}

func TestPredecessor(t *testing.T) {
	tree := buildTestTree()

	tests := []struct {
		name    string
		target  int
		wantKey int
		wantVal string
		wantOK  bool
	}{
		{"root", 10, 7, "seven", true},
		{"leaf", 12, 10, "ten", true},
		{"internal", 15, 12, "twelve", true},
		{"minimum", 2, 0, "", false},
		{"maximum", 20, 15, "fifteen", true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			key, value, ok := tree.Predecessor(test.target)

			if key != test.wantKey || value != test.wantVal || ok != test.wantOK {
				t.Errorf(
					"Predecessor(%d) = (%d, %q, %v), want (%d, %q, %v)",
					test.target,
					key,
					value,
					ok,
					test.wantKey,
					test.wantVal,
					test.wantOK,
				)
			}
		})
	}
}

func TestSuccessor(t *testing.T) {
	tree := buildTestTree()

	tests := []struct {
		name    string
		target  int
		wantKey int
		wantVal string
		wantOK  bool
	}{
		{"root", 10, 12, "twelve", true},
		{"leaf", 7, 10, "ten", true},
		{"internal", 15, 20, "twenty", true},
		{"minimum", 2, 5, "five", true},
		{"maximum", 20, 0, "", false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			key, value, ok := tree.Successor(test.target)

			if key != test.wantKey || value != test.wantVal || ok != test.wantOK {
				t.Errorf(
					"Successor(%d) = (%d, %q, %v), want (%d, %q, %v)",
					test.target,
					key,
					value,
					ok,
					test.wantKey,
					test.wantVal,
					test.wantOK,
				)
			}
		})
	}
}

func TestInsert(t *testing.T) {
	t.Run("insert root", func(t *testing.T) {
		tree := &BST[string]{}

		tree.Insert(10, "ten")

		if tree.root == nil {
			t.Fatal("expected root to be created, got nil")
		}

		if tree.root.key != 10 || tree.root.value != "ten" {
			t.Errorf(
				"root = (%d, %q), want (%d, %q)",
				tree.root.key,
				tree.root.value,
				10,
				"ten",
			)
		}

		if tree.root.parent != nil {
			t.Errorf("root parent = %v, want nil", tree.root.parent)
		}
	})

	t.Run("insert new element", func(t *testing.T) {
		tree := buildTestTree()

		tree.Insert(6, "six")

		n := tree.root.search(6)

		if n == nil {
			t.Fatal("expected key 6 to be inserted")
		}

		if n.value != "six" {
			t.Errorf("value = %q, want %q", n.value, "six")
		}

		if n.parent == nil || n.parent.key != 7 {
			t.Errorf("parent key is incorrect")
		}
	})

	t.Run("insert existing element", func(t *testing.T) {
		tree := buildTestTree()

		tree.Insert(7, "updated seven")

		n := tree.root.search(7)

		if n == nil {
			t.Fatal("expected key 7 to exist")
		}

		if n.value != "updated seven" {
			t.Errorf(
				"value = %q, want %q",
				n.value,
				"updated seven",
			)
		}
	})
}

func TestDelete(t *testing.T) {
	tests := []struct {
		name   string
		key    int
		wantOK bool
		check  func(t *testing.T, tree *BST[string])
	}{
		{
			name:   "delete root",
			key:    10,
			wantOK: true,
			check: func(t *testing.T, tree *BST[string]) {
				if tree.root.key != 12 {
					t.Errorf("root key = %d, want 12", tree.root.key)
				}

				if tree.root.parent != nil {
					t.Errorf("root parent = %v, want nil", tree.root.parent)
				}

				if tree.root.left.key != 5 {
					t.Errorf("root left key = %d, want 5", tree.root.left.key)
				}

				if tree.root.right.key != 15 {
					t.Errorf("root right key = %d, want 15", tree.root.right.key)
				}
			},
		},
		{
			name:   "delete leaf",
			key:    2,
			wantOK: true,
			check: func(t *testing.T, tree *BST[string]) {
				n := tree.root.search(5)

				if n.left != nil {
					t.Errorf("node 5 left = %v, want nil", n.left)
				}
			},
		},
		{
			name:   "delete middle node",
			key:    15,
			wantOK: true,
			check: func(t *testing.T, tree *BST[string]) {
				n := tree.root.search(20)

				if n.parent != tree.root {
					t.Errorf("node 20 parent is incorrect")
				}

				if n.left == nil || n.left.key != 12 {
					t.Errorf("node 20 left child is incorrect")
				}

				if n.left.parent != n {
					t.Errorf("node 12 parent is incorrect")
				}
			},
		},
		{
			name:   "delete nonexistent",
			key:    99,
			wantOK: false,
			check: func(t *testing.T, tree *BST[string]) {
				if tree.root.key != 10 {
					t.Errorf("root key = %d, want 10", tree.root.key)
				}
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tree := buildTestTree()

			ok := tree.Delete(test.key)

			if ok != test.wantOK {
				t.Errorf(
					"Delete(%d) = %v, want %v",
					test.key,
					ok,
					test.wantOK,
				)
			}

			// Deleted key should no longer exist.
			if test.wantOK && tree.root.search(test.key) != nil {
				t.Errorf("expected key %d to be deleted", test.key)
			}

			test.check(t, tree)
		})
	}
}
