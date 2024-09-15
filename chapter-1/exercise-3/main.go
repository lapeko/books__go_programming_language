package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

import (
	"log"
	"net/http"
	"time"
)

var palette = []color.Color{
	color.RGBA{G: 0xFF, A: 0xFF},
	color.White,
	color.Black,
	color.RGBA{R: 0xFF, A: 0xFF},
	color.RGBA{B: 0xFF, A: 0xFF},
	color.RGBA{R: 0xFF, B: 0xFF, A: 0xFF},
}

const (
	greenIndex = 0
	whiteIndex = 1
	blackIndex = 2
	redIndex   = 3
	blueIndex  = 4
	pinkIndex  = 5
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	if len(os.Args) > 1 && os.Args[1] == "web" {
		handler := func(w http.ResponseWriter, r *http.Request) {
			lissajous(w)
		}
		http.HandleFunc("/", handler)
		log.Fatal(http.ListenAndServe("localhost:8000", nil))
		return
	}
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			var colorIndex int
			switch {
			case t < 2*math.Pi:
				colorIndex = whiteIndex
			case t < 2*2*math.Pi:
				colorIndex = blackIndex
			case t < 3*2*math.Pi:
				colorIndex = redIndex
			case t < 4*2*math.Pi:
				colorIndex = blueIndex
			default:
				colorIndex = pinkIndex
			}
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				uint8(colorIndex))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
