package intset

import (
	"fmt"
	"reflect"
	"testing"
)

func TestString(t *testing.T) {
	tests := []struct {
		s    *IntSet
		want string
	}{
		{&IntSet{}, "{}"},
		{&IntSet{[]uint{1}}, "{0}"},
		{&IntSet{[]uint{0, 1}}, "{64}"},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("%#v.String()", test.s)
		got := test.s.String()
		if got != test.want {
			t.Errorf("%s = %q, want %q", descr, got, test.want)
		}
	}
}

func TestHas(t *testing.T) {
	tests := []struct {
		s    *IntSet
		x    int
		want bool
	}{
		{&IntSet{}, 0, false},
		{&IntSet{[]uint{1}}, 0, true},
		{&IntSet{[]uint{0, 1}}, 64, true},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("%v.Has(%v)", test.s, test.x)
		got := test.s.Has(test.x)
		if got != test.want {
			t.Errorf("%s = %v, want %v", descr, got, test.want)
		}
	}
}

func TestAdd(t *testing.T) {
	tests := []struct {
		s    *IntSet
		x    int
		want *IntSet
	}{
		{&IntSet{}, 0, &IntSet{[]uint{1}}},
		{&IntSet{[]uint{1}}, 1, &IntSet{[]uint{3}}},
		{&IntSet{[]uint{0}}, 64, &IntSet{[]uint{0, 1}}},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("%v.Add(%v)", test.s, test.x)
		test.s.Add(test.x)
		if !reflect.DeepEqual(test.s, test.want) {
			t.Errorf("%s => %v, want %v", descr, test.s, test.want)
		}
	}
}

func TestUnionWith(t *testing.T) {
	tests := []struct {
		s, t *IntSet
		want *IntSet
	}{
		{&IntSet{}, &IntSet{}, &IntSet{}},
		{&IntSet{[]uint{1}}, &IntSet{[]uint{2}}, &IntSet{[]uint{3}}},
		{&IntSet{[]uint{1, 2}}, &IntSet{[]uint{2, 1}}, &IntSet{[]uint{3, 3}}},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("%v.UnionWith(%v)", test.s, test.t)
		test.s.UnionWith(test.t)
		if !reflect.DeepEqual(test.s, test.want) {
			t.Errorf("%s => %v, want %v", descr, test.s, test.want)
		}
	}
}
