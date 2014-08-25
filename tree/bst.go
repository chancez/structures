package tree

import (
	"fmt"
)

type Tree struct {
	root *Node
}

type Node struct {
	value               Orderable
	left, right, parent *Node
}

func (t *Tree) Root() BstNode {
	return t.root
}

func (n *Node) Value() Orderable {
	return n.value
}

func (n *Node) Left() BstNode {
	return n.left
}

func (n *Node) Right() BstNode {
	return n.right
}

func (n *Node) Parent() BstNode {
	return n.parent
}

func InOrderWalk(n *Node) {
	if n == nil {
		return
	}
	InOrderWalk(n.left)
	fmt.Println(n.value)
	InOrderWalk(n.right)
}

func (n *Node) Search(v Orderable) *Node {
	if Eq(v, n.value) {
		return n
	}
	if n.left != nil && Lt(v, n.value) {
		return n.left.Search(v)
	}
	if n.right != nil {
		return n.right.Search(v)
	}
	return nil
}

func (n *Node) SearchIter(v Orderable) *Node {
	curr := n
	for curr != nil && !Eq(v, curr.value) {
		if Lt(v, curr.value) {
			curr = curr.left
		} else {
			curr = curr.right
		}
	}
	return curr
}

func (n *Node) Min() *Node {
	for n.left != nil {
		n = n.left
	}
	return n
}

func (n *Node) Max() *Node {
	for n.right != nil {
		n = n.right
	}
	return n
}

func (n *Node) Succ() *Node {
	if n.right != nil {
		return n.right.Min()
	}
	succ := n.parent
	for succ != nil && n == succ.right {
		n = succ
		succ = succ.parent
	}
	return succ
}

func (n *Node) Pred() *Node {
	if n.left != nil {
		return n.left.Max()
	}
	pred := n.parent
	for pred != nil && n == pred.left {
		n = pred
		pred = pred.parent
	}
	return pred
}

func (t *Tree) Insert(v *Node) {
	var tmp *Node
	curr := t.root
	for curr != nil {
		tmp = curr
		if Lt(v.value, curr.value) {
			curr = curr.left
		} else {
			curr = curr.right
		}
	}
	v.parent = tmp
	if tmp == nil {
		t.root = v // tree was empty
	} else if Lt(v.value, tmp.value) {
		tmp.left = v
	} else {
		tmp.right = v
	}
}

func (t *Tree) Delete(n *Node) {
	if n.left == nil {
		t.Transplant(n, n.right)
	} else if n.right == nil {
		t.Transplant(n, n.left)
	} else {
		min := n.right.Min()
		if min.parent != n {
			t.Transplant(min, min.right)
			min.right = n.right
			min.right.parent = min
		}
		t.Transplant(n, min)
		n.left = min.left
		n.left.parent = min
	}
}

// Transplant replaces one subtree as a child of its parent with another subtree
func (t *Tree) Transplant(original, replacement *Node) {
	if original.parent == nil { // Root of the tree
		t.root = replacement
	} else if original == original.parent.left { // Left child
		original.parent.left = replacement
	} else { // Right child
		original.parent.right = replacement
	}
	// Update the parent node for the new subtree
	if replacement != nil {
		replacement.parent = original.parent
	}
}

func (t *Tree) Search(v Orderable) *Node {
	if t.root != nil {
		return t.root.Search(v)
	}
	return nil
}

func (t *Tree) Min() *Node {
	if t.root != nil {
		return t.root.Min()
	}
	return nil
}

func (t *Tree) Max() *Node {
	if t.root != nil {
		return t.root.Max()
	}
	return nil
}
