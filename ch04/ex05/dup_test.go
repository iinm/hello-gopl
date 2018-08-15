package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDup(t *testing.T) {
	var tests = []struct {
		xs   []string
		want []string
	}{
		{[]string{}, []string{}},
		{[]string{"a", "a", "a"}, []string{"a"}},
		{[]string{"a", "b", "c"}, []string{"a", "b", "c"}},
		{[]string{"a", "b", "b", "c"}, []string{"a", "b", "c"}},
		{[]string{"a", "b", "c", "b"}, []string{"a", "b", "c", "b"}},
		{[]string{"a", "a", "b", "b", "c"}, []string{"a", "b", "c"}},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("dup(%v)", test.xs)
		got := dup(test.xs)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("%s = %v, want %v", descr, test.xs, test.want)
		}
	}
}
