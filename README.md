# Design Patterns Code Demo Cheat Sheet

## Creational Patterns

### Singleton Pattern (`singleton.go`)
- **What it does:** Ensures that only one instance of a class exists and provides a global access point.
- **How it works:**
  - Uses `sync.Once` to ensure the instance is created only once.
  - The `GetLoggerInstance()` function returns the same instance every time.
- **Output Explanation:**
  - Prints `true`, confirming that both variables hold the same instance.

### Factory Pattern (`factory.go`)
- **What it does:** Creates objects without specifying the exact class that will be instantiated.
- **How it works:**
  - Defines an interface `User`.
  - Uses `UserFactory(userType string)` to create either an `Admin` or `Guest`.
- **Output Explanation:**
  - Prints `Admin`, demonstrating that the factory successfully created an Admin user.

### Builder Pattern (`builder.go`)
- **What it does:** Constructs complex objects step by step.
- **How it works:**
  - `RequestBuilder` allows setting `method` and `url` separately.
  - Calls `SetMethod("GET")` and `SetURL("https://example.com")` before `Build()`.
- **Output Explanation:**
  - Prints `GET https://example.com`, showing that the object was built correctly.

---

## Structural Patterns

### Facade Pattern (`facade.go`)
- **What it does:** Provides a simplified interface to a larger body of code.
- **How it works:**
  - `OrderFacade` calls `CheckInventory()` and `ProcessPayment()` internally.
  - The client only calls `PlaceOrder()` instead of interacting with multiple subsystems.
- **Output Explanation:**
  - Shows messages for checking inventory, processing payment, and order success, hiding internal details.

### Adapter Pattern (`adapter.go`)
- **What it does:** Converts one interface into another that the client expects.
- **How it works:**
  - `CelsiusSensor` provides temperature in Celsius.
  - `FahrenheitAdapter` converts it to Fahrenheit.
- **Output Explanation:**
  - Prints `Temperature in Fahrenheit: 77`, showing the adapter in action.

---

## Behavioral Patterns

### Strategy Pattern (`strategy.go`)
- **What it does:** Allows selecting an algorithm at runtime.
- **How it works:**
  - Defines `TransportStrategy` interface.
  - Implements `CarStrategy` and `BikeStrategy`.
  - Switches between strategies dynamically.
- **Output Explanation:**
  - First prints `Driving a car`, then `Riding a bike`, demonstrating the strategy change.

### Observer Pattern (`observer.go`)
- **What it does:** Allows objects to subscribe and receive updates when an event occurs.
- **How it works:**
  - `YouTubeChannel` maintains a list of subscribers.
  - Calls `Notify("New Go Tutorial")` to alert all subscribers.
- **Output Explanation:**
  - Prints `Alice received notification for new video: New Go Tutorial`, showing the observer in action.

---

## Bonus: How `main.go` Runs Everything
- Calls each pattern demo one by one.
- Groups them under `Creational`, `Structural`, and `Behavioral` sections.
- Uses `fmt.Println("🔹 Pattern Name:")` to label outputs for clarity.

🚀 **Use this cheat sheet to explain each pattern during your presentation!**

