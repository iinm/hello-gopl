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
var targetFunction = distSin

func main() {
	fmt.Fprintf(out, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	var i, j float64
	for i = 0; i < cells; i++ {
		for j = 0; j < cells; j++ {
			ax, ay, _, aOk := corner(i+1, j)
			bx, by, _, bOk := corner(i, j)
			cx, cy, _, cOk := corner(i, j+1)
			dx, dy, _, dOk := corner(i+1, j+1)
			if !aOk || !bOk || !cOk || !dOk {
				continue
			}

			// polygonの内部に凹凸があるか？
			top, bottom := false, false
			step := 0.1
			for di := step; di < 1.0; di += step {
				_, _, z1, _ := corner(i+di, j)
				_, _, z2, _ := corner(i+di, j+1)
				for dj := step; dj < 1.0; dj += step {
					_, _, z3, _ := corner(i, j+dj)
					_, _, z4, _ := corner(i+1, j+dj)
					_, _, z, _ := corner(i+di, j+dj)
					if z < z1 && z < z2 && z < z3 && z < z4 {
						bottom = true
						break
					} else if z > z1 && z > z2 && z > z3 && z > z4 {
						top = true
						break
					}
				}
			}

			color := "white"
			if bottom {
				color = "blue"
			} else if top {
				color = "red"
			}
			fmt.Fprintf(out, "<polygon points='%g,%g %g,%g %g,%g %g,%g' style='fill: %s'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, color)
		}
	}
	fmt.Fprintln(out, "</svg>")
}

func corner(i, j float64) (sx, sy, z float64, ok bool) {
	// ます目の(i, j)のかどの点(x, y)を見つける
	x := xyrange * (i/cells - 0.5)
	y := xyrange * (j/cells - 0.5)

	// 面の高さzを計算する
	z = targetFunction(x, y)
	if math.IsInf(z, 0) || math.IsNaN(z) {
		return 0, 0, 0, false
	}

	// (x, y, z)を2-D SVGキャンパス (sx, sy) へ等角的に投影
	sx = width/2 + (x-y)*cos30*xyscale
	sy = height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, z, true
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
