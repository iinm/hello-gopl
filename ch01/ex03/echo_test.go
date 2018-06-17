package main

import (
	"bytes"
	"fmt"
	"reflect"
	"testing"
)

var args = makeTestArgs(10000)

func benchmarkSlowEcho(b *testing.B, size int) {
	for i := 0; i < b.N; i++ {
		slowEcho(args[:size], new(bytes.Buffer))
	}
}

func benchmarkEcho(b *testing.B, size int) {
	for i := 0; i < b.N; i++ {
		echo(args[:size], new(bytes.Buffer))
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

func makeTestArgs(size int) []string {
	testArgs := make([]string, 0, size)
	for i := 0; i < size; i++ {
		testArgs = append(testArgs, "foo")
	}
	return testArgs
}

func TestMakeTestArgs(t *testing.T) {
	var tests = []struct {
		size int
		want []string
	}{
		{0, []string{}},
		{1, []string{"foo"}},
		{3, []string{"foo", "foo", "foo"}},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("makeTestArgs(%d)", test.size)
		got := makeTestArgs(test.size)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("%s = %q, want %q", descr, got, test.want)
		}
	}
}

// 学習のためのテスト: スライスのキャパシティを設定したほうが速いか？
func makeTestArgsNoCap(size int) []string {
	testArgs := []string{}
	for i := 0; i < size; i++ {
		testArgs = append(testArgs, "foo")
	}
	return testArgs
}

func BenchmarkMakeTestArgs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		makeTestArgs(1000)
	}
}

func BenchmarkMakeTestArgsNoCap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		makeTestArgsNoCap(1000)
	}
}
