package builder

import "fmt"

// STEP 1: Define the struct you want to use
type HelloWorld struct {
	hello string
	world string
}

func (hw HelloWorld) Say() {
	fmt.Printf("%s %s\n", hw.hello, hw.world)
}