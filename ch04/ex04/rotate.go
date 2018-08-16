package main

func rotate(s []int, step int) {
	lenS := len(s)
	step = step % lenS

	// 長さが偶数で、長さ/2の要素をずらす場合
	if lenS == 2*step || lenS == -2*step {
		for i := 0; i < lenS/2; i++ {
			s[i], s[lenS/2+i] = s[lenS/2+i], s[i]
		}
		return
	}

	i := 0
	tmp := s[0]
	for range s {
		// 移動先
		dest := i - step
		if dest < 0 {
			dest = lenS + dest
		} else if lenS <= dest {
			dest = dest - lenS
		}

		// iの要素をdestに移動、destの要素をtmpに退避
		s[dest], tmp = tmp, s[dest]
		// 次のループでdestの要素を移動する
		i = dest
	}
}
