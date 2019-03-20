# hello-design-pattern
Hello world in GoLang using all 23 kinds of GoF design patterns. Thanks for https://github.com/code4craft/hello-design-pattern

## Creational

### AbstractFactory

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
	println("Hello World!")
}

type SimpleNegativeSpeaker struct {}

func (s SimpleNegativeSpeaker) BadSay() {
	println("Goodbye World!")
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
	println("Hello Hello Hello World Very Very Very Much!!!")
}

type ExtremeNegativeSpeaker struct {}

func (e ExtremeNegativeSpeaker) BadSay() {
	println("Goodbye Goodbye Goodbye World Very Very Very Much!!!")
}

type ExtremeSpeakerFactory struct {}

func (s ExtremeSpeakerFactory) CreatePositiveSpeaker() abstractFactory.PositiveSpeaker {
	return ExtremePositiveSpeaker{}
}

func (s ExtremeSpeakerFactory) CreateNegativeSpeaker() abstractFactory.NegativeSpeaker {
	return ExtremeNegativeSpeaker{}
}
```