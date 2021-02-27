package configuring

import "math/rand"

type input int

const (
	buttonX input = iota
	buttonY
	buttonA
	buttonB
)

type InputHandler struct {
	ButtonX, ButtonY, ButtonA, ButtonB Command
}

func NewInputHandler() InputHandler {
	return InputHandler{
		ButtonX: JumpCommand{},
		ButtonY: FireCommand{},
		ButtonA: LurchCommand{},
		ButtonB: SwapCommand{},
	}
}

func (ih InputHandler) HandleInput() {
	if isPressed(buttonX) {
		ih.ButtonX.Execute()
	} else if isPressed(buttonY) {
		ih.ButtonY.Execute()
	} else if isPressed(buttonA) {
		ih.ButtonA.Execute()
	} else if isPressed(buttonB) {
		ih.ButtonB.Execute()
	}
}

func isPressed(button input) bool {
	return rand.Intn(2) > 0
}
