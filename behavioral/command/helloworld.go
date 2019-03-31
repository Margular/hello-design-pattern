package command

import (
	"fmt"
	"math/rand"
)

// STEP 1: Define a receiver
type Person struct {
	name 	string
	age		uint
}

func (p Person) Say() {
	fmt.Printf("I'm %s, I am %d years old.\n", p.name, p.age)
}

// STEP 2: Define a command interface that have a method called execute()
type Command interface {
	Execute()
}

// STEP 3: Implementation 1
type RandName struct {
	P *Person
}

func (r RandName) Execute() {
	r.P.name = map[int]string {
		0 : "Tom",
		1 : "Jerry",
		2 : "Faker",
	}[rand.Intn(3)]
}

// STEP 4: Implementation 2
type IncreaseAge struct {
	P *Person
}

func (i IncreaseAge) Execute() {
	i.P.age++
}

// STEP 4: Define an invoker
type Grow struct {
	Cmds []Command
}

func (g Grow) Call() {
	for _, cmd := range g.Cmds {
		cmd.Execute()
	}
}