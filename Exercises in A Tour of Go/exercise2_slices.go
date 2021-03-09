package main

import (
	"golang.org/x/tour/pic"
	"math"
)

func Pic(dx, dy int) [][]uint8 {
	a := make([][]uint8, dy)
	for x := range a {
		a[x] = make([]uint8, dx)
		for y := range a[x] {
			a[x][y] = uint8(float64(x) * math.Log(float64(y)))
		}
	}
	return a
}

func main() {
	pic.Show(Pic)
}
