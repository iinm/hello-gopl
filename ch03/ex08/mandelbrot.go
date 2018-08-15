// マンデルブロフラクタルのPNG画像を生成する
package main

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"math/big"
	"math/cmplx"
	"os"
)

const (
	//windowSize    = .00001 // 64, 128で見た目に差異が出る倍率
	//xmin, ymin    = -.5430, -.6157
	windowSize    = .0000000000001
	xmin, ymin    = -.54300149349598, -.61570054194875
	xmax, ymax    = xmin + windowSize, ymin + windowSize
	width, height = 1024, 1024
)

var out io.Writer = os.Stdout

func main() {
	//img := makeImageComplex64()
	//img := makeImageComplex128()
	img := makeImageComplexBigFloat()
	png.Encode(out, img)
}

func makeImageComplex64() *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// 画像の点 (px, py) は複素数値zを表している
			img.Set(px, py, mandelbrotComplex64(complex64(z)))
		}
	}
	return img
}

func mandelbrotComplex64(z complex64) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex64
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(complex128(v)) > 2 {
			return color.YCbCr{128, 255 - contrast*n, 0}
		}
	}
	return color.Black
}

func makeImageComplex128() *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// 画像の点 (px, py) は複素数値zを表している
			img.Set(px, py, mandelbrotComplex128(z))
		}
	}
	return img
}

func mandelbrotComplex128(z complex128) color.Color {
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

type ComplexBigFloat struct {
	real, imag *big.Float
}

func (c *ComplexBigFloat) Add(a, b *ComplexBigFloat) *ComplexBigFloat {
	c.real = new(big.Float).Add(a.real, b.real)
	c.imag = new(big.Float).Add(a.imag, b.imag)
	return c
}

func (c *ComplexBigFloat) Mul(a, b *ComplexBigFloat) *ComplexBigFloat {
	c.real = new(big.Float).Sub(
		new(big.Float).Mul(a.real, b.real),
		new(big.Float).Mul(a.imag, b.imag),
	)
	c.imag = new(big.Float).Add(
		new(big.Float).Mul(a.real, b.imag),
		new(big.Float).Mul(b.real, a.imag),
	)
	return c
}

func (c *ComplexBigFloat) Abs(a *ComplexBigFloat) *big.Float {
	return new(big.Float).Sqrt(
		new(big.Float).Add(
			new(big.Float).Mul(a.real, a.real),
			new(big.Float).Mul(a.imag, a.imag),
		),
	)
}

func makeImageComplexBigFloat() *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		//y := big.NewFloat(float64(py)/height*(ymax-ymin) + ymin)
		y := new(big.Float).Add(
			new(big.Float).Mul(
				new(big.Float).Quo(big.NewFloat(float64(py)), big.NewFloat(height)),
				big.NewFloat(ymax-ymin),
			),
			big.NewFloat(ymin),
		)
		for px := 0; px < width; px++ {
			//x := big.NewFloat(float64(px)/width*(xmax-xmin) + xmin)
			x := new(big.Float).Add(
				new(big.Float).Mul(
					new(big.Float).Quo(big.NewFloat(float64(px)), big.NewFloat(width)),
					big.NewFloat(xmax-xmin),
				),
				big.NewFloat(xmin),
			)
			z := &ComplexBigFloat{x, y}
			// 画像の点 (px, py) は複素数値zを表している
			img.Set(px, py, mandelbrotComplexBigFloat(z))
		}
	}
	return img
}

func mandelbrotComplexBigFloat(z *ComplexBigFloat) color.Color {
	const iterations = 200
	const contrast = 15

	v := &ComplexBigFloat{big.NewFloat(0), big.NewFloat(0)}
	for n := uint8(0); n < iterations; n++ {
		//v = v*v + z
		v = new(ComplexBigFloat).Add(
			new(ComplexBigFloat).Mul(v, v),
			z,
		)
		if new(ComplexBigFloat).Abs(v).Cmp(big.NewFloat(2)) > 0 {
			return color.YCbCr{128, 255 - contrast*n, 0}
		}
	}
	return color.Black
}
