# hello-design-pattern
Hello world in GoLang using all 23 kinds of GoF design patterns. Thanks for https://github.com/code4craft/hello-design-pattern

## AbstractFactory
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

We are creational patterns!

1. Abstract Factory: 
Hello World!
Goodbye World!
Hello Hello Hello World Very Very Very Much!!!
Goodbye Goodbye Goodbye World Very Very Very Much!!!
```