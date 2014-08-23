package bst

import (
	"testing"
)

func TestInsert(t *testing.T) {
	tr := new(Tree)
	n1 := &Node{Value: 5}
	tr.Insert(n1)
	if tr.Root != n1 {
		t.Fatalf("Expected tr.Root=%v, got tr.Root=%v\n", n1, tr.Root)
	}
	n2 := &Node{Value: 10}
	tr.Insert(n2)
	if tr.Root.Right != n2 {
		t.Fatalf("Expected tr.Root=%v, got tr.Root=%v\n", n2, tr.Root.Right)
	}
	n3 := &Node{Value: 6}
	tr.Insert(n3)
	if tr.Root.Right.Left != n3 {
		t.Fatalf("Expected tr.Root=%v, got tr.Root=%v\n", n3, tr.Root.Right.Left)
	}
	n4 := &Node{Value: 2}
	tr.Insert(n4)
	if tr.Root.Left != n4 {
		t.Fatalf("Expected tr.Root=%v, got tr.Root=%v\n", n4, tr.Root.Left)
	}
}
