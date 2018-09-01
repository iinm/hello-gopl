package intset

func NewIntSet(xs ...int) *IntSet {
	s := &IntSet{}
	s.AddAll(xs...)
	return s
}
