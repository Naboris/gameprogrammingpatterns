package main

import (
	"github.com/naboris/gameprogrammingpatterns/revisited/command/direction"
)

func main() {
	const N = 10

	ih := direction.NewInputHandler()
	a0 := &direction.GameActor{}

	for i := 0; i < N; i++ {
		aN := &direction.GameActor{N: i + 1}
		c := ih.HandleInput()
		if c != nil {
			c.Execute(a0)
			c.Execute(aN)
		}
	}
}
