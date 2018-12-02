package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestMain(t *testing.T) {
	var tests = []struct {
		input string
		want  string
	}{
		{
			input: `abcあいう`,
			want: `rune	count
'a'	1
'b'	1
'c'	1
'あ'	1
'い'	1
'う'	1

len	count
1	3
2	0
3	3
4	0
`,
		},
	}

	for _, test := range tests {
		in = strings.NewReader(test.input)
		out = &bytes.Buffer{}
		main()
		got := out.(*bytes.Buffer).String()
		if got != test.want {
			t.Errorf("input:\n%s\noutput:\n%s\nwant:\n%s", test.input, got, test.want)
		}
	}

}
