package main

func rotate(s []int, step int) {
	lenS := len(s)
	step = step % lenS
	i := 0
	tmp := s[0]
	for {
		// 移動先
		dest := i - step
		if dest < 0 {
			dest = lenS + dest
		} else if lenS <= dest {
			dest = dest - lenS
		}
		s[dest], tmp = tmp, s[dest]
		i = dest
		if i == 0 {
			break
		}
	}
}
