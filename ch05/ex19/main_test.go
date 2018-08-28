package main

import (
	"fmt"
	"testing"
)

func TestF(t *testing.T) {
	tests := []struct {
		val  int
		want int
	}{
		{0, 0},
		{1, 1},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("f(%d)", test.val)
		got := f(test.val)
		if got != test.want {
			t.Errorf("%s = %d, want %d", descr, test.val, test.want)
		}
	}
}
