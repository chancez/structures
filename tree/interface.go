package tree

type Orderable interface {
	Lt(Orderable) bool
	Eq(Orderable) bool
}

type IntValue int
type StringValue string
type FloatValue float64

func (i IntValue) Lt(j Orderable) bool {
	return i < j.(IntValue)
}

func (i IntValue) Eq(j Orderable) bool {
	return i == j.(IntValue)
}

func (i StringValue) Lt(j Orderable) bool {
	return i < j.(StringValue)
}

func (i StringValue) Eq(j Orderable) bool {
	return i == j.(StringValue)
}

func (i FloatValue) Lt(j Orderable) bool {
	return i < j.(FloatValue)
}

func (i FloatValue) Eq(j Orderable) bool {
	return i == j.(FloatValue)
}

func Lt(i, j Orderable) bool {
	return i.Lt(j)
}

func Eq(i, j Orderable) bool {
	return i.Eq(j)
}

type NodeBase interface {
	Search(Orderable) Node
	Min() Node
	Max() Node
}

type Node interface {
	NodeBase
	Left() Node
	Right() Node
	Parent() Node
	Value() Orderable
}

type Tree interface {
	NodeBase // A tree is simply a node with a root
	Root() Node
	// These are implementation specific
	Insert(Orderable)
	Delete(Orderable)
}
