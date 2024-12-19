package main

import "math/rand"

type Apple struct {
	Image string
	X     int
	Y     int
}

func CreateApple(s *Snake) Apple {
	a := Apple{X: rand.Intn(10), Y: rand.Intn(10)}

	for _, f := range s.Fragments {
		if f.X == a.X && f.Y == a.Y {
			a = CreateApple(s)
		}
	}
	return a
}
