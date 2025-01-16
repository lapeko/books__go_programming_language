## Exercise 3.1
If the function f returns a non-finite float64 value, the SVG file will contain
invalid <polygon> elements (although many SVG renderers handle this gracefully). Modify
the program to skip invalid polygons.

## Exercise 3.2
Experiment with visualizations of other functions from the math package. Can
you produce an egg box, moguls, or a saddle?

## Exercise 3.3
Color each polygon based on its height, so that the peaks are colored red
(#ff0000) and the valleys blue (#0000ff).

## Exercise 3.4
Following the approach of the Lissajous example in Section 1.7, construct a web
server that computes surfaces and writes SVG data to the client. The server must set the Con-
tent-Type header like this:
```
w.Header().Set("Content-Type", "image/svg+xml")
```