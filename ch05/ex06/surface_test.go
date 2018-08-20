package main

import (
	"fmt"
	"testing"
)

func TestCorner(t *testing.T) {
	var tests = []struct {
		i, j int
		want bool
	}{
		{1, 1, true},
		{cells / 2, cells / 2, false},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("corner(%d, %d)", test.i, test.j)
		_, _, ok := corner(test.i, test.j)
		if ok != test.want {
			t.Errorf("%s = %v, want %v\n", descr, ok, test.want)
		}
	}
}
