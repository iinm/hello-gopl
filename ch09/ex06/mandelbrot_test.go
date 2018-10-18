package main

import (
	"bytes"
	"testing"
)

var out = &bytes.Buffer{}

func BenchmarkRender(b *testing.B) {
	for i := 0; i < b.N; i++ {
		render(out)
	}
}

func BenchmarkRenderWithGoRoutine(b *testing.B) {
	for i := 0; i < b.N; i++ {
		renderWithGoRoutine(out)
	}
}

func BenchmarkRenderWithGoRoutine2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		renderWithGoRoutine2(out)
	}
}
