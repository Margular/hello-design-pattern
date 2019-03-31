package chainOfResponsibility

import "fmt"

// STEP 1: Define an interface that has a method HandleRequest(...)
type PersonHandler interface {
	HandleRequest(string)
}

// STEP 2: Define an object that implements the interface and has a member Successor and something else
type Person struct {
	Successor *Person
	Name string
}

func (p Person) HandleRequest(name string) {
	if p.Name == name {
		fmt.Println("Hello World! I'm " + name)
		return
	}

	if p.Successor == nil {
		fmt.Printf("%s do not found in this world!", name)
		return
	}

	p.Successor.HandleRequest(name)
}