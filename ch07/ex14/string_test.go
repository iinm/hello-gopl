package eval

import (
	"fmt"
	"math"
	"testing"
)

func TestString(t *testing.T) {
	tests := []struct {
		exprStr string
		env     Env
	}{
		{"sqrt(A / pi)", Env{"A": 87616, "pi": math.Pi}},
		{"pow(x, 3) + pow(y, 3)", Env{"x": 12, "y": 1}},
		{"5 / 9 * (F - 32)", Env{"F": -40}},
		{"|x * 2|", Env{"x": -1}},
	}
	for _, test := range tests {
		fmt.Printf("Parse(%q) = ", test.exprStr)
		parsedExpr, err := Parse(test.exprStr)
		if err != nil {
			t.Error(err)
			continue
		}
		fmt.Printf("%q, ", parsedExpr.String())

		fmt.Printf("Parse(%q) = ", parsedExpr.String())
		reparsedExpr, err := Parse(parsedExpr.String())
		if err != nil {
			t.Error(err)
		}
		fmt.Printf("%q\n", reparsedExpr.String())

		if parsedExpr.String() != reparsedExpr.String() {
			t.Errorf("Parse(Parse(%q).String()) = %q, want %q\n",
				test.exprStr, reparsedExpr.String(), parsedExpr.String())
		}
	}
}
