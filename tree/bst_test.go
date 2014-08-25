package tree

import (
	"testing"
)

func TestInsert(t *testing.T) {
	tr := new(Tree)
	n1 := &Node{value: IntValue(5)}
	tr.Insert(n1)
	if tr.root != n1 {
		t.Fatalf("Expected tr.root=%v, got tr.root=%v\n", n1, tr.root)
	}
	n2 := &Node{value: IntValue(10)}
	tr.Insert(n2)
	if tr.root.right != n2 {
		t.Fatalf("Expected tr.root=%v, got tr.root=%v\n", n2, tr.root.right)
	}
	n3 := &Node{value: IntValue(6)}
	tr.Insert(n3)
	if tr.root.right.left != n3 {
		t.Fatalf("Expected tr.root=%v, got tr.root=%v\n", n3, tr.root.right.left)
	}
	n4 := &Node{value: IntValue(2)}
	tr.Insert(n4)
	if tr.root.left != n4 {
		t.Fatalf("Expected tr.root=%v, got tr.root=%v\n", n4, tr.root.left)
	}
}

func TestSearch(t *testing.T) {
	tr := new(Tree)
	n := tr.Search(IntValue(100))
	if n != nil {
		t.Fatal("Searching an empty tree should return nil")
	}
	tr.Insert(&Node{value: IntValue(5)})
	tr.Insert(&Node{value: IntValue(3)})
	tr.Insert(&Node{value: IntValue(10)})
	n = tr.Search(IntValue(5))
	if !Eq(n.value, IntValue(5)) {
		t.Fatalf("Expected value=%v, got value=%v\n", 5, n.value)
	}
	n = tr.Search(IntValue(10))
	if !Eq(n.value, IntValue(10)) {
		t.Fatalf("Expected value=%v, got value=%v\n", 10, n.value)
	}
	n = tr.Search(IntValue(1000))
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
	tr.Insert(&Node{value: IntValue(5)})
	tr.Insert(&Node{value: IntValue(10)})
	tr.Insert(&Node{value: IntValue(1)})
	tr.Insert(&Node{value: IntValue(11)})
	tr.Insert(&Node{value: IntValue(3)})
	n = tr.Max()
	if !Eq(n.value, IntValue(11)) {
		t.Fatalf("Expected value=%v, got value=%v\n", 11, n.value)
	}
}

func TestMin(t *testing.T) {
	tr := new(Tree)
	n := tr.Min()
	if n != nil {
		t.Fatal("Min of empty tree should be nil!")
	}
	tr.Insert(&Node{value: IntValue(5)})
	tr.Insert(&Node{value: IntValue(10)})
	tr.Insert(&Node{value: IntValue(1)})
	tr.Insert(&Node{value: IntValue(11)})
	tr.Insert(&Node{value: IntValue(3)})
	n = tr.Min()
	if !Eq(n.value, IntValue(1)) {
		t.Fatalf("Expected value=%v, got value=%v\n", 1, n.value)
	}
}

func TestNodeSucc(t *testing.T) {
	tr := new(Tree)
	tr.Insert(&Node{value: IntValue(10)})
	tr.Insert(&Node{value: IntValue(5)})
	n := &Node{value: IntValue(8)}
	tr.Insert(n)
	succ := n.Succ()
	if !Eq(succ.value, IntValue(10)) {
		t.Fatalf("Expected Succ=%v, got Succ=%v\n", 10, succ.value)
	}
	tr.Insert(&Node{value: IntValue(9)})
	succ = n.Succ()
	if !Eq(succ.value, IntValue(9)) {
		t.Fatalf("Expected Succ=%v, got Succ=%v\n", 9, succ.value)
	}
}

func TestNodePred(t *testing.T) {
	tr := new(Tree)
	tr.Insert(&Node{value: IntValue(10)})
	tr.Insert(&Node{value: IntValue(5)})
	n := &Node{value: IntValue(7)}
	tr.Insert(n)
	pred := n.Pred()
	if !Eq(pred.value, IntValue(5)) {
		t.Fatalf("Expected Pred=%v, got Pred=%v\n", 5, pred.value)
	}
	tr.Insert(&Node{value: IntValue(6)})
	pred = n.Pred()
	if !Eq(pred.value, IntValue(6)) {
		t.Fatalf("Expected Pred=%v, got Pred=%v\n", 6, pred.value)
	}
}

func TestTransplant(t *testing.T) {
	tr := new(Tree)
	tr.Insert(&Node{value: IntValue(7)})
	// Create a new tree with a root of 10, left of 5, and right of 15
	// (has 2 children)
	subTree1 := new(Tree)
	subTree1.Insert(&Node{value: IntValue(10)})
	subTree1.Insert(&Node{value: IntValue(15)})
	subTree1.Insert(&Node{value: IntValue(5)})
	// Test replacing the root of a tree
	tr.Transplant(tr.root, subTree1.root)
	if tr.root != subTree1.root {
		t.Fatalf("Expected tree.root=%v, got tree.root=%v", subTree1.root, tr.root)
	}
	subTree2 := new(Tree)
	subTree2.Insert(&Node{value: IntValue(4)})
	subTree2.Insert(&Node{value: IntValue(7)})
	tr.Transplant(tr.root.left, subTree2.root)
	if tr.root.left != subTree2.root {
		t.Fatalf("Expected tree.root.left=%v, got tree.root.left=%v", subTree2.root, tr.root.left)
	}
	subTree3 := new(Tree)
	subTree3.Insert(&Node{value: IntValue(20)})
	subTree3.Insert(&Node{value: IntValue(21)})
	// Replace the right child subtree
	tr.Transplant(tr.root.right, subTree3.root)
	if tr.root.right != subTree3.root {
		t.Fatalf("Expected tree.root.right=%v, got tree.root.right=%v", subTree3.root, tr.root.right)
	}
}
