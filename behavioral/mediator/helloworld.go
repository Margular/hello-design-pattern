package mediator

import "fmt"

// STEP 1: Define the collegue object
type Collegue struct {
	Id int
	M *Mediator
}

func (c Collegue) receiveMsg(msg string) {
	fmt.Printf("Hello World! I'm collegue %d, I got msg: %s\n", c.Id, msg)
}

func (c Collegue) SendMsg(id int, msg string) {
	c.M.operation(id, msg)
}

// STEP 2: Define the Mediator
type Mediator struct {
	collegues map[int]Collegue
}

func (m Mediator) operation(id int, msg string) {
	m.collegues[id].receiveMsg(msg)
}

func (m *Mediator) Register(c *Collegue) {
	if m.collegues == nil {
		m.collegues = map[int]Collegue{c.Id : *c}
		return
	}

	m.collegues[c.Id] = *c
}