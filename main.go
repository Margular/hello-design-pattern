package main

import (
	"github.com/Margular/hello-design-pattern/creational/abstractFactory"
	"github.com/Margular/hello-design-pattern/creational/abstractFactory/factory"
)

func tryAbstractFactory(speakerFactory abstractFactory.SpeakerFactory) {
	positiveSpeaker := speakerFactory.CreatePositiveSpeaker()
	negativeSpeaker := speakerFactory.CreateNegativeSpeaker()
	positiveSpeaker.GoodSay()
	negativeSpeaker.BadSay()
}

func main() {
	println("We are creational patterns!\n")

	println("1. Abstract Factory: ")
	tryAbstractFactory(factory.SimpleSpeakerFactory{})
	tryAbstractFactory(factory.ExtremeSpeakerFactory{})
	println()

	// TODO:
}