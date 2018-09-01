package intset

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewIntSet(t *testing.T) {
	tests := []struct {
		xs   []int
		want *IntSet
	}{
		{nil, &IntSet{}},
		{[]int{}, &IntSet{}},
		{[]int{0, 1, 2}, &IntSet{[]uint64{7}}},
		{[]int{0, 64}, &IntSet{[]uint64{1, 1}}},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("NewIntSet(%v)", test.xs)
		got := NewIntSet(test.xs...)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("%s = %v, want %v", descr, got, test.want)
		}
	}
}
