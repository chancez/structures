package tree

// func NodeMin(node Node) Node {
// 	return node.Min()
// }

// func NodeMax(node Node) Node {
// 	return node.Max()
// }

// func NodeSearch(node Node, v Orderable) Node {
// 	return node.Search(v)
// }

func NodeSucc(node Node) Node {
	if node == nil {
		return nil
	}
	if node.Right() != nil {
		return node.Right().Min()
	}
	succ := node.Parent()
	for succ != nil && node == succ.Right() {
		node = succ
		succ = succ.Parent()
	}
	return succ
}

func NodePred(node Node) Node {
	if node == nil {
		return nil
	}
	if node.Left() != nil {
		return node.Left().Max()
	}
	pred := node.Parent()
	for pred != nil && node == pred.Left() {
		node = pred
		pred = pred.Parent()
	}
	return pred
}
