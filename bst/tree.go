package bst

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

func Search(n *Node, v int) *Node {
	if n == nil || v == n.Value {
		return n
	}
	if v < n.Value {
		return Search(n.Left, v)
	}
	return Search(n.Right, v)
}

func SearchIter(n *Node, v int) *Node {
	for n != nil && v != n.Value {
		if v < n.Value {
			n = n.Left
		} else {
			n = n.Right
		}
	}
	return n
}

func Min(n *Node) *Node {
	for n.Left != nil {
		n = n.Left
	}
	return n
}

func Max(n *Node) *Node {
	for n.Right != nil {
		n = n.Right
	}
	return n
}

func Succ(n *Node) *Node {
	if n.Right != nil {
		return Min(n.Right)
	}
	succ := n.Parent
	for succ != nil && succ == n.Right {
		n = succ
		succ = succ.Parent
	}
	return succ
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
