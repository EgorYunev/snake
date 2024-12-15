package main

import "math/rand"

type Apple struct {
	Image string
	X     int
	Y     int
}

func (a *Apple) Create(f *Field) {

	a.X = rand.Intn(10) + 1
	a.Y = rand.Intn(10) + 1

	f.Diametr[a.X][a.Y] = a.Image

}
