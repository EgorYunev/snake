package main

import "time"

type Snake struct {
	Fragments []string

	Head string

	X int

	Y int
}

func (s *Snake) Create(f *Field) {
	f.Diametr[s.X][s.Y] = "0"
}

func (s *Snake) Eat(a *Apple) {
	if s.X == a.X && s.Y == a.Y {
		s.Fragments = append(s.Fragments, ">")
	}

}

func (s *Snake) Go() {

	for {
		s.X++
		time.Sleep(time.Millisecond * 100)
	}

}
