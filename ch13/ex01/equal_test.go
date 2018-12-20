package equal

import "testing"

func TestEqual(t *testing.T) {
	tests := []struct {
		x, y interface{}
		want bool
	}{
		{1, 1. + threshold, false},
		{1, 1. + threshold/2, true},
	}
	for _, test := range tests {
		got := equal(test.x, test.y)
		if got != test.want {
			t.Errorf("equal(%v, %v) = %v, want %v", test.x, test.y, got, test.want)
		}
	}
}
