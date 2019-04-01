package memento

// STEP 1: Define a memento object that contains some data
type Memento struct {
	State string
}

// STEP 2: Define an originator that has some data, too
type Originator struct {
	State string
}

func (o Originator) SaveStateToMemento() Memento {
	return Memento{o.State}
}

func (o *Originator) GetStateFromMemento(m Memento) {
	o.State = m.State
}

// STEP 3: Define a caretaker that stores a list of Memento objects
type CareTaker struct {
	mementos []Memento
}

func (ct *CareTaker) Add(m Memento) {
	if ct.mementos == nil {
		ct.mementos = []Memento{m}
		return
	}

	ct.mementos = append(ct.mementos, m)
}

func (ct CareTaker) Get(index uint) Memento{
	return ct.mementos[index]
}