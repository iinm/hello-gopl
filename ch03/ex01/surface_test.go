package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestSurface(t *testing.T) {
	var tests = []struct {
		f            func(float64, float64) float64
		mapper       func(int, int) (float64, float64, bool)
		validPolygon bool
		descr        string
	}{
		// 修正後のcorner関数では、SVGに無効な座標が含まれないことを確認
		{distSin, corner, true, "main() with distSin and corner"},
		{distSqrt, corner, true, "main() with distSqrt and corner"},
		// 修正前のcorner関数では、SVGに無効な座標が含まれることを確認
		{distSin, invalidCorner, false, "main() with distSin and invalidCorner"},
		{distSqrt, invalidCorner, false, "main() with distSqrt and invalidCorner"},
	}

	for _, test := range tests {
		out = new(bytes.Buffer)
		targetFunction = test.f
		mapper = test.mapper

		main()
		got := out.(*bytes.Buffer).String()
		if !strings.Contains(got, "NaN") != test.validPolygon {
			t.Errorf("%s output invalid polygon", test.descr)
		}
	}
}

func invalidCorner(i, j int) (sx float64, sy float64, ok bool) {
	// ます目の(i, j)のかどの点(x, y)を見つける
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// 面の高さzを計算する
	z := targetFunction(x, y)

	// (x, y, z)を2-D SVGキャンパス (sx, sy) へ等角的に投影
	sx = width/2 + (x-y)*cos30*xyscale
	sy = height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, true
}
