package main

import "fmt"

func computeGrade(gotPoints, maxPoints float64) float64 {
	return (gotPoints/maxPoints)*5 + 1
}

func main() {
	fmt.Println(computeGrade(17.5, 28.0)) // 4.125
	fmt.Println(computeGrade(23.0, 28.0)) // 5.107
	fmt.Println(computeGrade(10.0, 28.0)) // 2.786
}
