package main

import (
	"flag"
	"fmt"

	"./tempconv"
)

// *celsiusFlagはflag.Valueインタフェースを満足する
type celsiusFlag struct{ tempconv.Celsius }

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit)
	switch unit {
	case "C", "°C":
		f.Celsius = tempconv.Celsius(value)
		return nil
	case "F", "°F":
		f.Celsius = tempconv.FToC(tempconv.Fahrenheit(value))
		return nil
	case "K", "°K":
		f.Celsius = tempconv.KToC(tempconv.Kelvin(value))
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}

//func (f *celsiusFlag) String() string {
//	return f.Celsius.String() + "##test##"
//}

func CelsiusFlag(name string, value tempconv.Celsius, usage string) *tempconv.Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}

var temp = CelsiusFlag("temp", 20.0, "the temperature")

func main() {
	flag.Parse()
	fmt.Println(*temp)
}
