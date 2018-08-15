package main

import (
	"fmt"
	"testing"
)

func TestBitDiffCount(t *testing.T) {
	var tests = []struct {
		c1, c2 [32]byte
		want   int
	}{
		{[32]byte{0: 1}, [32]byte{}, 1},
		{[32]byte{0: 1}, [32]byte{31: 3}, 3},
		{[32]byte{0: 1}, [32]byte{0: 1}, 0},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("BitDiffCount(%v, %v)", &test.c1, &test.c2)
		got := BitDiffCount(&test.c1, &test.c2)
		if got != test.want {
			t.Errorf("%s = %v, want %v", descr, got, test.want)
		}
	}
}
