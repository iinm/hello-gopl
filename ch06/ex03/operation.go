package intset

func (s *IntSet) IntersectWith(t *IntSet) {
	var i int
	for i = 0; i < len(s.words) && i < len(t.words); i++ {
		s.words[i] &= t.words[i]
	}
	s.words = s.words[:i]
}

func (s *IntSet) DifferenceWith(t *IntSet) {
	for i := 0; i < len(s.words) && i < len(t.words); i++ {
		s.words[i] &^= t.words[i]
	}
}

func (s *IntSet) SymmetricDifference(t *IntSet) {
	var i int
	for i = 0; i < len(s.words) && i < len(t.words); i++ {
		s.words[i] ^= t.words[i]
	}
	if i < len(t.words) {
		s.words = append(s.words, t.words[i:]...)
	}
}
