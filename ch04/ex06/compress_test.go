package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCompressSpaces(t *testing.T) {
	var tests = []struct {
		b    []byte
		want []byte
	}{
		{[]byte("こんにちは、世界"), []byte("こんにちは、世界")},
		{[]byte("こんにちは、\t\n世界"), []byte("こんにちは、 世界")},
		{[]byte("こんにちは、\t\u0085世界"), []byte("こんにちは、 世界")},
		{[]byte("こんにちは、\u00a0\t世界"), []byte("こんにちは、 世界")},
		{[]byte("こんにちは、\u00a0\u0085世界"), []byte("こんにちは、 世界")},
		{[]byte("こんにちは、世界\t\n"), []byte("こんにちは、世界 ")},
		{[]byte("\t\nこんにちは、世界"), []byte(" こんにちは、世界")},
		{[]byte("\u00a0\nこんにちは、世界"), []byte(" こんにちは、世界")},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("compressSpaces(%q)", test.b)
		got, _ := compressSpaces(test.b)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("%s = %q, want %q", descr, got, test.want)
		}
	}
}
