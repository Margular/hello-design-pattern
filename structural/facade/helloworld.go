package facade

import "fmt"

// STEP 1: Define an object
type Tom struct {}

func (t Tom) SayTom() {
	fmt.Println("Hello World! I'm Tom")
}

// STEP 2: Define another object
type Jerry struct {}

func (j Jerry) SayJerry() {
	fmt.Println("Hello World! I'm Jerry")
}

// STEP 3: Define a facade object to get all objects in one, and privode a method to invoke them internally
type House struct {
	t Tom
	j Jerry
}

func (h House) AllSay() {
	h.t.SayTom()
	h.j.SayJerry()
}