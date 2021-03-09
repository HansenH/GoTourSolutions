package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
	"math"
)

type Image struct {
	bounds     image.Rectangle
	colorModel color.Model
	pixColor   func(x, y int) uint8
}

func (m Image) Bounds() image.Rectangle {
	return m.bounds
}

func (m Image) ColorModel() color.Model {
	return m.colorModel
}

func (m Image) At(x, y int) color.Color {
	v := m.pixColor(x, y)
	return color.RGBA{v, v, 255, 255}
}

func main() {
	m := Image{
		bounds:     image.Rect(0, 0, 256, 256),
		colorModel: color.RGBAModel,
		pixColor: func(x, y int) uint8 {
			return uint8(float64(y) * math.Log(float64(x)))
		},
	}
	pic.ShowImage(m)
}
