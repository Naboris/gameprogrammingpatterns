package direction

type command interface {
	Execute(*GameActor)
}

type jumpCommand struct{}

func (jumpCommand) Execute(actor *GameActor) {
	actor.jump()
}

type fireCommand struct{}

func (fireCommand) Execute(actor *GameActor) {
	actor.fire()
}

type lurchCommand struct{}

func (lurchCommand) Execute(actor *GameActor) {
	actor.lurch()
}

type swapCommand struct{}

func (swapCommand) Execute(actor *GameActor) {
	actor.swap()
}
