package tree

func Root(tr Tree) Node {
	return tr.Root()
}

func Search(tr Tree, v Orderable) Node {
	return tr.Search(v)
}

func Min(tr Tree) Node {
	return tr.Min()
}

func Max(tr Tree) Node {
	return tr.Max()
}

func Insert(tr Tree, v Orderable) {
	tr.Insert(v)
}

func Delete(tr Tree, v Orderable) {
	tr.Delete(v)
}
