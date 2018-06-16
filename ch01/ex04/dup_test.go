package main

import (
	"fmt"
	"io"
	"reflect"
	"strings"
	"testing"
)

var testFileContent = strings.TrimSpace(`
foo
bar
baz
foo
baz
`)

func TestCountLines(t *testing.T) {
	var tests = []struct {
		f               io.Reader
		filename        string
		counts          map[string]int
		locations       map[string][]string
		wantedCounts    map[string]int
		wantedLocations map[string][]string
	}{
		{
			strings.NewReader(testFileContent),
			"test.txt",
			make(map[string]int),
			make(map[string][]string),
			map[string]int{"foo": 2, "bar": 1, "baz": 2},
			map[string][]string{
				"foo": []string{"test.txt:1", "test.txt:4"},
				"bar": []string{"test.txt:2"},
				"baz": []string{"test.txt:3", "test.txt:5"},
			},
		},
		{
			strings.NewReader(testFileContent),
			"test.txt",
			map[string]int{"foo": 1, "bar": 1, "baz": 1},
			map[string][]string{
				"foo": []string{"dummy.txt:1"},
				"bar": []string{"dummy.txt:2"},
				"baz": []string{"dummy.txt:3"},
			},
			map[string]int{"foo": 3, "bar": 2, "baz": 3},
			map[string][]string{
				"foo": []string{"dummy.txt:1", "test.txt:1", "test.txt:4"},
				"bar": []string{"dummy.txt:2", "test.txt:2"},
				"baz": []string{"dummy.txt:3", "test.txt:3", "test.txt:5"},
			},
		},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("countLines(%v, %s, %v, %v)", test.f, test.filename, test.counts, test.locations)
		countLines(test.f, test.filename, test.counts, test.locations)
		if !reflect.DeepEqual(test.counts, test.wantedCounts) {
			t.Errorf("%s => counts = %v, want %v", descr, test.counts, test.wantedCounts)
		}
		if !reflect.DeepEqual(test.locations, test.wantedLocations) {
			t.Errorf("%s => locations = %v, want %v", descr, test.locations, test.wantedLocations)
		}
	}
}
