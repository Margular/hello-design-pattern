package state

import "fmt"

// STEP 1: Define some states
type State interface {
	handle()
}

type StateGood struct {}

func (sg StateGood) handle() {
	fmt.Println("Hello World! I'm good")
}

type StateBad struct {}

func (sb StateBad) handle() {
	fmt.Println("Hello World! I'm bad")
}

// STEP 2:Define the context object that manages state transform
type Context struct {
	state State
}

func (ctx *Context) ChangeState(s *State) {
	ctx.state = *s
}

func (ctx *Context) Request() {
	switch ctx.state.(type) {
	case StateBad:
		ctx.state = StateGood{}
	case StateGood:
		ctx.state = StateBad{}
	default:
		ctx.state = StateGood{}
	}

	ctx.state.handle()
}