package popcount

import (
	"fmt"
	"testing"
)

const x = 2 ^ 64 - 1

var result int

func BenchmarkPopCount(b *testing.B) {
	// 戻り値を使わないとコンパイラの最適化により実行されない
	total := 0
	for i := 0; i < b.N; i++ {
		total += PopCount(x)
	}
	result = total
}

func BenchmarkPopCountByLoop(b *testing.B) {
	total := 0
	for i := 0; i < b.N; i++ {
		total += PopCountByLoop(x)
	}
	result = total
}

func BenchmarkPopCountByShift(b *testing.B) {
	total := 0
	for i := 0; i < b.N; i++ {
		total += PopCountByShift(x)
	}
	result = total
}

func BenchmarkPopCountByClear(b *testing.B) {
	total := 0
	for i := 0; i < b.N; i++ {
		total += PopCountByClear(x)
	}
	result = total
}

func TestPopCountByShift(t *testing.T) {
	var tests = []struct {
		x    uint64
		want int
	}{
		{0, 0},
		{255, 8},
		{256, 1},
		{511, 9},
		{512, 1},
		{18446744073709551615, 64},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("PopCountByShift(%v)", test.x)
		got := PopCountByShift(test.x)
		if got != test.want {
			t.Errorf("%s = %v, want %v", descr, got, test.want)
		}
	}
}

func TestPopCountByLoop(t *testing.T) {
	var tests = []struct {
		x    uint64
		want int
	}{
		{0, 0},
		{255, 8},
		{256, 1},
		{511, 9},
		{512, 1},
		{18446744073709551615, 64},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("PopCountByLoop(%v)", test.x)
		got := PopCountByLoop(test.x)
		//got := PopCount(test.x)
		if got != test.want {
			t.Errorf("%s = %v, want %v", descr, got, test.want)
		}
	}
}

func TestPopCountByClear(t *testing.T) {
	var tests = []struct {
		x    uint64
		want int
	}{
		{0, 0},
		{255, 8},
		{256, 1},
		{511, 9},
		{512, 1},
		{18446744073709551615, 64},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("PopCountByClear(%v)", test.x)
		got := PopCountByClear(test.x)
		if got != test.want {
			t.Errorf("%s = %v, want %v", descr, got, test.want)
		}
	}
}
