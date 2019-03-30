package singleton

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// STEP 1: Define the object you want to use
type helloWorld struct {
	id uint64
}

func (hw helloWorld) Say() {
	fmt.Printf("hello world: %d\n", hw.id)
}

// STEP 2: Define a pointer point to the object, and a locker "once" for thread-safety
var (
	hw *helloWorld
	once sync.Once
)

// STEP 3: Define a function return a pointer to always to the same object
func HelloWorld() *helloWorld {
	once.Do(func() {
		rand.Seed(time.Now().UnixNano())
		hw = &helloWorld{
			rand.Uint64(),
		}
	})

	return hw
}