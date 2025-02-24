package strategy

import "fmt"

// TransportStrategy defines an interface for different transport methods.
type TransportStrategy interface {
	Travel()
}

// CarStrategy is a concrete strategy for traveling by car.
type CarStrategy struct{}

func (c *CarStrategy) Travel() { fmt.Println("Driving a car") }

// BikeStrategy is a concrete strategy for traveling by bike.
type BikeStrategy struct{}

func (b *BikeStrategy) Travel() { fmt.Println("Riding a bike") }

// DemoStrategy demonstrates the strategy pattern.
func DemoStrategy() {
	var strategy TransportStrategy = &CarStrategy{}
	strategy.Travel() // Driving a car

	strategy = &BikeStrategy{}
	strategy.Travel() // Riding a bike
}
