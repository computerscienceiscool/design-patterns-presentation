package adapter

import "fmt"

// CelsiusSensor represents a temperature sensor providing readings in Celsius.
type CelsiusSensor struct{}

// GetTemperature returns the temperature in Celsius.
func (c *CelsiusSensor) GetTemperature() float64 {
	return 25.0
}

// FahrenheitAdapter adapts CelsiusSensor to provide temperature in Fahrenheit.
type FahrenheitAdapter struct {
	sensor *CelsiusSensor
}

// GetTemperature converts Celsius to Fahrenheit.
func (fa *FahrenheitAdapter) GetTemperature() float64 {
	return (fa.sensor.GetTemperature() * 9 / 5) + 32
}

// DemoAdapter demonstrates the usage of the adapter pattern.
func DemoAdapter() {
	sensor := &CelsiusSensor{}
	adapter := &FahrenheitAdapter{sensor}
	fmt.Println("Temperature in Fahrenheit:", adapter.GetTemperature()) // 77°F
}
