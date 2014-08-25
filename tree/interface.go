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

type Node interface {
	Left() Node
	Right() Node
	Parent() Node
	Value() Orderable
}

type Tree interface {
	Root() Node
}
