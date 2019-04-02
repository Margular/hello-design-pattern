# hello-design-pattern
Hello world in GoLang using all 23 kinds of GoF design patterns. Thanks for https://github.com/code4craft/hello-design-pattern

## Creational

### Abstract Factory

```go
func tryAbstractFactory(speakerFactory abstractFactory.SpeakerFactory) {
	positiveSpeaker := speakerFactory.CreatePositiveSpeaker()
	negativeSpeaker := speakerFactory.CreateNegativeSpeaker()
	positiveSpeaker.GoodSay()
	negativeSpeaker.BadSay()
}

func main() {
	tryAbstractFactory(factory.SimpleSpeakerFactory{})
	tryAbstractFactory(factory.ExtremeSpeakerFactory{})
}
```

```text
output:
Hello World!
Goodbye World!
Hello Hello Hello World Very Very Very Much!!!
Goodbye Goodbye Goodbye World Very Very Very Much!!!
```

```go
// STEP 1: Define some interfaces you want to use
type PositiveSpeaker interface {
	GoodSay()
}

type NegativeSpeaker interface {
	BadSay()
}

// STEP 2: Define the abstract factory, it commonly has some methods to create some instances of interface
// that you have defined in STEP 1
type SpeakerFactory interface {
	CreatePositiveSpeaker() PositiveSpeaker
	CreateNegativeSpeaker() NegativeSpeaker
}

// STEP 3: Define an implementation
type SimplePositiveSpeaker struct {}

func (s SimplePositiveSpeaker) GoodSay() {
	fmt.Println("Hello World!")
}

type SimpleNegativeSpeaker struct {}

func (s SimpleNegativeSpeaker) BadSay() {
	fmt.Println("Goodbye World!")
}

type SimpleSpeakerFactory struct {}

func (s SimpleSpeakerFactory) CreatePositiveSpeaker() abstractFactory.PositiveSpeaker {
	return SimplePositiveSpeaker{}
}

func (s SimpleSpeakerFactory) CreateNegativeSpeaker() abstractFactory.NegativeSpeaker {
	return SimpleNegativeSpeaker{}
}

// STEP 4,5,...: Define any other implementations
type ExtremePositiveSpeaker struct {}

func (s ExtremePositiveSpeaker) GoodSay() {
	fmt.Println("Hello Hello Hello World Very Very Very Much!!!")
}

type ExtremeNegativeSpeaker struct {}

func (e ExtremeNegativeSpeaker) BadSay() {
	fmt.Println("Goodbye Goodbye Goodbye World Very Very Very Much!!!")
}

type ExtremeSpeakerFactory struct {}

func (s ExtremeSpeakerFactory) CreatePositiveSpeaker() abstractFactory.PositiveSpeaker {
	return ExtremePositiveSpeaker{}
}

func (s ExtremeSpeakerFactory) CreateNegativeSpeaker() abstractFactory.NegativeSpeaker {
	return ExtremeNegativeSpeaker{}
}
```

### Builder
```go
func main() {
	hwBuilder := builder.HelloWorldBuilder{}
	hw := hwBuilder.SetHello("hello").SetWorld("world").Build()
	hw.Say()
}
```

```text
output:
hello world
```

```go
// STEP 1: Define the struct you want to use
type HelloWorld struct {
	hello string
	world string
}

func (hw HelloWorld) Say() {
	fmt.Printf("%s %s\n", hw.hello, hw.world)
}

// STEP 2: Define the builder you want to use to build the struct we've defined before
type HelloWorldBuilder struct {
	hello string
	world string
}

func (b *HelloWorldBuilder) SetHello(hello string) *HelloWorldBuilder {
	b.hello = hello
	return b
}

func (b *HelloWorldBuilder) SetWorld(world string) *HelloWorldBuilder {
	b.world = world
	return b
}

func (b *HelloWorldBuilder) Build() HelloWorld {
	return HelloWorld{
		b.hello,
		b.world,
	}
}
```

### Factory Method
```go
func main() {
	var speaker factoryMethod.Speaker
	hwf := factoryMethod.HelloWorldFactory{}
	speaker = hwf.CreateSpeaker()
	fmt.Println(speaker.Words)
}
```

```text
output:
Hello World
```

```go
// STEP 1: Define the object you want to use
type Speaker struct {
	Words string
}

// STEP 2: Define an interface which contains a method that can be used to create the object
type speakerFactory interface {
	CreateSpeaker() Speaker
}

// STEP 3: Define a factory which implements the interface we've defined before
type HelloWorldFactory struct{}

func (hwf HelloWorldFactory) CreateSpeaker() Speaker {
	return Speaker{"Hello World"}
}
```

### Prototype
```go
func main() {
	options := prototype.NewOptions("hello", "world")
	options.Print()
}
```

```text
output:
hello world
```

```go
// STEP 1: Define the object you want to use
type Options struct {
	hello string
	world string
}

func (opt Options) Print() {
	fmt.Printf("%s %s\n", opt.hello, opt.world)
}

// STEP 2: Define a method that return an new object
func NewOptions(hello, world string) Options {
	return Options{
		hello,
		world,
	}
}
```

### Singleton
```go
func main() {
	singleton.HelloWorld().Say()
	singleton.HelloWorld().Say()
}
```

```text
output:
hello world: 3611123106840782002
hello world: 3611123106840782002
```

```go
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
```

## Structural

### Adapter
```go
func main() {
	var target adapter.Speaker
	target = adapter.Adapter{adapter.HelloWorld{}}
	target.Say()
}
```

```text
output:
Hello World
```

```go
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
```

### Bridge
```go
func main() {
	var tom = bridge.Person{}
	tom.M = bridge.HelloMouth{}
	tom.Speak()
}
```

```text
output:
Hello World
```

```go
// STEP 1: Define an object you want to use and an interface member
type Person struct {
	M Mouth
}

// STEP 2: Define a function that using the ability of the interface member
func (p Person) Speak() {
	p.M.Say()
}

// STEP 3: Define the interface
type Mouth interface {
	Say()
}

// STEP 4: Implements the interface
type HelloMouth struct {}

func (m HelloMouth) Say() {
	fmt.Println("Hello World")
}
```

### Composite
```go
func main() {
	t := composite.Tom{}
	j := composite.Jerry{}
	world := composite.World{
		[]composite.Speaker{
			t,
			j,
		},
	}
	world.Say()
}
```

```text
output:
Hello World! I'm Tom
Hello World! I'm Jerry
```

```go
// STEP 1: Define an interface that have several methods
type Speaker interface {
	Say()
}

// STEP 2: Define some objects the implements the interface
type Tom struct {}

func (t Tom) Say() {
	fmt.Println("Hello World! I'm Tom")
}

type Jerry struct {}

func (j Jerry) Say() {
	fmt.Println("Hello World! I'm Jerry")
}

// STEP 3: Define an object that implements the interface and have a list memeber of the interface
type World struct{
	Speakers []Speaker
}

// STEP 4: Define those methods, commonly invoke all implements of the list member
func (w World) Say() {
	for _, speaker := range w.Speakers {
		speaker.Say()
	}
}
```

### Decorator
```go
func main() {
	var dec decorator.PersonDecorator
	dec = decorator.AnoisingPerson{decorator.Tom{}}
	dec.Say()
}
```

```text
output:
Hello World! I'm Tom
I'm anoising!
```

```go
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
```

### Facade
```go
func main() {
	var house facade.House
	house.AllSay()
}
```

```text
output:
Hello World! I'm Tom
Hello World! I'm Jerry
```

```go
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
```

### Flyweight
```go
func main() {
	flyweight.AnimalFactory().GetAnimal("Tom").Say()
	flyweight.AnimalFactory().GetAnimal("Tom").Say()
	flyweight.AnimalFactory().GetAnimal("Jerry").Say()
	flyweight.AnimalFactory().GetAnimal("Tom").Say()
	flyweight.AnimalFactory().GetAnimal("Jerry").Say()
}
```

```text
output:
Hello World! I'm animal 1815427177354915630
Hello World! I'm animal 1815427177354915630
Hello World! I'm animal 13895412513112872184
Hello World! I'm animal 1815427177354915630
Hello World! I'm animal 13895412513112872184
```

```go
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
```

### Proxy
```go
func main() {
	var p proxy.Subject
	p := proxy.Proxy{}
	p.Request()
}
```

```text
output:
Hello World! I'm Tom
```

```go
// STEP 1: Define the subject interface
type Subject interface {
	Request()
}

// STEP 2: Define a real subject
type TomSubject struct {}

func (t TomSubject) Request() {
	fmt.Println("Hello World! I'm Tom")
}

// STEP 3: Define a proxy that implements the subject and have a real subject member
type Proxy struct {
	tom *TomSubject
}

func (p *Proxy) preRequest() {
	if p.tom == nil {
		p.tom = &TomSubject{}
	}
}

func (p Proxy) Request() {
	p.preRequest()
	p.tom.Request()
}
```

## Behavioral

### Chain of responsibility
```go
func main() {
	peter := chainOfResponsibility.Person{Name : "Peter"}
	fox := chainOfResponsibility.Person{Name : "Fox", Successor : &peter}
	steven := chainOfResponsibility.Person{Name : "Steven", Successor : &fox}
	var theWorld chainOfResponsibility.PersonHandler = chainOfResponsibility.Person{Successor : &steven}

	theWorld.HandleRequest("Peter")
	theWorld.HandleRequest("Fox")
	theWorld.HandleRequest("Tom")
}
```

```text
output:
Hello World! I'm Peter
Hello World! I'm Fox
Tom do not found in this world!
```

```go
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
```

### Command
```go
func main() {
	person := &command.Person{}
	randName := command.RandName{person}
	increaseAge := command.IncreaseAge{person}
	grow := command.Grow{[]command.Command{randName, increaseAge}}
	grow.Call(); person.Say()
	grow.Call(); person.Say()
	grow.Call(); person.Say()
}
```

```text
output:
I'm Jerry, I am 1 years old.
I'm Jerry, I am 2 years old.
I'm Faker, I am 3 years old.
```

```go
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
```

### Interpreter
```go
func main() {
	helloInter := interpreter.HelloInterpreter{}
	helloInter.RegFunc("println", interpreter.MyPrintln{})
	helloInter.Interpret("println('hello world')")
}
```

```text
output:
hello world
```

```go
// STEP 1: Define an interface
type Func interface {
	call(string)
}

// STEP 2: Implements the interface
type MyPrintln struct {}

func (pl MyPrintln) call(param string) {
	fmt.Println(param)
}

// STEP 3: Define an interpreter object that have ability to invoke registered functions
type HelloInterpreter struct {
	funcs map[string]Func
}

func (hw *HelloInterpreter) RegFunc(name string, f Func) {
	if hw.funcs == nil {
		hw.funcs = map[string]Func{}
	}

	hw.funcs[name] = f
}

func (hw *HelloInterpreter) Interpret(expr string) {
	funcName := strings.TrimSpace(expr[:strings.Index(expr, "(")])
	f := hw.funcs[funcName]

	after := expr[strings.Index(expr, "(") + 1:]
	param := strings.Trim(after[:strings.LastIndex(after, ")")], "'")

	f.call(param)
}
```

### Mediator
```go
func main() {
	m := mediator.Mediator{}

	c1 := mediator.Collegue{1, &m}
	c2 := mediator.Collegue{2, &m}
	c3 := mediator.Collegue{3, &m}


	m.Register(&c1)
	m.Register(&c2)
	m.Register(&c3)

	c1.SendMsg(2, "Hello Tom")
	c3.SendMsg(1, "Hello Jerry")
	c2.SendMsg(1, "Hello Jerry")
}
```

```text
output:
Hello World! I'm collegue 2, I got msg: Hello Tom
Hello World! I'm collegue 1, I got msg: Hello Jerry
Hello World! I'm collegue 1, I got msg: Hello Jerry
```

```go
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
```

### Memento
```go
func main() {
	originator := memento.Originator{}
	careTaker := memento.CareTaker{}
	originator.State = "Hello World"
	careTaker.Add(originator.SaveStateToMemento())
	originator.State = "Goodbye World"
	careTaker.Add(originator.SaveStateToMemento())
	fmt.Println("Current State: " + originator.State)
	originator.GetStateFromMemento(careTaker.Get(0))
	fmt.Println("Current State: " + originator.State)
}
```

```text
output:
Current State: Goodbye World
Current State: Hello World
```

```go
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
```

### State
```go
func main() {
	ctx := state.Context{}
	ctx.Request()
	ctx.Request()
	ctx.Request()
}
```

```text
output:
Hello World! I'm good
Hello World! I'm bad
Hello World! I'm good
```

```go
// STEP 1: Define some states
type State interface {
	handle()
}

type StateGood struct {}

func (sg StateGood) handle() {
	fmt.Println("Hello World! I'm good")
}

type StateBad struct {}

func (sb StateBad) handle() {
	fmt.Println("Hello World! I'm bad")
}

// STEP 2:Define the context object that manages state transform
type Context struct {
	state State
}

func (ctx *Context) ChangeState(s *State) {
	ctx.state = *s
}

func (ctx *Context) Request() {
	switch ctx.state.(type) {
	case StateBad:
		ctx.state = StateGood{}
	case StateGood:
		ctx.state = StateBad{}
	default:
		ctx.state = StateGood{}
	}

	ctx.state.handle()
}
```

### Visitor
```go
func main() {
	h := visitor.House{}
	h.Accept(visitor.PersonVisitor{})
}
```

```text
output:
Hello World! You are Tom
Hello World! You are Jerry
```

```go
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
```