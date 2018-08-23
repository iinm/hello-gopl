package main

import (
	"fmt"
	"testing"
)

func TestExpand(t *testing.T) {
	tests := []struct {
		s    string
		f    func(string) string
		want string
	}{
		{"", func(s string) string { return "#" + s + "#" }, ""},
		{"$", func(s string) string { return "#" + s + "#" }, "$"},
		{"$$", func(s string) string { return "#" + s + "#" }, "$$"},
		{" $ $ ", func(s string) string { return "#" + s + "#" }, " $ $ "},
		{"$foo $bar", func(s string) string { return "#" + s + "#" }, "#foo# #bar#"},
		{" $foo $bar ", func(s string) string { return "#" + s + "#" }, " #foo# #bar# "},
		{"$foo$bar", func(s string) string { return "#" + s + "#" }, "#foo##bar#"},
		{"$foo$", func(s string) string { return "#" + s + "#" }, "#foo#$"},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("expand(%q, %T)", test.s, test.f)
		got := expand(test.s, test.f)
		if got != test.want {
			t.Errorf("%s = %q, want %q", descr, got, test.want)
		}
	}
}
