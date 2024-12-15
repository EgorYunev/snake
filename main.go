package main

//Реализовать удлинение змейки
//Поедание яблока
//Движение

func main() {
	start()
}

func start() {
	field := Field{
		SizeX: 10,
		SizeY: 10,
	}

	field.Diametr = [10][10]string{}

	field.Create()

	snake := Snake{
		X: 5,
		Y: 5,
	}

	apple := Apple{
		Image: "A",
	}

	snake.Create(&field)
	apple.Create(&field)
	// TODO - гоурутины не работают, доделать
	go snake.Go()
	go field.Draw()

}
