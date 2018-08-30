package intset

import (
	"fmt"
	"reflect"
	"testing"
)

func TestLen(t *testing.T) {
	tests := []struct {
		s    *IntSet
		want int
	}{
		{&IntSet{}, 0},                  // {}
		{&IntSet{[]uint64{3}}, 2},       // {0, 1}
		{&IntSet{[]uint64{3, 3}}, 4},    // {0, 1, 64+0, 64+1}
		{&IntSet{[]uint64{3, 3, 7}}, 7}, // {0, 1, 64+0, 64+1, 128+0, 128+1, 128+2}
	}

	for _, test := range tests {
		descr := fmt.Sprintf("%s.Len()", test.s)
		got := test.s.Len()
		if got != test.want {
			t.Errorf("%s = %d, want %d", descr, got, test.want)
		}
	}
}

func TestRemove(t *testing.T) {
	tests := []struct {
		s    *IntSet
		x    int
		want *IntSet
	}{
		{&IntSet{}, 0, &IntSet{}},                              // {} - {0} = {}
		{&IntSet{[]uint64{3}}, 2, &IntSet{[]uint64{3}}},        // {0, 1} - {2} = {0, 1}
		{&IntSet{[]uint64{3}}, 1, &IntSet{[]uint64{1}}},        // {0, 1} - {1} = {0}
		{&IntSet{[]uint64{3, 3}}, 65, &IntSet{[]uint64{3, 1}}}, // {0, 1, 64+0, 64+1} - {64+1} = {0, 1, 64+0}
	}

	for _, test := range tests {
		descr := fmt.Sprintf("%s.Remove(%d)", test.s, test.x)
		test.s.Remove(test.x)
		if !reflect.DeepEqual(test.s, test.want) {
			t.Errorf("%s = %s, want %s", descr, test.s, test.want)
		}
	}
}

func TestClear(t *testing.T) {
	want := &IntSet{}
	tests := []struct {
		s *IntSet
	}{
		{&IntSet{}},
		{&IntSet{[]uint64{}}},
		{&IntSet{[]uint64{0, 1, 1}}},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("%s.Clear()", test.s)
		test.s.Clear()
		if !reflect.DeepEqual(test.s, want) {
			t.Errorf("%s = %s, want {}", descr, want)
		}
	}
}

func TestCopy(t *testing.T) {
	tests := []struct {
		s *IntSet
	}{
		{&IntSet{}},
		{&IntSet{[]uint64{}}},
		{&IntSet{[]uint64{3}}},
		{&IntSet{[]uint64{3, 3}}},
		{&IntSet{[]uint64{3, 3, 7}}},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("%s.Copy()", test.s)
		clone := test.s.Copy()
		if !reflect.DeepEqual(clone, test.s) {
			t.Errorf("%s = %s, want %s", descr, clone, test.s)
		}
	}
}
