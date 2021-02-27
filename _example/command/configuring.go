package main

import (
	"fmt"

	"github.com/naboris/gameprogrammingpatterns/revisited/command/configuring"
)

func main() {
	const N = 10
	ih := configuring.NewInputHandler()
	fmt.Print("Default:\n")
	for i := 0; i < N; i++ {
		ih.HandleInput()
	}
	fmt.Print("\nModified:\n")
	ih.ButtonX = customCommand{}
	for i := 0; i < N; i++ {
		ih.HandleInput()
	}
}

type customCommand struct{}

func (customCommand) Execute() {
	fmt.Println("CUSTOM")
}
