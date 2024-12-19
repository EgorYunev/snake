package main

import "time"

type Snake struct {
	Fragments []Fragment

	Head string

	X int

	Y int
}

type Fragment struct {
	X, Y  int
	image string
}

func (s *Snake) Go(whichKey chan rune, isEnd chan bool, a *Apple) {

	for {

		for i := len(s.Fragments) - 1; i > 0; i-- {
			s.Fragments[i] = s.Fragments[i-1]
		}

		s.Fragments[0] = Fragment{s.X, s.Y, "0"}

		switch <-whichKey {
		case 's':
			s.X++
		case 'w':
			s.X--
		case 'a':
			s.Y--
		case 'd':
			s.Y++
		}

		switch s.X {
		case 10:
			s.X = 0
		case -1:
			s.X = 9
		}

		switch s.Y {
		case 10:
			s.Y = 0
		case -1:
			s.Y = 9
		}

		if s.X == a.X && s.Y == a.Y {
			s.Fragments = append(s.Fragments, Fragment{image: "0"})
			*a = Create()
		}

		time.Sleep(time.Second / 4)
	}

}
