package main

import "math/rand"

type Apple struct {
	Image string
	X     int
	Y     int
}

func Create() Apple {

	return Apple{
		X: rand.Intn(10),
		Y: rand.Intn(10),
	}

}
