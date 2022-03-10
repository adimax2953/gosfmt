package main

import (
	"time"

	"github.com/adimax2953/gosfmt"
)

func main() {
	rng := gosfmt.NewSFMT()
	rng.Seed(time.Now().UnixNano())
}
