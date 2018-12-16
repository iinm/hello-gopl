package sexpr

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMarchal(t *testing.T) {
	type testInterface interface{}
	tests := []struct {
		v    interface{}
		want []byte
	}{
		{v: []int{1, 2, 3}, want: []byte("(1 2 3)")},
		{v: []bool{true, false}, want: []byte("(t nil)")},
		{v: []float64{0.1, 3.1415}, want: []byte("(0.1 3.1415)")},
		{v: complex(1, 2), want: []byte("#C(1 2)")},
		{v: struct{ i testInterface }{[]int{1, 2, 3}}, want: []byte(`((i ("sexpr.testInterface" '(1 2 3))))`)},
	}
	for _, test := range tests {
		descr := fmt.Sprintf("Marchal(%v)", test.v)
		got, err := Marshal(test.v)
		if err != nil || !reflect.DeepEqual(got, test.want) {
			t.Errorf("%s = (%q, %v), want (%q, nil)", descr, got, err, test.want)
		}
	}
}
