package direction

import "fmt"

type GameActor struct {
	N int
}

func (ga *GameActor) jump() {
	fmt.Printf("Actor %d: JUMP\n", ga.N)
}
func (ga *GameActor) fire() {
	fmt.Printf("Actor %d: FIRE\n", ga.N)
}
func (ga *GameActor) lurch() {
	fmt.Printf("Actor %d: LURCH\n", ga.N)
}
func (ga *GameActor) swap() {
	fmt.Printf("Actor %d: SWAP\n", ga.N)
}
