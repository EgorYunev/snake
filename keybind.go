package main

import (
	"time"

	"github.com/eiannone/keyboard"
)

func GetKeys(whichKey chan rune) {

	if err := keyboard.Open(); err != nil {
		panic(err)
	}

	defer keyboard.Close()

	for {
		char, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}
		if key == keyboard.KeyEsc {
			break
		}
		whichKey <- char
		time.Sleep(time.Millisecond * 100)
	}

}
