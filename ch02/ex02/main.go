package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/iinm/hello-gopl/ch02/lengthconv"
	"github.com/iinm/hello-gopl/ch02/massconv"
	"github.com/iinm/hello-gopl/ch02/tempconv"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		var err error
		args, err = getArgs(os.Stdin)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: %v\n", os.Args[0], err)
			os.Exit(1)
		}
	}

	for _, arg := range args {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: %v\n", os.Args[0], err)
			os.Exit(1)
		}
		showConvertResults(t)
	}
}

func getArgs(reader io.Reader) ([]string, error) {
	bytes, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return strings.Fields(string(bytes)), nil
}

func showConvertResults(t float64) {
	f := tempconv.Fahrenheit(t)
	c := tempconv.Celsius(t)
	m := lengthconv.Metre(t)
	ft := lengthconv.Foot(t)
	kg := massconv.Kilogram(t)
	lb := massconv.Pound(t)
	fmt.Println("---")
	fmt.Printf("%s = %s, %s = %s\n", f, tempconv.FToC(f), c, tempconv.CToF(c))
	fmt.Printf("%s = %s, %s = %s\n", m, lengthconv.MToFT(m), ft, lengthconv.FTToM(ft))
	fmt.Printf("%s = %s, %s = %s\n", kg, massconv.KgToLB(kg), lb, massconv.LBToKg(lb))
}
