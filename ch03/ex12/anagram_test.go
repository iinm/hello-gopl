package main

import (
	"fmt"
	"testing"
)

func TestAnagram(t *testing.T) {
	var tests = []struct {
		s1, s2 string
		want   bool
	}{
		{"", "", false},
		{"anagrams", "ars magna", true},
		{"ars magna", "anagrams", true},
		{"ars magna", "anagram", false},
		{"anagram", "ars magna", false},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("isAnagram(%s, %s)", test.s1, test.s2)
		got := isAnagran(test.s1, test.s2)
		if got != test.want {
			t.Errorf("%s = %v, want %v", descr, got, test.want)
		}
	}
}
