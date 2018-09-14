package intset

import (
	"fmt"
	"testing"
)

func TestIntersectWith(t *testing.T) {
	tests := []struct {
		s, t *IntSet
		want *IntSet
	}{
		{NewIntSet(), NewIntSet(), NewIntSet()},
		{NewIntSet(0, 1), NewIntSet(), NewIntSet()},
		{NewIntSet(), NewIntSet(0, 1), NewIntSet()},
		{NewIntSet(0, 1), NewIntSet(1, 2), NewIntSet(1)},
		{NewIntSet(0, 1, 64, 65), NewIntSet(1, 2, 65, 66), NewIntSet(1, 65)},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("%v.IntersectWith(%v)", test.s, test.t)
		test.s.IntersectWith(test.t)
		if !test.s.Equal(test.want) {
			t.Errorf("%s = %v, want %v", descr, test.s, test.want)
		}
	}
}

func TestDifferenceWith(t *testing.T) {
	tests := []struct {
		s, t *IntSet
		want *IntSet
	}{
		{NewIntSet(), NewIntSet(), NewIntSet()},
		{NewIntSet(), NewIntSet(0, 1), NewIntSet()},
		{NewIntSet(0, 1), NewIntSet(), NewIntSet(0, 1)},
		{NewIntSet(0, 1, 64, 65), NewIntSet(1, 65), NewIntSet(0, 64)},
		{NewIntSet(0, 1, 65), NewIntSet(1, 65), NewIntSet(0)},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("%v.DifferenceWith(%v)", test.s, test.t)
		test.s.DifferenceWith(test.t)
		if !test.s.Equal(test.want) {
			t.Errorf("%s = %v, want %v", descr, test.s, test.want)
		}
	}
}

// todo: 内部のスライスの長さのバリエーションを増やしたほうが良い
func TestSymmetricDifference(t *testing.T) {
	tests := []struct {
		s, t *IntSet
		want *IntSet
	}{
		{NewIntSet(), NewIntSet(), NewIntSet()},
		{NewIntSet(0, 1, 2), NewIntSet(1, 2, 3), NewIntSet(0, 3)},
		{NewIntSet(0, 1, 64, 65), NewIntSet(1, 2, 65, 66), NewIntSet(0, 2, 64, 66)},
		{NewIntSet(0, 64), NewIntSet(), NewIntSet(0, 64)},
		{NewIntSet(), NewIntSet(0, 64), NewIntSet(0, 64)},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("%v.SymmetricDifference(%v)", test.s, test.t)
		test.s.SymmetricDifference(test.t)
		if !test.s.Equal(test.want) {
			t.Errorf("%s = %v, want %v", descr, test.s, test.want)
		}
	}
}
