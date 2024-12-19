package main

import "time"

type Snake struct {
	Fragments []Fragment

	Head Fragment

	Direction rune
}

type Fragment struct {
	X, Y  int
	image string
}

func (s *Snake) Go(whichKey chan rune, isEnd chan bool, a *Apple) {

	for {

		for i := len(s.Fragments) - 1; i > 0; i-- {
			s.Fragments[i].X = s.Fragments[i-1].X
			s.Fragments[i].Y = s.Fragments[i-1].Y
		}

		switch s.Direction {
		case 's':
			s.Head.X++
		case 'w':
			s.Head.X--
		case 'a':
			s.Head.Y--
		case 'd':
			s.Head.Y++
		}

		switch s.Head.X {
		case 10:
			s.Head.X = 0
		case -1:
			s.Head.X = 9
		}

		switch s.Head.Y {
		case 10:
			s.Head.Y = 0
		case -1:
			s.Head.Y = 9
		}
		for i, fragment := range s.Fragments {
			if s.Head.X == fragment.X && s.Head.Y == fragment.Y && i != 0 {
				isEnd <- true
			}
		}
		s.Fragments[0] = s.Head

		if s.Head.X == a.X && s.Head.Y == a.Y {
			s.Fragments = append(s.Fragments, Fragment{image: "o"})
			*a = CreateApple(s)
		}
		select {
		case key := <-whichKey:
			s.Direction = key
		default:
		}

		time.Sleep(time.Second / 4)
	}

}
