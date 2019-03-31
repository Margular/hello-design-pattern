package bridge

import "fmt"

// STEP 1: Define an object you want to use and an interface member
type Person struct {
	M Mouth
}

// STEP 2: Define a function that using the ability of the interface member
func (p Person) Speak() {
	p.M.Say()
}

// STEP 3: Define the interface
type Mouth interface {
	Say()
}

// STEP 4: Implements the interface
type HelloMouth struct {}

func (m HelloMouth) Say() {
	fmt.Println("Hello World")
}