package massconv

import (
	"fmt"
	"math"
	"testing"
)

const epsilon = 0.001 // 許容する誤差

func TestKgToLB(t *testing.T) {
	var tests = []struct {
		kg   Kilogram
		want Pound
	}{
		{1, 2.20462},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("KgToLB(%v)", test.kg)
		got := KgToLB(test.kg)
		if math.Abs(float64(got-test.want)) > epsilon {
			t.Errorf("%s = %v, want %v", descr, got, test.want)
		}
	}
}

func TestFTToM(t *testing.T) {
	var tests = []struct {
		lb   Pound
		want Kilogram
	}{
		{2.20462, 1},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("LBToKg(%v)", test.lb)
		got := LBToKg(test.lb)
		if math.Abs(float64(got-test.want)) > epsilon {
			t.Errorf("%s = %v, want %v", descr, got, test.want)
		}
	}
}
