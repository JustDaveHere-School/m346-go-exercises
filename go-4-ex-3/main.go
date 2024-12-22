package main

import (
	"fmt"
	"math"
)

func computeQuadraticFormula(a, b, c float64) []float64 {
	discriminant := math.Pow(b, 2) - 4*a*c

	if discriminant > 0 {
		// Zwei Lösungen
		root1 := (-b + math.Sqrt(discriminant)) / (2 * a)
		root2 := (-b - math.Sqrt(discriminant)) / (2 * a)
		return []float64{root1, root2}
	} else if discriminant == 0 {
		// Eine Lösung
		root := -b / (2 * a)
		return []float64{root}
	} else {
		// Keine Lösung
		return []float64{}
	}
}

func main() {
	// Testfall 1: Zwei Lösungen (D > 0)
	fmt.Println(computeQuadraticFormula(3, 4, 1)) // [ -0.333, -1.0 ]

	// Testfall 2: Eine Lösung (D == 0)
	fmt.Println(computeQuadraticFormula(2, 4, 2)) // [ -1.0 ]

	// Testfall 3: Keine Lösung (D < 0)
	fmt.Println(computeQuadraticFormula(3, 4, 2)) // []
}
