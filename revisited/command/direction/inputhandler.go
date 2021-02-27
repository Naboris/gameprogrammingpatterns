package direction

import "math/rand"

type input int

const (
	buttonX input = iota
	buttonY
	buttonA
	buttonB
)

type inputHandler struct {
	buttonX, buttonY, buttonA, buttonB command
}

func NewInputHandler() inputHandler {
	return inputHandler{
		buttonX: jumpCommand{},
		buttonY: fireCommand{},
		buttonA: lurchCommand{},
		buttonB: swapCommand{},
	}
}

func (ih inputHandler) HandleInput() command {
	switch {
	case isPressed(buttonX):
		return ih.buttonX
	case isPressed(buttonY):
		return ih.buttonY
	case isPressed(buttonA):
		return ih.buttonA
	case isPressed(buttonB):
		return ih.buttonB
	default:
		return nil
	}
}

func isPressed(button input) bool {
	return rand.Intn(2) > 0
}
