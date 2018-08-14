// 3-Dg画面の関数のSVGレンダリングを計算する。
package main

import (
	"fmt"
	"io"
	"math"
	"os"
)

const (
	width, height = 600, 320            // キャンパスの大きさ（画素数）
	cells         = 100                 // 格子のます目の数
	xyrange       = 30.0                // 軸の範囲 (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // x単位 および y単位あたりの画素数
	zscale        = height * 0.4        // z単位あたりの画素数
	angle         = math.Pi / 6         // x, y軸の角度 (= 30度)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

// settings
var out io.Writer = os.Stdout
var targetFunction = cobbs

func main() {
	fmt.Fprintf(out, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, aOk := corner(i+1, j)
			bx, by, bOk := corner(i, j)
			cx, cy, cOk := corner(i, j+1)
			dx, dy, dOk := corner(i+1, j+1)
			if !aOk || !bOk || !cOk || !dOk {
				continue
			}
			fmt.Fprintf(out, "<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Fprintln(out, "</svg>")
}

func corner(i, j int) (sx float64, sy float64, ok bool) {
	// ます目の(i, j)のかどの点(x, y)を見つける
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// 面の高さzを計算する
	z := targetFunction(x, y)
	if math.IsInf(z, 0) || math.IsNaN(z) {
		return 0, 0, false
	}

	// (x, y, z)を2-D SVGキャンパス (sx, sy) へ等角的に投影
	sx = width/2 + (x-y)*cos30*xyscale
	sy = height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, true
}

func saddle(x, y float64) float64 {
	return 0.0015 * (y*y - x*x)
}

func cobbs(x, y float64) float64 {
	r1 := math.Hypot(x, 0)
	r2 := math.Hypot(0, y)
	return 0.05 * (math.Cos(r1) * math.Cos(r2))
}

func distSin(x, y float64) float64 {
	r := math.Hypot(x, y) // (0, 0)からの距離
	return math.Sin(r) / r
}

func distSqrt(x, y float64) float64 {
	r := math.Hypot(x, y) // (0, 0)からの距離
	return math.Sqrt(r) / r
}
