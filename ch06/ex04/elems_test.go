package intset

import (
	"fmt"
	"reflect"
	"testing"
)

func TestElems(t *testing.T) {
	tests := []struct {
		s    *IntSet
		want []int
	}{
		{NewIntSet(), []int{}},
		{NewIntSet(0, 1), []int{0, 1}},
		{NewIntSet(0, 1, 64, 65), []int{0, 1, 64, 65}},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("%v.Elems()", test.s)
		got := test.s.Elems()
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("%s = %v, want %v", descr, got, test.want)
		}
	}
}
