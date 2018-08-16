package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestReverseUnicodeBytes(t *testing.T) {
	var tests = []struct {
		b    []byte
		want []byte
	}{
		{[]byte("こんにちは、世界"), []byte("界世、はちにんこ")},
		{[]byte("こんにちは、世界abc"), []byte("cba界世、はちにんこ")},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("reverseUnicodeBytes(%q)", test.b)
		got, _ := reverseUnicodeBytes(test.b)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("%s = %q, want %q", descr, got, test.want)
		}
	}
}
