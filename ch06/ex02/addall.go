package intset

// AddAll はxsのすべての要素をsに追加する。
func (s *IntSet) AddAll(xs ...int) {
	for _, x := range xs {
		s.Add(x)
	}
}
