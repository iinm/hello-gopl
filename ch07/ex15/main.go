package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"./eval"
)

func main() {
	input := bufio.NewScanner(os.Stdin)

	fmt.Print("expression: ")
	input.Scan()
	exprStr := input.Text()

	expr, err := eval.Parse(exprStr)
	if err != nil {
		panic(err)
	}
	fmt.Printf("parsed: %s\n", expr)

	vars := make(map[eval.Var]bool)
	if err := expr.Check(vars); err != nil {
		panic(err)
	}

	env := eval.Env{}
	for v, _ := range vars {
		fmt.Printf("%s = ", v)
		input.Scan()
		value, err := strconv.ParseFloat(input.Text(), 64)
		if err != nil {
			panic(err)
		}
		env[v] = value
	}

	result := expr.Eval(env)
	fmt.Printf("result = %g\n", result)
}
