package tempconv

import (
	"fmt"
	"math"
	"testing"
)

const epsilon = 0.001 // 許容する誤差

func TestKToC(t *testing.T) {
	var tests = []struct {
		k    Kelvin
		want Celsius
	}{
		{0, -273.15},    // 絶対零度
		{183.95, -89.2}, // 地球表面の最低気温
	}

	for _, test := range tests {
		descr := fmt.Sprintf("KToC(%v)", test.k)
		got := KToC(test.k)
		if math.Abs(float64(got-test.want)) > epsilon {
			t.Errorf("%s = %v, want %v", descr, got, test.want)
		}
	}
}

func TestKToF(t *testing.T) {
	var tests = []struct {
		k    Kelvin
		want Fahrenheit
	}{
		{0, -459.67},      // 絶対零度
		{183.95, -128.56}, // 地球表面の最低気温
	}

	for _, test := range tests {
		descr := fmt.Sprintf("KToF(%v)", test.k)
		got := KToF(test.k)
		if math.Abs(float64(got-test.want)) > epsilon {
			t.Errorf("%s = %v, want %v", descr, got, test.want)
		}
	}
}

func TestCToF(t *testing.T) {
	var tests = []struct {
		c    Celsius
		want Fahrenheit
	}{
		{-273.15, -459.67}, // 絶対零度
		{-89.2, -128.56},   // 地球表面の最低気温
	}

	for _, test := range tests {
		descr := fmt.Sprintf("CToF(%v)", test.c)
		got := CToF(test.c)
		if math.Abs(float64(got-test.want)) > epsilon {
			t.Errorf("%s = %v, want %v", descr, got, test.want)
		}
	}
}

func TestCToK(t *testing.T) {
	var tests = []struct {
		c    Celsius
		want Kelvin
	}{
		{-273.15, 0},    // 絶対零度
		{-89.2, 183.95}, // 地球表面の最低気温
	}

	for _, test := range tests {
		descr := fmt.Sprintf("CToK(%v)", test.c)
		got := CToK(test.c)
		if math.Abs(float64(got-test.want)) > epsilon {
			t.Errorf("%s = %v, want %v", descr, got, test.want)
		}
	}
}

func TestFToC(t *testing.T) {
	var tests = []struct {
		f    Fahrenheit
		want Celsius
	}{
		{-459.67, -273.15}, // 絶対零度
		{-128.56, -89.2},   // 地球表面の最低気温
	}

	for _, test := range tests {
		descr := fmt.Sprintf("CToF(%v)", test.f)
		got := FToC(test.f)
		if math.Abs(float64(got-test.want)) > epsilon {
			t.Errorf("%s = %v, want %v", descr, got, test.want)
		}
	}
}

func TestFToK(t *testing.T) {
	var tests = []struct {
		f    Fahrenheit
		want Kelvin
	}{
		{-459.67, 0},      // 絶対零度
		{-128.56, 183.95}, // 地球表面の最低気温
	}

	for _, test := range tests {
		descr := fmt.Sprintf("KToF(%v)", test.f)
		got := FToK(test.f)
		if math.Abs(float64(got-test.want)) > epsilon {
			t.Errorf("%s = %v, want %v", descr, got, test.want)
		}
	}
}
