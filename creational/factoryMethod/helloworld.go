package factoryMethod

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