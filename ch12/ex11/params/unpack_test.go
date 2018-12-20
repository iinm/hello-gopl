package params

import (
	"fmt"
	"testing"
)

func TestUnpack(t *testing.T) {
	type data struct {
		Labels      []string `http:"l"`
		MaxRequests int      `http:"max"`
		Exact       bool     `http:"x"`
	}
	tests := []struct {
		input data
		want  string
	}{
		{input: data{[]string{"hello"}, 10, true}, want: `l=hello&max=10&x=true`},
		{input: data{[]string{"hello", "world"}, 0, false}, want: `l=hello&l=world&max=0&x=false`},
		{input: data{[]string{}, 0, false}, want: `max=0&x=false`},
		{input: data{nil, 0, false}, want: `max=0&x=false`},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("Unpack(%#v)", test.input)
		got := Unpack(test.input)
		if got != test.want {
			t.Errorf("%s = %q, want %q", descr, got, test.want)
		}
	}
}
