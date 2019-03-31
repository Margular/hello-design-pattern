package flyweight

import (
	"fmt"
	"math/rand"
)

// STEP 1: Define an object
type Animal struct {
	id 		uint64	// intrinsic
}

func (a Animal) Say() {
	fmt.Printf("Hello World! I'm animal %d\n", a.id)
}

// STEP 2: Define a factory of the object that have a map of objects
type animalFactory struct {
	animals map[string]Animal
}

// STEP 3: Define getObject method of this factory
func (af animalFactory) GetAnimal(name string) Animal {
	animal, ok := af.animals[name]

	if ok {
		return animal
	}

	animal = Animal{rand.Uint64()}
	af.animals[name] = animal
	return animal
}

// STEP 4(Optional): Singleton design pattern combined
var af *animalFactory

func AnimalFactory() *animalFactory{
	if af == nil {
		af = &animalFactory{
			map[string]Animal{},
		}
	}

	return af
}