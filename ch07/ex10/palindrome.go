package palindrom

import "sort"

func IsPandrome(s sort.Interface) bool {
	for i := 0; i < s.Len()/2; i++ {
		if !equal(s, i, s.Len()-1-i) {
			return false
		}
	}
	return true
}

func equal(s sort.Interface, i, j int) bool {
	return !s.Less(i, j) && !s.Less(j, i)
}
