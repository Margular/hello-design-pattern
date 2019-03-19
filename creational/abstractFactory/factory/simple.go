package factory

import "github.com/Margular/hello-design-pattern/creational/abstractFactory"

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