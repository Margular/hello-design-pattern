package builder

// STEP 2: Define the builder you want to use to build the struct we've defined before
type HelloWorldBuilder struct {
	hello string
	world string
}

func (b *HelloWorldBuilder) SetHello(hello string) *HelloWorldBuilder {
	b.hello = hello
	return b
}

func (b *HelloWorldBuilder) SetWorld(world string) *HelloWorldBuilder {
	b.world = world
	return b
}

func (b *HelloWorldBuilder) Build() HelloWorld {
	return HelloWorld{
		b.hello,
		b.world,
	}
}