package main

import (
	"bytes"
	"testing"
)

func TestLissajous(t *testing.T) {
	err := lissajous(new(bytes.Buffer))
	if err != nil {
		t.Errorf("lissajous() = %q, want nil", err)
	}
}
