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

// var palette = []color.Color{color.White, color.Black}
var palette = []color.Color{color.White, color.RGBA{0xAA, 0xDD, 0xBB, 0xff}, color.RGBA{0xAA, 0xDD, 0xB2, 0x33}}

const (
	whiteIndex = 0
	blackIndex = 1
	otherIndex = 2
)

func main() {
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycle   = 5
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)
	freq := rand.Float64() * 3.0
	anmi := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycle*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), otherIndex)
		}
		phase += 0.1
		anmi.Delay = append(anmi.Delay, delay)
		anmi.Image = append(anmi.Image, img)
	}
	gif.EncodeAll(out, &anmi)
}
