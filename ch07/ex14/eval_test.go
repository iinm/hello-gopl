package eval

import (
	"fmt"
	"math"
	"testing"
)

func TestParseAndEval(t *testing.T) {
	tests := []struct {
		expr string
		env  Env
		want string
	}{
		{"sqrt(A / pi)", Env{"A": 87616, "pi": math.Pi}, "167"},
		{"pow(x, 3) + pow(y, 3)", Env{"x": 12, "y": 1}, "1729"},
		{"pow(x, 3) + pow(y, 3)", Env{"x": 9, "y": 10}, "1729"},
		{"5 / 9 * (F - 32)", Env{"F": -40}, "-40"},
		{"5 / 9 * (F - 32)", Env{"F": 32}, "0"},
		{"5 / 9 * (F - 32)", Env{"F": 212}, "100"},
		{"-1 + -x", Env{"x": 1}, "-2"},
		{"-1 - x", Env{"x": 1}, "-2"},
		{"|x|", Env{"x": 1}, "1"},
		{"|x|", Env{"x": -1}, "1"},
		{"|x - 2|", Env{"x": -1}, "3"},
		{"|x - 2| * 2", Env{"x": -1}, "6"},
	}
	var prevExpr string
	for _, test := range tests {
		if test.expr != prevExpr {
			fmt.Printf("\n%s\n", test.expr)
			prevExpr = test.expr
		}
		expr, err := Parse(test.expr)
		if err != nil {
			t.Error(err)
			continue
		}
		got := fmt.Sprintf("%.6g", expr.Eval(test.env))
		fmt.Printf("\t%v => %s\n", test.env, got)
		if got != test.want {
			t.Errorf("%s.Eval() in %v = %q, want %q\n",
				test.expr, test.env, got, test.want)
		}
	}
}

//func TestEval(t *testing.T) {
//	tests := []struct {
//		expr Expr
//		env  Env
//		want string
//	}{
//		{
//			call{"sqrt", []Expr{binary{'/', Var("A"), Var("pi")}}},
//			Env{"A": 87616, "pi": math.Pi},
//			"167",
//		},
//		{
//			binary{
//				'+',
//				call{"pow", []Expr{Var("x"), literal(3)}},
//				call{"pow", []Expr{Var("y"), literal(3)}},
//			},
//			Env{"x": 12, "y": 1},
//			"1729",
//		},
//		{
//			binary{
//				'+',
//				call{"pow", []Expr{Var("x"), literal(3)}},
//				call{"pow", []Expr{Var("y"), literal(3)}},
//			},
//			Env{"x": 9, "y": 10},
//			"1729",
//		},
//		{
//			binary{
//				'*',
//				binary{'/', literal(5), literal(9)},
//				binary{'-', Var("F"), literal(32)},
//			},
//			Env{"F": 32},
//			"0",
//		},
//	}
//	for _, test := range tests {
//		fmt.Printf("\n%s\n", test.expr)
//		got := fmt.Sprintf("%.6g", test.expr.Eval(test.env))
//		fmt.Printf("\t%v => %s\n", test.env, got)
//		if got != test.want {
//			t.Errorf("%s.Eval() in %v = %q, want %q\n",
//				test.expr, test.env, got, test.want)
//		}
//	}
//}
