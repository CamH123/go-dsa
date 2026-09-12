# Binary Search Tree

Referencing CLRS (Ch. 13)

## 13.1 Properties of Red-Black Trees

1. Red-black trees are binary search trees where each node has an extra color bit (red or black).

2. Ensures that no path is twice as long as any other --> approximately balanced --> h <= 2log(n+1)

3. Properties:

   1. Every node is red or black.
   2. The root is black.
   3. Every leaf (nil) is black.
   4. If a node is red, both children are black.
   5. For each node, all paths to descendant leaves contain the same number of black nodes.

4. Use sentinel T.nil for leaves to save space.

## 13.2 Rotations

1. Node rotations are O(1), as they only change pointers and maintain the BST property.

## 13.3 Insertion

1. First, insert the node into the tree like a regular BST. Then, color the node red. Run RB-INSERT-FIXUP.

2. Insert runs in O(log(n)) time, as the fixup may walk up to the root. It maintains the properties of the RB tree through rotations and color changes.

## 13.4 Deletion

1. Delete also runs in O(log(n)) time, first following the standard BST deletion procedure. After replacement with the successor when necessary, RB-DELETE-FIXUP performs rotations and recolorings to bring back the standard properties.

2. Delete fixup starts at the node x that replaces the node that was actually removed. If the removed node was black, it performs fixup up the tree to ensure the properties are kept.

## Lemma 13.1

A red-black tree with n internal nodes has height at most 2log(n+1).

Start by showing that the subtree rooted at any node x contains at least 2^bh(x) - 1 internal nodes through induction.

Base case: A leaf node has a bh of 0 and contains 0 internal nodes. 2^0 - 1 = 0. The base case holds.

Inductive Step: Consider a node x that has a positive height and is an internal node. This node has two children. If a child is black, then the child has a bh of bh(x) - 1. Otherwise, if the child is red, it has a bh of bh(x). Therefore, each child has a bh of at least bh(x) - 1. Using the inductive hypothesis, the number of nodes in a tree rooted at x is thus at least

2^(bh(x) - 1) - 1 + 2^(bh(x) - 1) - 1 + 1 = 2^bh(x) - 1,

proving the claim.

To complete the proof, let h be the height of the tree. By property 4, at least half the nodes on any of the paths must be black. Therefore, the black height of the root must be at least h/2, thus

n >= 2^(h/2) - 1.

Solving for h,

h <= 2 * log(n + 1).

## Problem 13-3

### a.

Prove an AVL tree of height h has at least Fh nodes with induction.

Base Case: When h = 0, the tree has 0 nodes, so 0 >= F_0 = 0. When h = 1, the tree has one node, so 1 >= F_1 = 1.

Inductive Step: Consider an AVL tree of height k. To minimize the number of nodes, its children must have heights of at least k - 1 and k - 2. By the inductive hypothesis, the two subtrees must have at least F_(k-1) and F_(k-2) nodes. Thus, the number of nodes in the AVL tree must be at least

F_(k-1) + F_(k-2) + 1 = F_k + 1 >= F_k,

proving that an AVL tree of height h has at least F_h nodes.

Since Fibonacci grows exponentially,

F_h = Theta(phi^h).

Thus,

n >= c * phi^h.

Taking the log of both sides,

log(n) >= log(c) + h * log(phi).

Thus,

h <= (log(n) - log(c)) / log(phi).

Since log(c) and log(phi) are constants, h = O(log(n)). Therefore, an AVL tree with n nodes has height O(log(n)).
