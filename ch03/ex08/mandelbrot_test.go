package main

import "testing"

func BenchmarkMakeImageComplex64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		makeImageComplex64()
	}
}

func BenchmarkMakeImageComplex128(b *testing.B) {
	for i := 0; i < b.N; i++ {
		makeImageComplex128()
	}
}

func BenchmarkMakeImageComplexBigFloat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		makeImageComplexBigFloat()
	}
}
