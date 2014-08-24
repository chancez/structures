package tree

import (
	"fmt"
)

type Tree struct {
	Root *Node
}

type Node struct {
	Value               int
	Left, Right, Parent *Node
}

func InOrderWalk(n *Node) {
	if n == nil {
		return
	}
	InOrderWalk(n.Left)
	fmt.Println(n.Value)
	InOrderWalk(n.Right)
}

func (n *Node) Search(v int) *Node {
	if v == n.Value {
		return n
	}
	if n.Left != nil && v < n.Value {
		return n.Left.Search(v)
	}
	if n.Right != nil {
		return n.Right.Search(v)
	}
	return nil
}

func (n *Node) SearchIter(v int) *Node {
	curr := n
	for curr != nil && v != curr.Value {
		if v < curr.Value {
			curr = curr.Left
		} else {
			curr = curr.Right
		}
	}
	return curr
}

func (n *Node) Min() *Node {
	for n.Left != nil {
		n = n.Left
	}
	return n
}

func (n *Node) Max() *Node {
	for n.Right != nil {
		n = n.Right
	}
	return n
}

func (n *Node) Succ() *Node {
	if n.Right != nil {
		return n.Right.Min()
	}
	succ := n.Parent
	for succ != nil && n == succ.Right {
		n = succ
		succ = succ.Parent
	}
	return succ
}

func (n *Node) Pred() *Node {
	if n.Left != nil {
		return n.Left.Max()
	}
	pred := n.Parent
	for pred != nil && n == pred.Left {
		n = pred
		pred = pred.Parent
	}
	return pred
}

func (t *Tree) Insert(v *Node) {
	var tmp *Node
	curr := t.Root
	for curr != nil {
		tmp = curr
		if v.Value < curr.Value {
			curr = curr.Left
		} else {
			curr = curr.Right
		}
	}
	v.Parent = tmp
	if tmp == nil {
		t.Root = v // tree was empty
	} else if v.Value < tmp.Value {
		tmp.Left = v
	} else {
		tmp.Right = v
	}
}

// Transplant replaces one subtree as a child of its parent with another subtree
func (t *Tree) Transplant(original, replacement *Node) {
	if original.Parent == nil { // Root of the tree
		t.Root = replacement
	} else if original == original.Parent.Left { // Left child
		original.Parent.Left = replacement
	} else { // Right child
		original.Parent.Right = replacement
	}
	// Don't forget to update the parent node
	if replacement != nil {
		replacement.Parent = original.Parent
	}
}

func (t *Tree) Search(v int) *Node {
	if t.Root != nil {
		return t.Root.Search(v)
	}
	return nil
}

func (t *Tree) Min() *Node {
	if t.Root != nil {
		return t.Root.Min()
	}
	return nil
}

func (t *Tree) Max() *Node {
	if t.Root != nil {
		return t.Root.Max()
	}
	return nil
}
