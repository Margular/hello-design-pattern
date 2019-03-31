package main

import (
	"fmt"
	"github.com/Margular/hello-design-pattern/behavioral/chainOfResponsibility"
	"github.com/Margular/hello-design-pattern/creational/abstractFactory"
	"github.com/Margular/hello-design-pattern/creational/abstractFactory/factory"
	"github.com/Margular/hello-design-pattern/creational/builder"
	"github.com/Margular/hello-design-pattern/creational/factoryMethod"
	"github.com/Margular/hello-design-pattern/creational/prototype"
	"github.com/Margular/hello-design-pattern/creational/singleton"
	"github.com/Margular/hello-design-pattern/structural/adapter"
	"github.com/Margular/hello-design-pattern/structural/bridge"
	"github.com/Margular/hello-design-pattern/structural/composite"
	"github.com/Margular/hello-design-pattern/structural/decorator"
	"github.com/Margular/hello-design-pattern/structural/facade"
	"github.com/Margular/hello-design-pattern/structural/flyweight"
	"github.com/Margular/hello-design-pattern/structural/proxy"
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

	fmt.Println("5. Singleton: ")
	singleton.HelloWorld().Say()
	singleton.HelloWorld().Say()

	fmt.Println("\nWe are structural patterns!\n")

	fmt.Println("6. Adapter: ")
	var target adapter.Speaker
	target = adapter.Adapter{adapter.HelloWorld{}}
	target.Say()
	fmt.Println()

	fmt.Println("7. Bridge: ")
	var tom = bridge.Person{}
	tom.M = bridge.HelloMouth{}
	tom.Speak()
	fmt.Println()

	fmt.Println("8. Composite: ")
	t := composite.Tom{}
	j := composite.Jerry{}
	world := composite.World{
		[]composite.Speaker{
			t,
			j,
		},
	}
	world.Say()
	fmt.Println()

	fmt.Println("9. Decorator: ")
	var dec decorator.PersonDecorator
	dec = decorator.AnoisingPerson{decorator.Tom{}}
	dec.Say()
	fmt.Println()

	fmt.Println("10. Facade: ")
	var house facade.House
	house.AllSay()
	fmt.Println()

	fmt.Println("11. Flyweight: ")
	flyweight.AnimalFactory().GetAnimal("Tom").Say()
	flyweight.AnimalFactory().GetAnimal("Tom").Say()
	flyweight.AnimalFactory().GetAnimal("Jerry").Say()
	flyweight.AnimalFactory().GetAnimal("Tom").Say()
	flyweight.AnimalFactory().GetAnimal("Jerry").Say()
	fmt.Println()

	fmt.Println("12. Proxy: ")
	var p proxy.Subject
	p = proxy.Proxy{}
	p.Request()
	fmt.Println()

	fmt.Println("We are behavioral patterns!\n")

	fmt.Println("13. Chain of Responsibility: ")
	peter := chainOfResponsibility.Person{Name : "Peter"}
	fox := chainOfResponsibility.Person{Name : "Fox", Successor : &peter}
	steven := chainOfResponsibility.Person{Name : "Steven", Successor : &fox}
	var theWorld chainOfResponsibility.PersonHandler = chainOfResponsibility.Person{Successor : &steven}

	theWorld.HandleRequest("Peter")
	theWorld.HandleRequest("Fox")
	theWorld.HandleRequest("Tom")
	fmt.Println()


}