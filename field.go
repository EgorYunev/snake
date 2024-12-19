package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"
)

type Field struct {
	SizeX   int
	SizeY   int
	Diametr [10][10]string
}

func (f *Field) Draw(s *Snake, a *Apple) {
	for {
		fmt.Println("Enter esc to exit game")
		for i := 0; i < len(f.Diametr); i++ {
		Y:
			for j := 0; j < len(f.Diametr[i]); j++ {
				for _, fragment := range s.Fragments {
					if i == fragment.X && j == fragment.Y {
						fmt.Print(fragment.image)
						continue Y
					}
				}
				if i == a.X && j == a.Y {
					fmt.Print("A")
				} else {
					fmt.Print(".")
				}

			}
			fmt.Println()
		}
		time.Sleep(time.Second / 4)
		CallClear()
	}

}

var clear map[string]func() //create a map for storing clear funcs

func init() {
	clear = make(map[string]func()) //Initialize it
	clear["linux"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func CallClear() {
	value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	if ok {                          //if we defined a clear func for that platform:
		value() //we execute it
	} else { //unsupported platform
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}
