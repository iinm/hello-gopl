package lengthconv

import (
	"fmt"
	"math"
	"testing"
)

const epsilon = 0.001 // 許容する誤差

func TestMToFT(t *testing.T) {
	var tests = []struct {
		m    Metre
		want Foot
	}{
		{1, 3.280},
		{3, 9.842},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("MToFT(%v)", test.m)
		got := MToFT(test.m)
		if math.Abs(float64(got-test.want)) > epsilon {
			t.Errorf("%s = %v, want %v", descr, got, test.want)
		}
	}
}

func TestFTToM(t *testing.T) {
	var tests = []struct {
		ft   Foot
		want Metre
	}{
		{3.280, 1},
		{9.842, 3},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("FTToM(%v)", test.ft)
		got := FTToM(test.ft)
		if math.Abs(float64(got-test.want)) > epsilon {
			t.Errorf("%s = %v, want %v", descr, got, test.want)
		}
	}
}
