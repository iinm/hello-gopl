// マンデルブロフラクタルのPNG画像を生成する
package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// 画像の点 (px, py) は複素数値zを表している
			img.Set(px, py, newton(z))
		}
	}
	png.Encode(os.Stdout, img) // 注意: エラーを無視
}

func next(z complex128) complex128 {
	// z^4 - 1 = 0
	// z - f(z) / f'(z)
	return z - (z*z*z*z-1)/(4*z*z*z)
	// test: x^5 + x^2 +1 = 0
	// http://www.math.u-ryukyu.ac.jp/~suga/ssh4/node6.html
	//return z - (z*z*z*z*z+z*z+1)/(5*z*z*z*z+2*z)
}

func newton(z complex128) color.Color {
	const iterations = 200
	const contrast = 25

	epsilon := 1e-5
	answer := z
	for n := uint8(0); n < iterations; n++ {
		if cmplx.Abs(1-answer) < epsilon {
			return color.CMYK{n * contrast, 0x00, 0x00, 0x00}
		}
		if cmplx.Abs(-1-answer) < epsilon {
			return color.CMYK{0x00, n * contrast, 0x00, 0x00}
		}
		if cmplx.Abs(1i-answer) < epsilon {
			return color.CMYK{0x00, n * contrast, n * contrast, 0x00}
		}
		if cmplx.Abs(-1i-answer) < epsilon {
			return color.CMYK{n * contrast, n * contrast, 0x00, 0x00}
		}
		answer = next(answer)
	}
	return color.Black
}
