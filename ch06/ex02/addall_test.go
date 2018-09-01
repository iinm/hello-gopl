package intset

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAddAll(t *testing.T) {
	tests := []struct {
		s    *IntSet
		xs   []int
		want *IntSet
	}{
		{&IntSet{}, []int{}, &IntSet{}},                               // {} + {} = {}
		{&IntSet{}, []int{0, 1, 2}, &IntSet{[]uint64{7}}},             // {} + {0, 1, 2} = {0, 1, 2}
		{&IntSet{[]uint64{3}}, []int{2, 3}, &IntSet{[]uint64{15}}},    // {0, 1} + {2} = {0, 1, 2}
		{&IntSet{[]uint64{3}}, []int{1}, &IntSet{[]uint64{3}}},        // {0, 1} + {1} = {0, 1}
		{&IntSet{[]uint64{3, 3}}, []int{66}, &IntSet{[]uint64{3, 7}}}, // {0, 1, 64+0, 64+1} + {64+2} = {0, 1, 64+0, 64+1, 64+2}
	}

	for _, test := range tests {
		descr := fmt.Sprintf("%s.AddAll(%v...)", test.s, test.xs)
		test.s.AddAll(test.xs...)
		if !reflect.DeepEqual(test.s, test.want) {
			t.Errorf("%s = %s, want %s", descr, test.s, test.want)
		}
	}

}
