package main

import (
	"fmt"
	"time"
)

type Field struct {
	SizeX   int
	SizeY   int
	Diametr [10][10]string
}

func (f *Field) Draw(s *Snake, a *Apple) {
	for {
		for i := 0; i < len(f.Diametr); i++ {
			for j := 0; j < len(f.Diametr[i]); j++ {
				if i == s.X && j == s.Y {
					fmt.Print("0")
				} else if i == a.X && j == a.Y {
					fmt.Print("A")
				} else {
					fmt.Print(".")
				}
			}
			fmt.Println()
		}
		time.Sleep(time.Second / 2)
	}
}
