package tree

import (
	"testing"
)

func TestInsert(t *testing.T) {
	tr := new(Bst)
	n1 := &BstNode{value: IntValue(5)}
	tr.Insert(n1)
	if tr.root != n1 {
		t.Fatalf("Expected tr.root=%v, got tr.root=%v\n", n1, tr.root)
	}
	n2 := &BstNode{value: IntValue(10)}
	tr.Insert(n2)
	if tr.root.right != n2 {
		t.Fatalf("Expected tr.root=%v, got tr.root=%v\n", n2, tr.root.right)
	}
	n3 := &BstNode{value: IntValue(6)}
	tr.Insert(n3)
	if tr.root.right.left != n3 {
		t.Fatalf("Expected tr.root=%v, got tr.root=%v\n", n3, tr.root.right.left)
	}
	n4 := &BstNode{value: IntValue(2)}
	tr.Insert(n4)
	if tr.root.left != n4 {
		t.Fatalf("Expected tr.root=%v, got tr.root=%v\n", n4, tr.root.left)
	}
}

func TestSearch(t *testing.T) {
	tr := new(Bst)
	n := tr.Search(IntValue(100))
	if n != nil {
		t.Fatalf("Searching an empty tree should return nil, got %v\n", n)
	}
	tr.Insert(&BstNode{value: IntValue(5)})
	tr.Insert(&BstNode{value: IntValue(3)})
	tr.Insert(&BstNode{value: IntValue(10)})
	n = tr.Search(IntValue(5))
	if !Eq(n.Value(), IntValue(5)) {
		t.Fatalf("Expected value=%v, got value=%v\n", IntValue(5), n.Value())
	}
	n = tr.Search(IntValue(10))
	if !Eq(n.Value(), IntValue(10)) {
		t.Fatalf("Expected value=%v, got value=%v\n", IntValue(10), n.Value())
	}
	n = tr.Search(IntValue(1))
	if n != (*BstNode)(nil) {
		t.Fatalf("Expected Node=%v, got Node=%v\n", nil, n)
	}
}

func TestMax(t *testing.T) {
	tr := new(Bst)
	n := tr.Max()
	if n != nil {
		t.Fatal("Max of empty tree should be nil!")
	}
	tr.Insert(&BstNode{value: IntValue(5)})
	tr.Insert(&BstNode{value: IntValue(10)})
	tr.Insert(&BstNode{value: IntValue(1)})
	tr.Insert(&BstNode{value: IntValue(11)})
	tr.Insert(&BstNode{value: IntValue(3)})
	n = tr.Max()
	if !Eq(n.Value(), IntValue(11)) {
		t.Fatalf("Expected value=%v, got value=%v\n", IntValue(11), n.Value())
	}
}

func TestMin(t *testing.T) {
	tr := new(Bst)
	n := tr.Min()
	if n != nil {
		t.Fatal("Min of empty tree should be nil!")
	}
	tr.Insert(&BstNode{value: IntValue(5)})
	tr.Insert(&BstNode{value: IntValue(10)})
	tr.Insert(&BstNode{value: IntValue(1)})
	tr.Insert(&BstNode{value: IntValue(11)})
	tr.Insert(&BstNode{value: IntValue(3)})
	n = tr.Min()
	if !Eq(n.Value(), IntValue(1)) {
		t.Fatalf("Expected value=%v, got value=%v\n", IntValue(1), n.Value())
	}
}

func TestNodeSucc(t *testing.T) {
	tr := new(Bst)
	tr.Insert(&BstNode{value: IntValue(8)})
	tr.Insert(&BstNode{value: IntValue(10)})
	tr.Insert(&BstNode{value: IntValue(5)})
	succ := NodeSucc(tr.Root())
	if !Eq(succ.Value(), IntValue(10)) {
		t.Fatalf("Expected Succ=%v, got Succ=%v\n", IntValue(10), succ.Value())
	}
	tr.Insert(&BstNode{value: IntValue(9)})
	succ = NodeSucc(tr.Root())
	if !Eq(succ.Value(), IntValue(9)) {
		t.Fatalf("Expected Succ=%v, got Succ=%v\n", IntValue(9), succ.Value())
	}
}

func TestNodePred(t *testing.T) {
	tr := new(Bst)
	tr.Insert(&BstNode{value: IntValue(10)})
	tr.Insert(&BstNode{value: IntValue(5)})
	pred := NodePred(tr.Root())
	if !Eq(pred.Value(), IntValue(5)) {
		t.Fatalf("Expected Pred=%v, got Pred=%v\n", IntValue(5), pred.Value())
	}
	tr.Insert(&BstNode{value: IntValue(6)})
	pred = NodePred(tr.Root())
	if !Eq(pred.Value(), IntValue(6)) {
		t.Fatalf("Expected Pred=%v, got Pred=%v\n", IntValue(6), pred.Value())
	}
}

func TestTransplant(t *testing.T) {
	tr := new(Bst)
	tr.Insert(&BstNode{value: IntValue(7)})
	// Create a new tree with a root of 10, left of 5, and right of 15
	// (has 2 children)
	subTree1 := new(Bst)
	subTree1.Insert(&BstNode{value: IntValue(10)})
	subTree1.Insert(&BstNode{value: IntValue(15)})
	subTree1.Insert(&BstNode{value: IntValue(5)})
	// Test replacing the root of a tree
	tr.Transplant(tr.root, subTree1.root)
	if tr.root != subTree1.root {
		t.Fatalf("Expected tree.root=%v, got tree.root=%v", subTree1.root, tr.root)
	}
	subTree2 := new(Bst)
	subTree2.Insert(&BstNode{value: IntValue(4)})
	subTree2.Insert(&BstNode{value: IntValue(7)})
	tr.Transplant(tr.root.left, subTree2.root)
	if tr.root.left != subTree2.root {
		t.Fatalf("Expected tree.root.left=%v, got tree.root.left=%v", subTree2.root, tr.root.left)
	}
	subTree3 := new(Bst)
	subTree3.Insert(&BstNode{value: IntValue(20)})
	subTree3.Insert(&BstNode{value: IntValue(21)})
	// Replace the right child subtree
	tr.Transplant(tr.root.right, subTree3.root)
	if tr.root.right != subTree3.root {
		t.Fatalf("Expected tree.root.right=%v, got tree.root.right=%v", subTree3.root, tr.root.right)
	}
}
