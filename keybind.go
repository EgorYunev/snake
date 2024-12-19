package main

import (
	"github.com/eiannone/keyboard"
)

func GetKeys(whichKey chan rune, isEnd chan bool) {

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
			isEnd <- true
			break
		}
		whichKey <- char
	}

}
