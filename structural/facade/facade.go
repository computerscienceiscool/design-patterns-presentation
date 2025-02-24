package facade

import "fmt"

// PaymentProcessor handles payment processing
type PaymentProcessor struct{}

func (p *PaymentProcessor) ProcessPayment() { fmt.Println("Processing payment...") }

// InventoryChecker checks inventory status
type InventoryChecker struct{}

func (i *InventoryChecker) CheckInventory() { fmt.Println("Checking inventory...") }

// OrderFacade provides a simple interface for placing an order
type OrderFacade struct {
	payment   PaymentProcessor
	inventory InventoryChecker
}

func (o *OrderFacade) PlaceOrder() {
	o.inventory.CheckInventory()
	o.payment.ProcessPayment()
	fmt.Println("Order placed successfully!")
}

// DemoFacade demonstrates the facade pattern
func DemoFacade() {
	order := &OrderFacade{}
	order.PlaceOrder()
}
