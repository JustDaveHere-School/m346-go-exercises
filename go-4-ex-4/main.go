package main

import "fmt"

// convertCelsiusToFahrenheit rechnet Celsius in Fahrenheit um.
func convertCelsiusToFahrenheit(celsius float64) float64 {
	return celsius*(9.0/5.0) + 32
}

// convertFahrenheitToCelsius rechnet Fahrenheit in Celsius um.
func convertFahrenheitToCelsius(fahrenheit float64) float64 {
	return (fahrenheit - 32) * (5.0 / 9.0)
}

func main() {
	// Beispielaufrufe für Celsius zu Fahrenheit
	f1 := convertCelsiusToFahrenheit(0)   // 32
	f2 := convertCelsiusToFahrenheit(25)  // 77
	f3 := convertCelsiusToFahrenheit(100) // 212

	fmt.Println(f1) // 32
	fmt.Println(f2) // 77
	fmt.Println(f3) // 212

	// Rückrechnung von Fahrenheit zu Celsius
	c1 := convertFahrenheitToCelsius(f1) // 0
	c2 := convertFahrenheitToCelsius(f2) // 25
	c3 := convertFahrenheitToCelsius(f3) // 100

	fmt.Println(c1) // 0
	fmt.Println(c2) // 25
	fmt.Println(c3) // 100
}
