package tree

import (
	"fmt"
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
	n := tr.Search(100)
	if n != nil {
		t.Fatal("Searching an empty tree should return nil")
	}
	tr.Insert(&Node{Value: 5})
	tr.Insert(&Node{Value: 3})
	tr.Insert(&Node{Value: 10})
	n = tr.Search(5)
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
	n := tr.Max()
	if n != nil {
		t.Fatal("Max of empty tree should be nil!")
	}
	tr.Insert(&Node{Value: 5})
	tr.Insert(&Node{Value: 10})
	tr.Insert(&Node{Value: 1})
	tr.Insert(&Node{Value: 11})
	tr.Insert(&Node{Value: 3})
	n = tr.Max()
	if n.Value != 11 {
		t.Fatalf("Expected Value=%v, got Value=%v\n", 11, n.Value)
	}
}

func TestMin(t *testing.T) {
	tr := new(Tree)
	n := tr.Min()
	if n != nil {
		t.Fatal("Min of empty tree should be nil!")
	}
	tr.Insert(&Node{Value: 5})
	tr.Insert(&Node{Value: 10})
	tr.Insert(&Node{Value: 1})
	tr.Insert(&Node{Value: 11})
	tr.Insert(&Node{Value: 3})
	n = tr.Min()
	if n.Value != 1 {
		t.Fatalf("Expected Value=%v, got Value=%v\n", 1, n.Value)
	}
}

func TestNodeSucc(t *testing.T) {
	tr := new(Tree)
	tr.Insert(&Node{Value: 10})
	tr.Insert(&Node{Value: 5})
	n := &Node{Value: 8}
	tr.Insert(n)
	succ := n.Succ()
	if succ.Value != 10 {
		t.Fatalf("Expected Succ=%v, got Succ=%v\n", 10, succ.Value)
	}
	tr.Insert(&Node{Value: 9})
	succ = n.Succ()
	if succ.Value != 9 {
		t.Fatalf("Expected Succ=%v, got Succ=%v\n", 9, succ.Value)
	}
}

func TestNodePred(t *testing.T) {
	tr := new(Tree)
	tr.Insert(&Node{Value: 10})
	tr.Insert(&Node{Value: 5})
	n := &Node{Value: 7}
	tr.Insert(n)
	pred := n.Pred()
	if pred.Value != 5 {
		t.Fatalf("Expected Pred=%v, got Pred=%v\n", 5, pred.Value)
	}
	tr.Insert(&Node{Value: 6})
	pred = n.Pred()
	if pred.Value != 6 {
		t.Fatalf("Expected Pred=%v, got Pred=%v\n", 6, pred.Value)
	}
}

func TestTransplant(t *testing.T) {
	tr := new(Tree)
	tr.Insert(&Node{Value: 7})
	// Create a new tree with a root of 10, left of 5, and right of 15
	// (has 2 children)
	subTree1 := new(Tree)
	subTree1.Insert(&Node{Value: 10})
	subTree1.Insert(&Node{Value: 15})
	subTree1.Insert(&Node{Value: 5})
	// Test replacing the root of a tree
	tr.Transplant(tr.Root, subTree1.Root)
	if tr.Root != subTree1.Root {
		t.Fatalf("Expected tree.Root=%v, got tree.Root=%v", subTree1.Root, tr.Root)
	}
	subTree2 := new(Tree)
	subTree2.Insert(&Node{Value: 4})
	subTree2.Insert(&Node{Value: 7})
	tr.Transplant(tr.Root.Left, subTree2.Root)
	if tr.Root.Left != subTree2.Root {
		t.Fatalf("Expected tree.Root.Left=%v, got tree.Root.Left=%v", subTree2.Root, tr.Root.Left)
	}
	subTree3 := new(Tree)
	subTree3.Insert(&Node{Value: 20})
	subTree3.Insert(&Node{Value: 21})
	// Replace the right child subtree
	fmt.Printf("%+v\n", tr.Root.Right)
	tr.Transplant(tr.Root.Right, subTree3.Root)
	if tr.Root.Right != subTree3.Root {
		t.Fatalf("Expected tree.Root.Right=%v, got tree.Root.Right=%v", subTree3.Root, tr.Root.Right)
	}
}
