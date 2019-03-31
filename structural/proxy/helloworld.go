package proxy

import "fmt"

// STEP 1: Define the subject interface
type Subject interface {
	Request()
}

// STEP 2: Define a real subject
type TomSubject struct {}

func (t TomSubject) Request() {
	fmt.Println("Hello World! I'm Tom")
}

// STEP 3: Define a proxy that implements the subject and have a real subject member
type Proxy struct {
	tom *TomSubject
}

func (p *Proxy) preRequest() {
	if p.tom == nil {
		p.tom = &TomSubject{}
	}
}

func (p Proxy) Request() {
	p.preRequest()
	p.tom.Request()
}