package equal

import (
	"math"
	"reflect"
)

const threshold = 1. / 1000000000.

func equal(xi, yi interface{}) bool {
	x := reflect.ValueOf(xi)
	y := reflect.ValueOf(yi)

	if !x.IsValid() || !y.IsValid() {
		return false
	}

	var xFloat, yFloat float64
	switch x.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		xFloat = float64(x.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		xFloat = float64(x.Uint())
	case reflect.Float32, reflect.Float64:
		xFloat = float64(x.Float())
	default:
		return false
	}

	switch y.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		yFloat = float64(y.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		yFloat = float64(y.Uint())
	case reflect.Float32, reflect.Float64:
		yFloat = float64(y.Float())
	default:
		return false
	}

	return math.Abs(xFloat-yFloat) < threshold
}
