package main

import (
	"fmt"
	"github.com/Margular/hello-design-pattern/creational/abstractFactory"
	"github.com/Margular/hello-design-pattern/creational/abstractFactory/factory"
	"github.com/Margular/hello-design-pattern/creational/builder"
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

	fmt.Println("2. Builder")
	hwBuilder := builder.HelloWorldBuilder{}
	hw := hwBuilder.SetHello("hello").SetWorld("world").Build()
	hw.Say()

	// TODO:
}