package intset

func (s *IntSet) Equal(t *IntSet) bool {
	i := 0
	for ; i < len(s.words) && i < len(t.words); i++ {
		if s.words[i] != t.words[i] {
			return false
		}
	}
	for ; i < len(s.words); i++ {
		if s.words[i] != 0 {
			return false
		}
	}
	for ; i < len(t.words); i++ {
		if t.words[i] != 0 {
			return false
		}
	}
	return true
}
