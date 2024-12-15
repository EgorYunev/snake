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

func (f *Field) Create() {

	for i := 0; i < len(f.Diametr); i++ {
		for j := 0; j < len(f.Diametr[i]); j++ {
			f.Diametr[i][j] = "."
		}
	}

}

func (f *Field) Draw() {
	for {
		for i := 0; i < len(f.Diametr); i++ {
			for j := 0; j < len(f.Diametr[i]); j++ {
				fmt.Print(f.Diametr[i][j])
			}
			fmt.Println()
		}
		time.Sleep(time.Millisecond * 100)
	}
}
