package eval

import (
	"bytes"
	"fmt"
	"math"
)

// Env は変数名を値へと対応付ける環境
type Env map[Var]float64

// Expr は算術式
type Expr interface {
	Eval(env Env) float64
	String() string
}

// Var は変数を特定する
type Var string

// literal は数値定数
type literal float64

// uranry は単行演算式
type unary struct {
	op rune // '+' or '-'
	x  Expr
}

// binary は二項演算子
type binary struct {
	op   rune // '+', '-', '*', '/'
	x, y Expr
}

// callは関数呼び出し式
type call struct {
	fn   string // "pow", "sin", "sqrt"
	args []Expr
}

func (v Var) Eval(env Env) float64 {
	return env[v]
}

func (l literal) Eval(_ Env) float64 {
	return float64(l)
}

func (u unary) Eval(env Env) float64 {
	switch u.op {
	case '+':
		return +u.x.Eval(env)
	case '-':
		return -u.x.Eval(env)
	}
	panic(fmt.Sprintf("unsupported unary operator: %q", u.op))
}

func (b binary) Eval(env Env) float64 {
	switch b.op {
	case '+':
		return b.x.Eval(env) + b.y.Eval(env)
	case '-':
		return b.x.Eval(env) - b.y.Eval(env)
	case '*':
		return b.x.Eval(env) * b.y.Eval(env)
	case '/':
		return b.x.Eval(env) / b.y.Eval(env)
	}
	panic(fmt.Sprintf("unsupported binary operator: %q", b.op))
}

func (c call) Eval(env Env) float64 {
	switch c.fn {
	case "pow":
		return math.Pow(c.args[0].Eval(env), c.args[1].Eval(env))
	case "sin":
		return math.Sin(c.args[0].Eval(env))
	case "sqrt":
		return math.Sqrt(c.args[0].Eval(env))
	}
	panic(fmt.Sprintf("unsupported function call: %q", c.fn))
}

func (v Var) String() string     { return string(v) }
func (l literal) String() string { return fmt.Sprintf("%.6g", l) }
func (u unary) String() string   { return fmt.Sprintf("%s%s", string(u.op), u.x.String()) }
func (b binary) String() string {
	return fmt.Sprintf("(%s %s %s)", b.x.String(), string(b.op), b.y.String())
}
func (c call) String() string {
	buf := new(bytes.Buffer)
	for i, arg := range c.args {
		if i > 0 {
			buf.WriteString(", ")
		}
		buf.WriteString(arg.String())
	}
	return fmt.Sprintf("%s(%s)", c.fn, buf.String())
}
