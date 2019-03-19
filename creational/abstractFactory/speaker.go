package abstractFactory

// STEP 1: Define some interfaces you want to use
type PositiveSpeaker interface {
	GoodSay()
}

type NegativeSpeaker interface {
	BadSay()
}