package main

func rotate(s []int, step int) {
	L := len(s)
	// 0 < step < L
	step = step % L
	if step == 0 {
		return
	}
	if step < 0 {
		step = L + step
	}

	start := 0
	dest, backup := start, s[start]
	for range s {
		src := (dest + step) % L
		if src == start {
			s[dest] = backup
			start++
			dest, backup = start, s[start]
			continue
		}
		s[dest] = s[src]
		dest = src
	}
}
