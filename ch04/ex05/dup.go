package main

func dup(xs []string) []string {
	if len(xs) < 2 {
		return xs
	}
	j := 1
	for i := 1; i < len(xs); i++ {
		if xs[i] != xs[i-1] {
			xs[j] = xs[i]
			j++
		}
	}
	return xs[:j]
}
