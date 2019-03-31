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