package main

import (
	"bytes"
	"fmt"
	"reflect"
	"testing"
)

var args = makeTestArgs(10000)

func benchmarkSlowEcho(b *testing.B, size int) {
	out = new(bytes.Buffer)
	for i := 0; i < b.N; i++ {
		slowEcho(args[:size])
	}
}

func benchmarkEcho(b *testing.B, size int) {
	out = new(bytes.Buffer)
	for i := 0; i < b.N; i++ {
		echo(args[:size])
	}
}

func BenchmarkSlowEcho10(b *testing.B)    { benchmarkSlowEcho(b, 10) }
func BenchmarkSlowEcho100(b *testing.B)   { benchmarkSlowEcho(b, 100) }
func BenchmarkSlowEcho1000(b *testing.B)  { benchmarkSlowEcho(b, 1000) }
func BenchmarkSlowEcho10000(b *testing.B) { benchmarkSlowEcho(b, 10000) }
func BenchmarkEcho10(b *testing.B)        { benchmarkEcho(b, 10) }
func BenchmarkEcho100(b *testing.B)       { benchmarkEcho(b, 100) }
func BenchmarkEcho1000(b *testing.B)      { benchmarkEcho(b, 1000) }
func BenchmarkEcho10000(b *testing.B)     { benchmarkEcho(b, 10000) }

func makeTestArgs(length int) []string {
	testArgs := []string{}
	for i := 0; i < length; i++ {
		testArgs = append(testArgs, "foo")
	}
	return testArgs
}

func TestMakeTestArgs(t *testing.T) {
	var tests = []struct {
		length int
		want   []string
	}{
		{0, []string{}},
		{1, []string{"foo"}},
		{3, []string{"foo", "foo", "foo"}},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("makeTestArgs(%d)", test.length)
		got := makeTestArgs(test.length)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("%s = %q, want %q", descr, got, test.want)
		}
	}
}
