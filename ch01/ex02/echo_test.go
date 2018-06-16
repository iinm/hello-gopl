package main

import (
	"bytes"
	"fmt"
	"testing"
)

func TestEcho(t *testing.T) {
	var tests = []struct {
		args []string
		want string
	}{
		{[]string{}, ""},
		{[]string{"echo", "foo"}, "0\techo\n1\tfoo\n"},
		{[]string{"echo", "foo", "bar"}, "0\techo\n1\tfoo\n2\tbar\n"},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("echo(%q)", test.args)

		out = new(bytes.Buffer)
		echo(test.args)
		got := out.(*bytes.Buffer).String()
		if got != test.want {
			t.Errorf("%s = %q, want %q", descr, got, test.want)
		}
	}
}
