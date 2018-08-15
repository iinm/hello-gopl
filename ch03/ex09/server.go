package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"
	"log"
	"math/cmplx"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	http.HandleFunc("/", mandelbrotHandler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func mandelbrotHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	x, y, zoomLevel := 0., 0., 1.
	if r.Form.Get(`x`) != `` {
		if formX, err := strconv.ParseFloat(r.Form.Get(`x`), 64); err == nil {
			x = formX
		} else {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "invalid parameter: x = %v", r.Form.Get(`x`))
			return
		}
	}
	if r.Form.Get(`y`) != `` {
		if formY, err := strconv.ParseFloat(r.Form.Get(`y`), 64); err == nil {
			y = formY
		} else {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "invalid parameter: y = %v", r.Form.Get(`y`))
			return
		}
	}
	if r.Form.Get(`zoomLevel`) != `` {
		if formZoomLevel, err := strconv.ParseFloat(r.Form.Get(`zoomLevel`), 64); err == nil {
			zoomLevel = formZoomLevel
		} else {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "invalid parameter: zoomLevel = %v", r.Form.Get(`zoomLevel`))
			return
		}
	}

	writeImage(w, x, y, zoomLevel)
}

func writeImage(out io.Writer, x, y, zoomLevel float64) {
	var (
		windowSize             = 2. / zoomLevel
		xmin, ymin, xmax, ymax = x - windowSize, y - windowSize, x + windowSize, y + windowSize
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/float64(height)*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/float64(width)*(xmax-xmin) + xmin
			z := complex(x, y)
			// 画像の点 (px, py) は複素数値zを表している
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(out, img) // 注意: エラーを無視
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.YCbCr{128, 255 - contrast*n, 0}
		}
	}
	return color.Black
}
