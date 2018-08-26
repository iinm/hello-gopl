package main

import (
	"fmt"
	"os"
)

var deps = map[string][]string{
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
}

// os.Args[1]で指定したコースを受講するために必要なコースの受講順序を表示する。
func main() {
	dependencies := courseOrder(os.Args[1], deps)
	for i := len(dependencies) - 1; 0 <= i; i-- {
		fmt.Printf("%d\t%s\n", len(dependencies)-i, dependencies[i])
	}
}

func courseOrder(course string, deps map[string][]string) (dependencies []string) {
	seen := make(map[string]bool)

	visit := func(item string) []string {
		dependencies = append(dependencies, item)
		seen[item] = true
		var items []string
		for _, d := range deps[item] {
			items = append(items, d)
		}
		return items
	}

	breadthFirst(visit, []string{course})

	return dependencies
}

func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}
