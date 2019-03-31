package composite

import "fmt"

// STEP 1: Define an interface that have several methods
type Speaker interface {
	Say()
}

// STEP 2: Define some objects the implements the interface
type Tom struct {}

func (t Tom) Say() {
	fmt.Println("Hello World! I'm Tom")
}

type Jerry struct {}

func (j Jerry) Say() {
	fmt.Println("Hello World! I'm Jerry")
}

// STEP 3: Define an object that implements the interface and have a list memeber of the interface
type World struct{
	Speakers []Speaker
}

// STEP 4: Define those methods, commonly invoke all implements of the list member
func (w World) Say() {
	for _, speaker := range w.Speakers {
		speaker.Say()
	}
}