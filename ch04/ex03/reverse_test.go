package main

import (
	"fmt"
	"testing"
)

func TestReverse(t *testing.T) {
	var tests = []struct {
		s    [5]int
		want [5]int
	}{
		{[5]int{0, 1, 2, 3, 4}, [5]int{4, 3, 2, 1, 0}},
		{[5]int{4, 3, 2, 1, 0}, [5]int{0, 1, 2, 3, 4}},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("reverse(%v)", test.s)
		reverse(&test.s)
		if test.s != test.want {
			t.Errorf("%s => %v, want %v", descr, test.s, test.want)
		}
	}
}
