package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
	"runtime"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	defer func() {
		fmt.Fprintln(os.Stderr, time.Since(start))
	}()

	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	threadsNum := runtime.NumCPU()
	wg := &sync.WaitGroup{}
	heightCh := make(chan int)

	img := image.NewRGBA(image.Rect(0, 0, width, height))

	for i := 0; i < threadsNum; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for h := range heightCh {
				y := float64(h)/height*(ymax-ymin) + ymin
				for px := 0; px < width; px++ {
					x := float64(px)/width*(xmax-xmin) + xmin
					z := complex(x, y)
					img.Set(px, h, mandelbrot(z))
				}
			}
		}()
	}

	for i := 0; i < height; i++ {
		heightCh <- i
	}
	close(heightCh)

	wg.Wait()
	png.Encode(os.Stdout, img)
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15
	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.RGBA{
				(n * 2) % 255,
				(100 + contrast*n) % 255,
				(contrast * n) % 255,
				255,
			}
		}
	}
	return color.Black
}
