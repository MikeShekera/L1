package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func main() {
	a := newPoint(1, 2)
	b := newPoint(4, 7)
	fmt.Println(findDistance(b, a))
}

func newPoint(xCoord, yCoord float64) Point {
	return Point{
		x: xCoord,
		y: yCoord,
	}
}

func findDistance(a, b Point) float64 {
	return math.Sqrt(math.Pow(a.x-b.x, 2) + math.Pow(a.y-b.y, 2))
}
