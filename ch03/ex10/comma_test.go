package main

import (
	"fmt"
	"testing"
)

func TestComma(t *testing.T) {
	var tests = []struct {
		s    string
		want string
	}{
		{"", ""},
		{"123", "123"},
		{"123456", "123,456"},
		{"1234567", "1,234,567"},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("comma(%s)", test.s)
		got := comma(test.s)
		if got != test.want {
			t.Errorf("%s = %q, want %q", descr, got, test.want)
		}
	}
}
