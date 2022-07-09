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

	err := termbox.Init()
	if err != nil {
		panic(err)
	}

	game := tetris.NewGame()
	game.Start()

	termbox.Close()
	fmt.Println("Thank you for your time. Good bye!")
}
