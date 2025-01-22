package main

import (
	"fmt"
	"io"
	"math"
)

const (
	width, height = 600, 320
	cells         = 100
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.4
	angle         = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func BuildPolygon(w io.Writer) {
	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)

			if hasInfinite(ax, ay, bx, by, cx, cy, dx, dy) {
				continue
			}

			fmt.Fprintf(w, "<polygon points='%g,%g %g,%g %g,%g %g,%g' fill='%s'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, getColor(i, j))
		}
	}

	fmt.Fprint(w, "</svg>")
}

func getXYZ(i, j int) (float64, float64, float64) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	return x, y, f(x, y)
}

func corner(i, j int) (float64, float64) {
	x, y, z := getXYZ(i, j)
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}

func getColor(i, j int) string {
	_, _, z := getXYZ(i, j)
	secondColor := 255 - int(math.Round(math.Abs(z)*255))
	r, g, b := 255, secondColor, 255
	if z > 0 {
		b = secondColor
	} else {
		r = secondColor
	}
	return fmt.Sprintf("rgb(%d, %d, %d)", r, g, b)
}

func hasInfinite(nums ...float64) bool {
	for _, num := range nums {
		if math.IsInf(num, 0) {
			return true
		}
	}
	return false
}
