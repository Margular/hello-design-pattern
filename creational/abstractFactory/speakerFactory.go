package abstractFactory

// STEP 2: Define the abstract factory, it commonly has some methods to create some instances of interface
// that you have defined in STEP 1
type SpeakerFactory interface {
	CreatePositiveSpeaker() PositiveSpeaker
	CreateNegativeSpeaker() NegativeSpeaker
}