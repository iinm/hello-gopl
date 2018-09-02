package counter

import (
	"fmt"
	"testing"
)

func TestLineCounter(t *testing.T) {
	tests := []struct {
		ps   []string
		want int
	}{
		{[]string{""}, 0},
		{[]string{" "}, 1},
		{[]string{"test\n"}, 2},
		{[]string{"test\n", ""}, 2},
		{[]string{"test\nte", "st"}, 2},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("LineCounter.Write(%#v) -> LineCounter.Count()", test.ps)
		counter := &LineCounter{}
		for _, p := range test.ps {
			counter.Write([]byte(p))
		}
		if counter.Count() != test.want {
			t.Errorf("%s = %d, want %d", descr, counter.Count(), test.want)
		}
	}
}

func TestWordCounter(t *testing.T) {
	tests := []struct {
		ps   []string
		want int
	}{
		{[]string{""}, 0},
		{[]string{" "}, 0},
		{[]string{"Hello World!"}, 2},
		{[]string{"こんにちは　世界！"}, 2},
		{[]string{" Hello World! "}, 2},
		{[]string{"Hel", "lo ", "Wor", "ld! "}, 2},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("WordCounter.Write(%#v) -> WordCounter.Count()", test.ps)
		counter := &WordCounter{}
		for _, p := range test.ps {
			counter.Write([]byte(p))
		}
		if counter.Count() != test.want {
			t.Errorf("%s = %d, want %d", descr, counter.Count(), test.want)
		}
	}
}
