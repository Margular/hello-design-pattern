package visitor

import "fmt"

// STEP 1: Define the object you want to be visited, it has an Accept(Visitor) method
type Person interface {
	Accept(PersonVisitor)
}

// STEP 2: Define some objects implements the interface
type Tom struct {}

func (t Tom) Accept(v PersonVisitor) {
	v.Visit(t)
}

type Jerry struct {}

func (j Jerry) Accept (v PersonVisitor) {
	v.Visit(j)
}

// STEP 3: Define an object contains all the objects you've defined above, stored as a list
type House struct {
	persons []Person
}

func (h House) Accept (v PersonVisitor) {
	if h.persons == nil {
		h.persons = []Person{Tom{}, Jerry{}}
	}

	for _, p := range h.persons {
		v.Visit(p)
	}
}

// STEP 4: Define an Visitor that can Visit any objects above except the last one
type PersonVisitor struct {}

func (v PersonVisitor) Visit(person interface{}) {
	switch person.(type) {
	case Tom:
		fmt.Println("Hello World! You are Tom")
	case Jerry:
		fmt.Println("Hello World! You are Jerry")
	default:
		fmt.Println("unknown person")
	}
}