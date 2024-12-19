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

	snake := Snake{
		X: 5,
		Y: 5,
	}

	whichKey := make(chan rune)
	isEnd := make(chan bool)

	apple := Create()
	go snake.Go(whichKey, isEnd, &apple)
	go field.Draw(&snake, &apple)
	go GetKeys(whichKey)

	if <-isEnd {
		fmt.Println("Game over")
	}

}
