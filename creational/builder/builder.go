package builder

// Pizza represents the final product being built
type Pizza struct {
	Size      string
	Cheese    bool
	Pepperoni bool
}

// PizzaBuilder provides step-by-step construction of a Pizza
type PizzaBuilder struct {
	size      string
	cheese    bool
	pepperoni bool
}

// SetSize sets the size of the pizza
func (pb *PizzaBuilder) SetSize(size string) *PizzaBuilder {
	pb.size = size
	return pb
}

// AddCheese adds cheese to the pizza
func (pb *PizzaBuilder) AddCheese() *PizzaBuilder {
	pb.cheese = true
	return pb
}

// AddPepperoni adds pepperoni to the pizza
func (pb *PizzaBuilder) AddPepperoni() *PizzaBuilder {
	pb.pepperoni = true
	return pb
}

// Build creates the final Pizza object
func (pb *PizzaBuilder) Build() Pizza {
	return Pizza{
		Size:      pb.size,
		Cheese:    pb.cheese,
		Pepperoni: pb.pepperoni,
	}
}
