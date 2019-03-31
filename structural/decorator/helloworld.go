package decorator

import "fmt"

// STEP 1: Define an interface
type Person interface {
	Say()
}

// STEP 2: Define some implementations
type Tom struct {}

func (t Tom) Say() {
	fmt.Println("Hello World! I'm Tom")
}

// STEP 2: Define a decorator that aggregation the interface
type PersonDecorator interface {
	Person
}

// STEP 3: Implements the decorator with 2 things:
// 1. invoke method of the interface
// 2. do something else(this calls decorating the interface)
type AnoisingPerson struct {
	P Person
}

func (ap AnoisingPerson) Say() {
	ap.P.Say()
	fmt.Println("I'm anoising!")
}