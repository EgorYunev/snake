package main

import "fmt"

func main() {
	start()
}

func start() {
	field := Field{
		SizeX: 10,
		SizeY: 10,
	}

	field.Diametr = [10][10]string{}

	head := Fragment{X: 5, Y: 5, image: "0"}

	snake := Snake{
		Fragments: []Fragment{head, Fragment{X: 5, Y: 4, image: "o"}, Fragment{X: 5, Y: 3, image: "o"}},
		Head:      head,
		Direction: 'd',
	}

	whichKey := make(chan rune)
	isEnd := make(chan bool)

	apple := CreateApple(&snake)
	go snake.Go(whichKey, isEnd, &apple)
	go field.Draw(&snake, &apple)
	go GetKeys(whichKey, isEnd)

	if <-isEnd {
		fmt.Println("Game over")
	}

}
