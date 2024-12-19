package main

import "time"

type Snake struct {
	Fragments []string

	Head string

	X int

	Y int
}

func (s *Snake) Go(whichKey chan rune, isEnd chan bool, a *Apple) {

	for {

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
		case 0:
			s.X = 10
		}

		switch s.Y {
		case 10:
			s.Y = 0
		case 0:
			s.Y = 10
		}

		if s.X == a.X && s.Y == a.Y {
			s.Fragments = append(s.Fragments, "o")
			*a = Create()
		}

		time.Sleep(time.Second / 2)
	}

}
