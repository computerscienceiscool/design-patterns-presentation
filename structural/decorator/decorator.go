package decorator

import "fmt"

// Coffee defines the interface for the base component
type Coffee interface {
	Cost() int
	Description() string
}

// SimpleCoffee is a concrete component implementing Coffee
type SimpleCoffee struct{}

func (s *SimpleCoffee) Cost() int {
	return 5
}

func (s *SimpleCoffee) Description() string {
	return "Simple Coffee"
}

// MilkDecorator is a decorator adding milk to the coffee
type MilkDecorator struct {
	coffee Coffee
}

func (m *MilkDecorator) Cost() int {
	return m.coffee.Cost() + 2
}

func (m *MilkDecorator) Description() string {
	return m.coffee.Description() + ", Milk"
}

// SugarDecorator is a decorator adding sugar to the coffee
type SugarDecorator struct {
	coffee Coffee
}

func (s *SugarDecorator) Cost() int {
	return s.coffee.Cost() + 1
}

func (s *SugarDecorator) Description() string {
	return s.coffee.Description() + ", Sugar"
}

// DemoDecorator demonstrates the decorator pattern
func DemoDecorator() {
	coffee := &SimpleCoffee{}
	fmt.Println(coffee.Description(), "costs", coffee.Cost())

	milkCoffee := &MilkDecorator{coffee: coffee}
	fmt.Println(milkCoffee.Description(), "costs", milkCoffee.Cost())

	sugarMilkCoffee := &SugarDecorator{coffee: milkCoffee}
	fmt.Println(sugarMilkCoffee.Description(), "costs", sugarMilkCoffee.Cost())
}
