package configuring

type Command interface {
	Execute()
}

type JumpCommand struct{}

func (JumpCommand) Execute() {
	jump()
}

type FireCommand struct{}

func (FireCommand) Execute() {
	fire()
}

type LurchCommand struct{}

func (LurchCommand) Execute() {
	lurch()
}

type SwapCommand struct{}

func (SwapCommand) Execute() {
	swap()
}
