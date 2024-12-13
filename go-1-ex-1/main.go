package main

import "fmt"

func main() {
	firstName, lastName := "Vorname", "Nachname"
	dayOfBirth, monthOfBirth, yearOfBirth := 21, 7, 2006
	numberOfSiblings := 1
	heightInMeters := 1.83
	zodiacSign := '\u264B'

	fmt.Printf("Vor- und Nachname: %s %s\n", firstName, lastName)
	fmt.Printf("Geburtsdatum: %d.%d.%d\n", dayOfBirth, monthOfBirth, yearOfBirth)
	fmt.Printf("Anzahl Geschwister: %d\n", numberOfSiblings)
	fmt.Printf("Grösse (in Metern): %.2f\n", heightInMeters)
	fmt.Printf("Sternzeichen: %c\n", zodiacSign)
}
