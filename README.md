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