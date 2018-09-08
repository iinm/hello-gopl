package reader

import (
	"bytes"
	"fmt"
	"io"
	"reflect"
	"strings"
	"testing"
)

func TestLimitReader(t *testing.T) {
	tests := []struct {
		src   io.Reader
		limit int64
		want  []byte
	}{
		{strings.NewReader(""), 0, []byte{}},
		{strings.NewReader("hello"), 0, []byte{}},
		{strings.NewReader("hello"), 4, []byte("hell")},
		{strings.NewReader("hello"), 5, []byte("hello")},
		{strings.NewReader("hello"), 7, []byte("hello")},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("LimitReader(%v)", test.src)
		r := LimitReader(test.src, test.limit)
		buf := new(bytes.Buffer)
		buf.ReadFrom(r)
		got := buf.Bytes()
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("%s -> %q, want %q", descr, got, test.want)
		}
	}
}
