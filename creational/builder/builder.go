
package main

import "fmt"

// Product: The object being built
type Pizza struct {
    size   string
    cheese bool
    pepperoni bool
}

// Builder: Defines the step-by-step construction
type PizzaBuilder struct {
    size   string
    cheese bool
    pepperoni bool
}

func (pb *PizzaBuilder) SetSize(size string) *PizzaBuilder {
    pb.size = size
    return pb
}

func (pb *PizzaBuilder) AddCheese() *PizzaBuilder {
    pb.cheese = true
    return pb
}

func (pb *PizzaBuilder) AddPepperoni() *PizzaBuilder {
    pb.pepperoni = true
    return pb
}

func (pb *PizzaBuilder) Build() Pizza {
    return Pizza{
        size:   pb.size,
        cheese: pb.cheese,
        pepperoni: pb.pepperoni,
    }
}

func main() {
    // Using the builder to create a customized pizza
    pizza := PizzaBuilder{}.
        SetSize("Large").
        AddCheese().
        AddPepperoni().
        Build()

    fmt.Printf("Pizza: %s, Cheese: %v, Pepperoni: %v\n", pizza.size, pizza.cheese, pizza.pepperoni)
}

