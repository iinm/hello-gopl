package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

const (
	//width, height = 600, 320            // キャンパスの大きさ（画素数）
	cells   = 100  // 格子のます目の数
	xyrange = 30.0 // 軸の範囲 (-xyrange..+xyrange)
	//xyscale       = width / 2 / xyrange // x単位 および y単位あたりの画素数
	//zscale = height * 0.4 // z単位あたりの画素数
	angle = math.Pi / 6 // x, y軸の角度 (= 30度)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

// settings
var targetFunction = distSin

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	http.HandleFunc("/", surfaceHandler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func surfaceHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	width, height := 600, 320 // default
	if r.Form.Get(`width`) != `` {
		if formWidth, err := strconv.Atoi(r.Form.Get(`width`)); err == nil {
			width = formWidth
		} else {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "invalid parameter: width = %v", r.Form.Get(`width`))
			return
		}
	}
	if r.Form.Get(`height`) != `` {
		if formHeight, err := strconv.Atoi(r.Form.Get(`height`)); err == nil {
			height = formHeight
		} else {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "invalid parameter: height = %v", r.Form.Get(`height`))
			return
		}
	}

	w.Header().Set("Content-Type", "image/svg+xml")
	surface(w, width, height)
}

func surface(out io.Writer, width, height int) {
	fmt.Fprintf(out, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	var i, j float64
	for i = 0; i < cells; i++ {
		for j = 0; j < cells; j++ {
			ax, ay, _, aOk := corner(i+1, j, width, height)
			bx, by, _, bOk := corner(i, j, width, height)
			cx, cy, _, cOk := corner(i, j+1, width, height)
			dx, dy, _, dOk := corner(i+1, j+1, width, height)
			if !aOk || !bOk || !cOk || !dOk {
				continue
			}

			// polygonの内部に凹凸があるか？
			top, bottom := false, false
			step := 0.1
			for di := step; di < 1.0; di += step {
				_, _, z1, _ := corner(i+di, j, width, height)
				_, _, z2, _ := corner(i+di, j+1, width, height)
				for dj := step; dj < 1.0; dj += step {
					_, _, z3, _ := corner(i, j+dj, width, height)
					_, _, z4, _ := corner(i+1, j+dj, width, height)
					_, _, z, _ := corner(i+di, j+dj, width, height)
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

func corner(i, j float64, width, height int) (sx, sy, z float64, ok bool) {
	// ます目の(i, j)のかどの点(x, y)を見つける
	x := xyrange * (i/cells - 0.5)
	y := xyrange * (j/cells - 0.5)

	// 面の高さzを計算する
	z = targetFunction(x, y)
	if math.IsInf(z, 0) || math.IsNaN(z) {
		return 0, 0, 0, false
	}

	// (x, y, z)を2-D SVGキャンパス (sx, sy) へ等角的に投影
	xyscale := float64(width) / 2 / xyrange
	zscale := float64(height) * 0.4
	sx = float64(width)/2 + (x-y)*cos30*xyscale
	sy = float64(height)/2 + (x+y)*sin30*xyscale - z*zscale
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
