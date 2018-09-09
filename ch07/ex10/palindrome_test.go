package palindrom

import (
	"fmt"
	"sort"
	"testing"
)

type runes []rune

func (rs runes) Len() int           { return len(rs) }
func (rs runes) Less(i, j int) bool { return rs[i] < rs[j] }
func (rs runes) Swap(i, j int)      { rs[i], rs[j] = rs[j], rs[i] }

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		s    sort.Interface
		want bool
	}{
		{runes(""), true},
		{runes("あいうえお"), false},
		{runes("あいうえおえういあ"), true},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("IsPandrome(%q)", test.s)
		got := IsPandrome(test.s)
		if got != test.want {
			t.Errorf("%s = %v, want %v", descr, got, test.want)
		}
	}
}
