package main

func reverse(s *[5]int) {
	for i, j := 0, 4; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
