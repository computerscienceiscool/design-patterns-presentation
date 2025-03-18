
class: center, middle

# Design Patterns in Software Development
## Understanding Key Patterns for Scalable Code

---

# Why Design Patterns Matter

- Software development is complex.
- Patterns provide **proven solutions** to common problems.
- Used across different languages and platforms.
- Make code **scalable, maintainable, and reusable**.

---

# The Gang of Four (GoF) Design Patterns

- In 1994, four developers (**Erich Gamma, Richard Helm, Ralph Johnson, John Vlissides**) formalized 23 design patterns.
- These patterns fall into **three categories**:
  - **Creational Patterns** – How objects are created.
  - **Structural Patterns** – How objects relate.
  - **Behavioral Patterns** – How objects communicate.

---

# Categories of Design Patterns

- **Creational Patterns** → Control how objects are created.
- **Structural Patterns** → Define relationships between objects.
- **Behavioral Patterns** → Manage communication between objects.

---

# Creational Patterns

- Focus on **object creation**.
- Example: **Building with Legos** – Creational patterns define what types of Legos you have.
- We’ll cover:
  - Singleton
  - Factory Method
  - Builder

---


# Singleton Pattern

Ensures a class has **only one instance** and provides a global access point.
Used in logging systems, database connections, and configuration managers.

 

---

# Singleton Pattern

Ensures a class has **only one instance** and provides a global access point.

Example: **Logging System**

```go
package main

import (
    "fmt"
    "sync"
)

type Logger struct{}

var instance *Logger
var once sync.Once

func GetLoggerInstance() *Logger {
    once.Do(func() {
        instance = &Logger{}
    })
    return instance
}

func main() {
    logger1 := GetLoggerInstance()
    logger2 := GetLoggerInstance()
    fmt.Println(logger1 == logger2) // true
}
```

---

# Factory Method Pattern

Encapsulates object creation logic so that calling code does not need to know the specific class being instantiated.
Commonly used in frameworks and libraries to provide extensibility.


---


# Factory Method Pattern

Encapsulates object creation logic to avoid direct instantiation.

Example: **User Roles in an Application**

```go
package main

import "fmt"

type User interface {
    GetRole() string
}

type Admin struct{}
func (a Admin) GetRole() string { return "Admin" }

type Guest struct{}
func (g Guest) GetRole() string { return "Guest" }

func UserFactory(userType string) User {
    if userType == "Admin" {
        return Admin{}
    }
    return Guest{}
}

func main() {
    user := UserFactory("Admin")
    fmt.Println(user.GetRole()) // Admin
}
```

---

# Builder Pattern

Used for **step-by-step object construction**, especially when many optional parameters exist.
 

---


# Builder Pattern
Example: **Building a Pizaa**

```go
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
```

---

# Structural Patterns

- Focus on **how objects fit together**.
- Example: **Building with Legos** – Structural patterns define how the pieces connect.
- We’ll cover:
  - Facade
  - Adapter

---


# Facade Pattern

Simplifies complex systems by providing a **unified interface** to underlying subsystems.
Used in APIs, smart home systems, and large enterprise applications.


---


# Facade Pattern

Simplifies interactions by providing a **unified interface**.

Example: **E-commerce Checkout**

```go
package main

import "fmt"

type PaymentProcessor struct{}
func (p *PaymentProcessor) ProcessPayment() { fmt.Println("Processing payment...") }

type InventoryChecker struct{}
func (i *InventoryChecker) CheckInventory() { fmt.Println("Checking inventory...") }

type OrderFacade struct {
    payment  PaymentProcessor
    inventory InventoryChecker
}

func (o *OrderFacade) PlaceOrder() {
    o.inventory.CheckInventory()
    o.payment.ProcessPayment()
    fmt.Println("Order placed successfully!")
}

func main() {
    order := &OrderFacade{}
    order.PlaceOrder()
}
```

---


# Adapter Pattern

Allows incompatible interfaces to work together by acting as a **bridge**.
Used for hardware adapters, third-party API integrations, and data format conversions.


---


# Adapter Pattern

Converts one interface into another to enable compatibility.

Example: **Converting Fahrenheit to Celsius**

```go
package main

import "fmt"

type CelsiusSensor struct{}
func (c *CelsiusSensor) GetTemperature() float64 { return 25.0 }

type FahrenheitAdapter struct {
    sensor *CelsiusSensor
}

func (fa *FahrenheitAdapter) GetTemperature() float64 {
    return (fa.sensor.GetTemperature() * 9/5) + 32
}

func main() {
    sensor := &CelsiusSensor{}
    adapter := &FahrenheitAdapter{sensor}
    fmt.Println("Temperature in Fahrenheit:", adapter.GetTemperature()) // 77°F
}
```

---

# Behavioral Patterns

- Focus on **how objects communicate**.
- Example: **Playing with Legos** – Behavioral patterns define interaction rules.
- We’ll cover:
  - Strategy
  - Observer

---

# Strategy Pattern

Encapsulates interchangeable behaviors, allowing dynamic selection of an appropriate algorithm.
Used in sorting algorithms, payment methods, and navigation apps.

---


# Strategy Pattern

Encapsulates interchangeable behaviors.

Example: **Transport Strategy**

```go
package main

import "fmt"

type TransportStrategy interface {
    Travel()
}

type CarStrategy struct{}
func (c *CarStrategy) Travel() { fmt.Println("Driving a car") }

type BikeStrategy struct{}
func (b *BikeStrategy) Travel() { fmt.Println("Riding a bike") }

func main() {
    var strategy TransportStrategy = &CarStrategy{}
    strategy.Travel() // Driving a car

    strategy = &BikeStrategy{}
    strategy.Travel() // Riding a bike
}
```

---

# Observer Pattern

Defines a one-to-many dependency where objects (observers) react to changes in a subject.
Used in event-driven systems, social media notifications, and stock market updates.


---


# Observer Pattern

Allows objects to subscribe to and receive updates when an event occurs.

Example: **YouTube Subscription Notifications**

```go
package main

import "fmt"

type Observer interface {
    Update(videoTitle string)
}

type Subscriber struct {
    name string
}

func (s *Subscriber) Update(videoTitle string) {
    fmt.Printf("%s received notification for new video: %s
", s.name, videoTitle)
}

type YouTubeChannel struct {
    observers []Observer
}

func (yc *YouTubeChannel) Subscribe(observer Observer) {
    yc.observers = append(yc.observers, observer)
}

func (yc *YouTubeChannel) Notify(videoTitle string) {
    for _, observer := range yc.observers {
        observer.Update(videoTitle)
    }
}

func main() {
    channel := &YouTubeChannel{}
    sub1 := &Subscriber{name: "Alice"}
    channel.Subscribe(sub1)

    channel.Notify("New Go Tutorial")
}
```

---


# List of Design Patterns

| **Creational Patterns  **      | **  Structural Patterns  ** | **  Behavioral Patterns  **  |
|------------------------------|------------------------|--------------------------|
| Singleton                    | Adapter                | Chain of Responsibility |
| Factory Method               | Bridge                 | Command                 |
| Abstract Factory             | Composite              | Interpreter             |
| Builder                      | Decorator              | Iterator                |
| Prototype                    | Facade                 | Mediator                |
|                              | Flyweight              | Memento                 |
|                              | Proxy                  | Observer                |
|                              |                        | State                   |
|                              |                        | Strategy                |
|                              |                        | Template Method         |
|                              |                        | Visitor                 |



---


# Summary: 7 Key Design Patterns

- **Singleton** → Ensures a single instance with global access.  
- **Factory Method** → Creates objects dynamically based on input.  
- **Builder** → Constructs complex objects step by step.  
- **Facade** → Provides a simple interface to a complex system.  
- **Adapter** → Bridges incompatible interfaces.  
- **Strategy** → Allows dynamic selection of behaviors.  
- **Observer** → Notifies multiple subscribers of changes.  

---


# Quick Quiz: Match the Patterns!

Can you match each pattern to a real-world scenario?  
We’ll discuss together!

1. **Singleton**
2. **Factory Method**
3. **Builder**
4. **Facade**
5. **Adapter**
6. **Strategy**
7. **Observer**


---

# Thank You!

- Questions?
- Thoughts?

