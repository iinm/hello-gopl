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
		{[]string{}, "\n"},
		{[]string{"echo", "foo"}, "echo foo\n"},
		{[]string{"echo", "foo", "bar"}, "echo foo bar\n"},
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
