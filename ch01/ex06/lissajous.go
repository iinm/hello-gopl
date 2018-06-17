package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

// https://colorhunt.co/palette/122094
var palette = []color.Color{
	color.RGBA{0x53, 0x52, 0x38, 0xff},
	color.RGBA{0x4b, 0xbb, 0x8b, 0xff},
	color.RGBA{0x6d, 0xda, 0xbe, 0xff},
	color.RGBA{0xc9, 0xff, 0xc7, 0xff},
}

const (
	bgIndex = 0
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	err := lissajous(os.Stdout)
	if err != nil {
		panic(err)
	}
}

func lissajous(out io.Writer) error {
	const (
		cycles  = 5     // 発振器xが完了する周回の回数
		res     = 0.001 // 回転の分解能
		size    = 100   // 画像キャンバスは[-size..+size]の範囲を扱う
		nframes = 64    // アニメーションフレーム数
		delay   = 8     // 10ms単位でのフレーム間の遅延
	)
	freq := rand.Float64() * 3.0 // 発振器yの相対周波数
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // 位相差
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			nCycle := int(t / (2 * math.Pi))
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			xt, yt := size+int(x*size+0.5), size+int(y*size+0.5)
			fgIndex := uint8(nCycle%3 + 1)
			for yi := yt; yi <= size*2; yi++ {
				img.SetColorIndex(xt, yi, fgIndex)
			}
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	return gif.EncodeAll(out, &anim)
}
