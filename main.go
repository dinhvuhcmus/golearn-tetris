package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/dinhvuhcmus/golearn-tetris/tetris"
	"github.com/nsf/termbox-go"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	// Init termbox library to use it.
	err := termbox.Init()
	if err != nil {
		panic(err)
	}

	game := tetris.NewGame()
	game.Start()

	//Close the library termbox.
	termbox.Close()
	//After the user quit the game. Print the goodbye sentence.
	fmt.Println("Thank you for your time. Good bye!")
}
