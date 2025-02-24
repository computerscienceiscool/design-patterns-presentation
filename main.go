package main

import (
	"design-patterns-presentation/behavioral/observer"
	"design-patterns-presentation/behavioral/strategy"
	"design-patterns-presentation/creational/builder"
	"design-patterns-presentation/creational/factory"
	"design-patterns-presentation/creational/singleton"
	"design-patterns-presentation/structural/adapter"
	"design-patterns-presentation/structural/facade"
)

func main() {
	println("Running Design Patterns Demo:")

	println("\n--- Creational Patterns ---")
	singleton.DemoSingleton()
	factory.DemoFactory()
	builder.DemoBuilder()

	println("\n--- Structural Patterns ---")
	facade.DemoFacade()
	adapter.DemoAdapter()

	println("\n--- Behavioral Patterns ---")
	strategy.DemoStrategy()
	observer.DemoObserver()
}
