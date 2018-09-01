package intset

import (
	"fmt"
	"testing"
)

func TestEqual(t *testing.T) {
	tests := []struct {
		s, t *IntSet
		want bool
	}{
		{&IntSet{}, &IntSet{}, true},
		{&IntSet{}, &IntSet{[]uint64{}}, true},
		{&IntSet{[]uint64{}}, &IntSet{}, true},
		{&IntSet{[]uint64{1, 2}}, &IntSet{[]uint64{1, 2}}, true},
		{&IntSet{[]uint64{1, 2, 0}}, &IntSet{[]uint64{1, 2}}, true},
		{&IntSet{[]uint64{1, 2}}, &IntSet{[]uint64{1, 2, 0}}, true},
		{&IntSet{[]uint64{1, 2}}, &IntSet{}, false},
		{&IntSet{[]uint64{1, 2}}, &IntSet{[]uint64{1}}, false},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("%#v.Equal(%#v)", test.s, test.t)
		got := test.s.Equal(test.t)
		if got != test.want {
			t.Errorf("%s = %#v, want %#v", descr, got, test.want)
		}
	}
}
