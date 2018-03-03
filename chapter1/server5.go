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
	"sync"
)

var mu sync.Mutex
var count int
var palette = []color.Color{color.White, color.RGBA{0xAA, 0xDD, 0xBB, 0xff}, color.RGBA{0xAA, 0xDD, 0xB2, 0x33}}

const (
	whiteIndex = 0
	blackIndex = 1
	otherIndex = 2
)

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	http.HandleFunc("/gif", gifp)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	// mu.Lock()
	// count++
	// mu.Unlock()
	// fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "From[%q] = %q\n", k, v)
	}
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}

func gifp(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	// for k, v := range r.Form {
	// 	fmt.Fprintf(w, "From[%q] = %q\n", k, v)
	// }
	// fmt.Fprintf(w, "Cycle = %q\n", r.Form.Get("cycle"))
	cycle, err := strconv.Atoi(r.Form.Get("cycle"))
	if err != nil {
		log.Print(err)
	}
	// fmt.Fprintf(w, "Cycle = %d\n", cycle)
	lissajous(w, cycle)

}

func lissajous(out io.Writer, set_cycle int) {
	const (
		// cycle   = 5
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)
	var cycle = float64(set_cycle)
	log.Printf("%f", cycle)
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
