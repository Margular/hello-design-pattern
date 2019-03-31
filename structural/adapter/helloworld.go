package adapter

import "fmt"

// STEP 1: Define an interface
type Speaker interface {
	Say()
}

// STEP 2: Define an object that do the same thing but not implements the interface
type HelloWorld struct {}

func (hw HelloWorld) SayHello() {
	fmt.Println("Hello World")
}

// STEP 3: Define an adapter that implements the interface and combines the object into itself
type Adapter struct {
	HW HelloWorld
}

func (a Adapter) Say() {
	a.HW.SayHello()
}