package configuring

import "math/rand"

type input int

const (
	BUTTON_X input = iota
	BUTTON_Y
	BUTTON_A
	BUTTON_B
)

type InputHandler struct {
	buttonX, buttonY, buttonA, buttonB Command
}

func (ih InputHandler) handleInput() {
	if isPressed(BUTTON_X) {
		ih.buttonX.Execute()
	} else if isPressed(BUTTON_Y) {
		ih.buttonY.Execute()
	} else if isPressed(BUTTON_A) {
		ih.buttonA.Execute()
	} else if isPressed(BUTTON_B) {
		ih.buttonB.Execute()
	}
}

func isPressed(button input) bool {
	return rand.Intn(2) > 0
}
