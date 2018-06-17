package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

var palette = []color.Color{color.Black, color.RGBA{0x00, 0xff, 0x00, 0xff}}

const (
	bgIndex = 0
	fgIndex = 1
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	http.HandleFunc("/", lissajousHandler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func lissajousHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// default params
	var (
		cycles  = 5.0   // 発振器xが完了する周回の回数
		res     = 0.001 // 回転の分解能
		size    = 100   // 画像キャンバスは[-size..+size]の範囲を扱う
		nframes = 64    // アニメーションフレーム数
		delay   = 8     // 10ms単位でのフレーム間の遅延
	)
	if r.Form.Get(`cycles`) != `` {
		if formCycles, err := strconv.ParseFloat(r.Form.Get(`cycles`), 64); err == nil {
			cycles = formCycles
		} else {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "invalid parameter: cycles = %v", r.Form.Get(`cycles`))
			return
		}
	}
	if r.Form.Get(`res`) != `` {
		if formRes, err := strconv.ParseFloat(r.Form.Get(`res`), 64); err == nil {
			res = formRes
		} else {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "invalid parameter: res = %v", r.Form.Get(`res`))
			return
		}
	}
	if r.Form.Get(`size`) != `` {
		if formSize, err := strconv.Atoi(r.Form.Get(`size`)); err == nil {
			size = formSize
		} else {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "invalid parameter: size = %v", r.Form.Get(`size`))
			return
		}
	}
	if r.Form.Get(`nframes`) != `` {
		if formNframes, err := strconv.Atoi(r.Form.Get(`nframes`)); err == nil {
			nframes = formNframes
		} else {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "invalid parameter: nframes = %v", r.Form.Get(`nframes`))
			return
		}
	}
	if r.Form.Get(`delay`) != `` {
		if formDelay, err := strconv.Atoi(r.Form.Get(`delay`)); err == nil {
			delay = formDelay
		} else {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "invalid parameter: delay = %v", r.Form.Get(`delay`))
			return
		}
	}

	if err := lissajous(w, cycles, res, size, nframes, delay); err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "internal server error")
	}
}

func lissajous(out io.Writer, cycles float64, res float64, size int, nframes int, delay int) error {
	freq := rand.Float64() * 3.0 // 発振器yの相対周波数
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // 位相差
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*float64(size)+0.5), size+int(y*float64(size)+0.5), fgIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	return gif.EncodeAll(out, &anim)
}
