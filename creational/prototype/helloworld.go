package prototype

import "fmt"

// STEP 1: Define the object you want to use
type Options struct {
	hello string
	world string
}

func (opt Options) Print() {
	fmt.Printf("%s %s\n", opt.hello, opt.world)
}

// STEP 2: Define a method that return an new object
func NewOptions(hello, world string) Options {
	return Options{
		hello,
		world,
	}
}