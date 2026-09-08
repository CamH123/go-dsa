# Binary Search Tree
Referencing CLRS (ch. 12)

## Supported Functions
Search, Minimum, Maximum, Predecessor, Successor, Insert, Delete
- Can thus be used as a dictionary and priority queue

## 12.1 What is a binary search tree?
- Each node contains a key, ptr to satellite data, left, right, and parent
- The tree itself has a root attribute
- The height of a BST is the number of nodes on the longest path from root r to a leaf
- Keys are stored to satisfy the binary search tree property (invariance): for node x, all nodes y in its left subtree follow y.key <= x.key and all nodes z in its right subtree follow z.key >= x.key
- This property allows you to print out all nodes in sorted order through an inorder tree walk - proof by induction (proof.1)
- Inorder traversal takes O(n) time since after the initial call, the procedure calls itself recursively exactly twice for each node in the tree, once for each child. - proof by substitution (proof.2)

## 12.2 Querying a binary search tree
- Tree search: O(h) recursively searches either left or right depending on target key value
- Minimum: find leftmost child
- Maximum: find rightmost child
- Successor: leftmost node in right subtree or lowest ancestor of x whose left child is also an ancestor of x

## 12.3 Insertion and deletion
- Insertion follows the rules until it reaches an opening to insert the value as a leaf node. O(h)
- Deletion: 1. No children = modify parent to replace z with nil 2. if z has one child, elevate child to replace z's position 3. if z has children, replace with its successor
- Transplant(u, v): replace one subtree as a child of its parent with another subtree; doesn't update v's pointers

## Proofs
1. Prove an inorder tree walk of a BST prints out all keys in sorted order.

We prove by strong induction on n (the number of nodes in a BST) that an inorder traversal outputs the keys in sorted order.

Base Case:
Let n = 0. Since the tree is empty, the inorder traversal outputs an empty sequence, which is trivially sorted. Therefore, the base case holds.

Inductive Hypothesis:
Assume for all BSTs with fewer than n nodes, an inorder traversal outputs the keys in sorted order.

Inductive Step:
Consider a BST T with n > 0 nodes. Let x be its root. Suppose the left subtree has k nodes, and the right subtree has n - k - 1 nodes. Since k < n and n - k - 1 < n, the inductive hypothesis applies to both subtrees. Therefore, the inorder traversal outputs the keys of the left subtree in sorted order and the keys of the right subtree in sorted order. Since an inorder traversal visits the left subtree first, then the root, then the right subtree, its output on the tree T will be in the form of [sorted keys in left subtree], x.key, [sorted keys in right subtree]. By the BST property of invariance, all keys in the left subtree are less than or equal to x.key, and all keys in the right subtree are >= x.key. Thus, the output of the keys in tree T is sorted, and the statement holds for a tree with n nodes.

Therefore, by strong induction, an inorder traversal outputs the keys of every binary search tree in sorted order.

2. Prove it takes Theta(n) time to walk an n-node BST.

Let T(n) be the time it takes to walk through a BST of n nodes. Since the walk must traverse every node, it performs at least a constant amount of work per node; thus, T(n) = Omega(n).

Now let's prove T(n) = O(n).

Let c be the constant amount of time it takes to run a tree walk on an empty tree: T(0) = c.

For n > 0, suppose the root is x, the left subtree contains k nodes, and the right subtree contains n - k - 1 nodes.

The time complexity of T(n) is thus upper bounded by T(k) + T(n-k-1) + d, where d is a constant.

Using substitution, suppose T(n) <= an + b.
Then T(n) <= ak + b + a(n-k-1) + b + d = an + 2b - a + d.

For the assumed bound to hold, show an + 2b - a + d <= an + b.
2b - a + d <= b.
a >= b + d.

For n = 0, we know T(0) = c <= a(0) + b; thus, c <= b.

Thus, the two constraints are:
b >= c.
a >= b + d.

These constraints are satisfied when b = c and a = c + d.
Therefore, the final bound is T(n) <= (c+d)n + c.
Since c and d are constants, T(n) = O(n).
