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

func TestEval(t *testing.T) {
	tests := []struct {
		expr Expr
		env  Env
		want string
	}{
		{
			call{"sqrt", []Expr{binary{'/', Var("A"), Var("pi")}}},
			Env{"A": 87616, "pi": math.Pi},
			"167",
		},
		{
			binary{
				'+',
				call{"pow", []Expr{Var("x"), literal(3)}},
				call{"pow", []Expr{Var("y"), literal(3)}},
			},
			Env{"x": 12, "y": 1},
			"1729",
		},
		{
			binary{
				'+',
				call{"pow", []Expr{Var("x"), literal(3)}},
				call{"pow", []Expr{Var("y"), literal(3)}},
			},
			Env{"x": 9, "y": 10},
			"1729",
		},
		{
			binary{
				'*',
				binary{'/', literal(5), literal(9)},
				binary{'-', Var("F"), literal(32)},
			},
			Env{"F": 32},
			"0",
		},
	}
	for _, test := range tests {
		fmt.Printf("\n%s\n", test.expr)
		got := fmt.Sprintf("%.6g", test.expr.Eval(test.env))
		fmt.Printf("\t%v => %s\n", test.env, got)
		if got != test.want {
			t.Errorf("%s.Eval() in %v = %q, want %q\n",
				test.expr, test.env, got, test.want)
		}
	}
}
