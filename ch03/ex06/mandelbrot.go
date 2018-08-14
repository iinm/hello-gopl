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
	samplingHeight := 0.5 / height * (ymax - ymin)
	samplingWidth := 0.5 / width * (xmax - xmin)
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			// supersampling
			samples := []color.YCbCr{
				mandelbrot(complex(x, y)),
				mandelbrot(complex(x+samplingWidth, y)),
				mandelbrot(complex(x+samplingWidth, y+samplingHeight)),
				mandelbrot(complex(x, y+samplingHeight)),
			}
			var sumY, sumCb, sumCr uint32
			for _, sample := range samples {
				sumY += uint32(sample.Y)
				sumCb += uint32(sample.Cb)
				sumCr += uint32(sample.Cr)
			}
			sampleSize := uint32(len(samples))
			color := color.YCbCr{
				uint8(sumY / sampleSize),
				uint8(sumCb / sampleSize),
				uint8(sumCr / sampleSize),
			}
			// 画像の点 (px, py) は複素数値zを表している
			img.Set(px, py, color)
		}
	}
	png.Encode(os.Stdout, img) // 注意: エラーを無視
}

func mandelbrot(z complex128) color.YCbCr {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			//return color.Gray{255 - contrast*n}
			return color.YCbCr{128, 255 - contrast*n, 0}
			//return color.YCbCr{128, 255 - contrast*n, 255 - contrast*n}
		}
	}
	return color.YCbCr{0, 128, 128}
}
