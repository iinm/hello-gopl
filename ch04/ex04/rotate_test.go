package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRotate(t *testing.T) {
	var tests = []struct {
		s    []int
		step int
		want []int
	}{
		{[]int{0, 1, 2, 3, 4}, 1, []int{1, 2, 3, 4, 0}},
		{[]int{0, 1, 2, 3, 4}, 2, []int{2, 3, 4, 0, 1}},
		{[]int{0, 1, 2, 3, 4}, 3, []int{3, 4, 0, 1, 2}},
		{[]int{0, 1, 2, 3, 4}, 4, []int{4, 0, 1, 2, 3}},
		{[]int{0, 1, 2, 3, 4}, 5, []int{0, 1, 2, 3, 4}},
		{[]int{0, 1, 2, 3, 4}, 6, []int{1, 2, 3, 4, 0}},

		{[]int{0, 1, 2, 3, 4}, -1, []int{4, 0, 1, 2, 3}},
		{[]int{0, 1, 2, 3, 4}, -2, []int{3, 4, 0, 1, 2}},
		{[]int{0, 1, 2, 3, 4}, -6, []int{4, 0, 1, 2, 3}},

		{[]int{0, 1, 2, 3}, 1, []int{1, 2, 3, 0}},
		{[]int{0, 1, 2, 3}, 2, []int{2, 3, 0, 1}},
		{[]int{0, 1, 2, 3}, 3, []int{3, 0, 1, 2}},
		{[]int{0, 1, 2, 3}, 4, []int{0, 1, 2, 3}},
		{[]int{0, 1, 2, 3}, 5, []int{1, 2, 3, 0}},
		{[]int{0, 1, 2, 3}, -1, []int{3, 0, 1, 2}},
		{[]int{0, 1, 2, 3}, -2, []int{2, 3, 0, 1}},

		{[]int{0, 1, 2, 3, 4, 5, 6, 7}, 2, []int{2, 3, 4, 5, 6, 7, 0, 1}},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("rotate(%v), %v", test.s, test.step)
		rotate(test.s, test.step)
		if !reflect.DeepEqual(test.s, test.want) {
			t.Errorf("%s => %v, want %v", descr, test.s, test.want)
		}
	}
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func rotateByReverse(s []int, step int) {
	reverse(s[:step])
	reverse(s[step:])
	reverse(s)
}

func BenchmarkRatateByReverse(b *testing.B) {
	s := []int{0, 1, 2, 3, 4}
	for i := 0; i < b.N; i++ {
		rotateByReverse(s, 2)
	}
}

func BenchmarkRatate(b *testing.B) {
	s := []int{0, 1, 2, 3, 4}
	for i := 0; i < b.N; i++ {
		rotate(s, 2)
	}
}
