package tree

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

func TestSearch(t *testing.T) {
	tr := new(Tree)
	tr.Insert(&Node{Value: 5})
	tr.Insert(&Node{Value: 3})
	tr.Insert(&Node{Value: 10})
	n := tr.Search(5)
	if n.Value != 5 {
		t.Fatalf("Expected Value=%v, got Value=%v\n", 5, n.Value)
	}
	n = tr.Search(10)
	if n.Value != 10 {
		t.Fatalf("Expected Value=%v, got Value=%v\n", 10, n.Value)
	}
	n = tr.Search(1000)
	if n != nil {
		t.Fatalf("Expected Node=%v, got Node=%v\n", nil, n)
	}
}

func TestMax(t *testing.T) {
	tr := new(Tree)
	tr.Insert(&Node{Value: 5})
	tr.Insert(&Node{Value: 10})
	tr.Insert(&Node{Value: 1})
	tr.Insert(&Node{Value: 11})
	tr.Insert(&Node{Value: 3})
	n := tr.Max()
	if n.Value != 11 {
		t.Fatalf("Expected Value=%v, got Value=%v\n", 11, n.Value)
	}
}

func TestMin(t *testing.T) {
	tr := new(Tree)
	tr.Insert(&Node{Value: 5})
	tr.Insert(&Node{Value: 10})
	tr.Insert(&Node{Value: 1})
	tr.Insert(&Node{Value: 11})
	tr.Insert(&Node{Value: 3})
	n := tr.Min()
	if n.Value != 1 {
		t.Fatalf("Expected Value=%v, got Value=%v\n", 1, n.Value)
	}
}
