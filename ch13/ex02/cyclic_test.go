package cyclic

import (
	"fmt"
	"testing"
)

func TestCyclic(t *testing.T) {
	type link struct {
		value string
		tail  *link
	}
	a, b := &link{value: "a"}, &link{value: "b"}
	a.tail, b.tail = b, a

	tests := []struct {
		in   interface{}
		want bool
	}{
		{in: struct{ x, y int }{1, 2}, want: false},
		{in: a, want: true},
		{in: []*link{a, b}, want: true},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("cyclic(%v)", test.in)
		got := cyclic(test.in)
		if got != test.want {
			t.Errorf("%s = %v, want %v", descr, got, test.want)
		}
	}
}
