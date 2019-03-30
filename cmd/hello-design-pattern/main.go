package main

import (
	"fmt"
	"github.com/Margular/hello-design-pattern/creational/abstractFactory"
	"github.com/Margular/hello-design-pattern/creational/abstractFactory/factory"
	"github.com/Margular/hello-design-pattern/creational/builder"
	"github.com/Margular/hello-design-pattern/creational/factoryMethod"
	"github.com/Margular/hello-design-pattern/creational/prototype"
	"github.com/Margular/hello-design-pattern/creational/singleton"
)

func tryAbstractFactory(speakerFactory abstractFactory.SpeakerFactory) {
	positiveSpeaker := speakerFactory.CreatePositiveSpeaker()
	negativeSpeaker := speakerFactory.CreateNegativeSpeaker()
	positiveSpeaker.GoodSay()
	negativeSpeaker.BadSay()
}

func main() {
	fmt.Println("We are creational patterns!\n")

	fmt.Println("1. Abstract Factory: ")
	tryAbstractFactory(factory.SimpleSpeakerFactory{})
	tryAbstractFactory(factory.ExtremeSpeakerFactory{})
	fmt.Println()

	fmt.Println("2. Builder: ")
	hwBuilder := builder.HelloWorldBuilder{}
	hw := hwBuilder.SetHello("hello").SetWorld("world").Build()
	hw.Say()
	fmt.Println()

	fmt.Println("3. Factory method: ")
	var speaker factoryMethod.Speaker
	hwf := factoryMethod.HelloWorldFactory{}
	speaker = hwf.CreateSpeaker()
	fmt.Println(speaker.Words)
	fmt.Println()

	fmt.Println("4. Prototype: ")
	options := prototype.NewOptions("hello", "world")
	options.Print()
	fmt.Println()

	fmt.Println("5. Singleton")
	singleton.HelloWorld().Say()
	singleton.HelloWorld().Say()
}