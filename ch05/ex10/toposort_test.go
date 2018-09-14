package main

import (
	"errors"
	"fmt"
	"testing"
)

// https://ja.wikipedia.org/wiki/トポロジカルソート#例
var wikipediaSampleDeps = map[string][]string{
	"2":  []string{"11"},
	"9":  []string{"11", "8"},
	"10": []string{"11", "3"},
	"11": []string{"7", "5"},
	"8":  []string{"7", "3"},
}

// todo: 非決定的なので何度かテストを回すようにすると良い
func TestTopoSort(t *testing.T) {
	tests := []struct {
		deps map[string][]string
	}{
		{wikipediaSampleDeps},
		// https://github.com/adonovan/gopl.io/blob/master/ch5/toposort/main.go
		{
			map[string][]string{
				"algorithms": {"data structures"},
				"calculus":   {"linear algebra"},
				//"linear algebra": {"calculus"},

				"compilers": {
					"data structures",
					"formal languages",
					"computer organization",
				},

				"data structures":       {"discrete math"},
				"databases":             {"data structures"},
				"discrete math":         {"intro to programming"},
				"formal languages":      {"discrete math"},
				"networks":              {"operating systems"},
				"operating systems":     {"data structures", "computer organization"},
				"programming languages": {"data structures", "computer organization"},
			},
		},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("topoSort(%v)", test.deps)
		got := topoSort(test.deps)
		if err := isTopological(got, test.deps); err != nil {
			t.Errorf("%s = %q => %q", descr, got, err)
		}
	}
}

// sortedがトポロジカル順になっているかdepsを使ってを検証する
func isTopological(sorted []string, deps map[string][]string) error {
	order := make(map[string]int)
	for i, item := range sorted {
		order[item] = i
	}

	for item, dependencies := range deps {
		for _, d := range dependencies {
			if order[item] < order[d] {
				return fmt.Errorf("not topological: %q -> %q", item, d)
			}
		}
	}

	return nil
}

func TestIsTopological(t *testing.T) {
	tests := []struct {
		sorted []string
		deps   map[string][]string
		want   error
	}{
		{
			[]string{"7", "5", "3", "11", "8", "2", "9", "10"},
			wikipediaSampleDeps,
			nil,
		},
		{
			[]string{"3", "5", "7", "8", "11", "2", "9", "10"},
			wikipediaSampleDeps,
			nil,
		},
		{
			[]string{"3", "5", "7", "8", "2", "11", "9", "10"},
			wikipediaSampleDeps,
			errors.New(`not topological: "2" -> "11"`),
		},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("isTopological(%v, %v)", test.sorted, test.deps)
		got := isTopological(test.sorted, test.deps)
		if test.want == nil && got != nil { // エラーがないことを期待する場合
			t.Errorf("%s = %v, want %v", descr, got, test.want)
		}
		if test.want != nil && got.Error() != test.want.Error() { // エラーを期待する場合
			t.Errorf("%s = %v, want %v", descr, got, test.want)
		}
	}
}
