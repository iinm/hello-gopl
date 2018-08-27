package main

import (
	"fmt"
	"testing"
)

func TestJoin(t *testing.T) {
	tests := []struct {
		sep  string
		strs []string
		want string
	}{
		{", ", []string{}, ""},
		{", ", []string{"a"}, "a"},
		{", ", []string{"a", "b"}, "a, b"},
		{", ", []string{"a", "b", "c"}, "a, b, c"},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("join(%q, %q)", test.sep, test.strs)
		got := join(test.sep, test.strs...)
		if got != test.want {
			t.Errorf("%s = %q, want %q", descr, got, test.want)
		}
	}
}
