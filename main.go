package main

import (
	"fmt"
	"strconv"
	"syscall/js"

	"github.com/golang/geo/s2"
)

func main() {
	calculateS2ID()

	done := make(chan struct{})

	calculateS2IDCallback := js.NewCallback(func(args []js.Value) {
		calculateS2ID()
	})
	calculateLatLngCallback := js.NewCallback(func(args []js.Value) {
		calculateLatLng()
	})

	js.Global().Get("document").
		Call("getElementById", "latitude").
		Call("addEventListener", "change", calculateS2IDCallback)
	js.Global().Get("document").
		Call("getElementById", "latitude").
		Call("addEventListener", "keyup", calculateS2IDCallback)

	js.Global().Get("document").
		Call("getElementById", "longitude").
		Call("addEventListener", "change", calculateS2IDCallback)
	js.Global().Get("document").
		Call("getElementById", "longitude").
		Call("addEventListener", "keyup", calculateS2IDCallback)

	js.Global().Get("document").
		Call("getElementById", "level").
		Call("addEventListener", "change", calculateS2IDCallback)
	js.Global().Get("document").
		Call("getElementById", "level").
		Call("addEventListener", "keyup", calculateS2IDCallback)

	js.Global().Get("document").
		Call("getElementById", "s2id").
		Call("addEventListener", "change", calculateLatLngCallback)
	js.Global().Get("document").
		Call("getElementById", "s2id").
		Call("addEventListener", "keyup", calculateLatLngCallback)

	<-done
}

func calculateS2ID() {
	latitudeStr := js.Global().Get("document").
		Call("getElementById", "latitude").
		Get("value").String()
	longitudeStr := js.Global().Get("document").
		Call("getElementById", "longitude").
		Get("value").String()
	levelStr := js.Global().Get("document").
		Call("getElementById", "level").
		Get("value").String()

	if latitudeStr == "" || longitudeStr == "" || levelStr == "" {
		return
	}

	latitude, err := strconv.ParseFloat(latitudeStr, 64)
	if err != nil {
		fmt.Println("Failed to parse latitude")
		return
	}

	longitude, err := strconv.ParseFloat(longitudeStr, 64)
	if err != nil {
		fmt.Println("Failed to parse longitude")
		return
	}

	level, err := strconv.Atoi(levelStr)
	if err != nil {
		fmt.Println("Failed to parse level")
		return
	}

	s2ID := strconv.Itoa(int(s2.CellIDFromLatLng(s2.LatLngFromDegrees(latitude, longitude)).Parent(level)))

	js.Global().Get("document").
		Call("getElementById", "s2id").
		Set("value", s2ID)
}

func calculateLatLng() {
	s2idStr := js.Global().Get("document").
		Call("getElementById", "s2id").
		Get("value").String()

	if s2idStr == "" {
		return
	}

	s2idInt, err := strconv.Atoi(s2idStr)
	if err != nil {
		fmt.Println("Failed to parse s2id")
		return
	}

	s2Cell := s2.CellID(s2idInt)
	ll := s2Cell.LatLng()
	latitude := fmt.Sprintf("%v", ll.Lat)
	longitude := fmt.Sprintf("%v", ll.Lng)
	level := s2Cell.Level()

	js.Global().Get("document").
		Call("getElementById", "latitude").
		Set("value", latitude)
	js.Global().Get("document").
		Call("getElementById", "longitude").
		Set("value", longitude)
	js.Global().Get("document").
		Call("getElementById", "level").
		Set("value", level)
}
