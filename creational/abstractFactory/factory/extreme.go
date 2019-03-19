package factory

import "github.com/Margular/hello-design-pattern/creational/abstractFactory"

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