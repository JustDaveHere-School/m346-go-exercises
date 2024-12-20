package main

import "fmt"

const (
	Aries       = '\u2648' // Widder
	Taurus      = '\u2649' // Stier
	Gemini      = '\u264a' // Zwillinge
	Cancer      = '\u264b' // Krebs
	Leo         = '\u264c' // Löwe
	Virgo       = '\u264d' // Jungfrau
	Libra       = '\u264e' // Waage
	Scorpius    = '\u264f' // Skorpion
	Sagittarius = '\u2650' // Schütze
	Capricornus = '\u2651' // Steinbock
	Aquarius    = '\u2652' // Wassermann
	Pisces      = '\u2653' // Fische
	Sun         = '\u2600' // Sun Symbol (non-zodiac)
)

func outputDateRange(zodiacSign rune) {
	fmt.Printf("%c: ", zodiacSign)
	// TODO: Replace if, else if branching with switch/case.
	// TODO: Define all 12 cases...
	switch zodiacSign {
	case Aries:
		fmt.Println("March 21 - April 19")
	case Taurus:
		fmt.Println("April 20 - May 20")
	case Gemini:
		fmt.Println("May 21 - June 20")
	case Cancer:
		fmt.Println("June 21 - July 22")
	case Leo:
		fmt.Println("July 23 - August 22")
	case Virgo:
		fmt.Println("August 23 - September 22")
	case Libra:
		fmt.Println("September 23 - October 22")
	case Scorpius:
		fmt.Println("October 23 - November 21")
	case Sagittarius:
		fmt.Println("November 22 - December 21")
	case Capricornus:
		fmt.Println("December 22 - January 19")
	case Aquarius:
		fmt.Println("January 20 - February 18")
	case Pisces:
		fmt.Println("February 19 - March 20")
	default:
		fmt.Println("Invalid Zodiac Sign")
	}
	// TODO: ...and consider a default case.
}

func main() {
	for zodiacSign := Aries; zodiacSign <= Pisces; zodiacSign++ {
		outputDateRange(zodiacSign)
	}
	outputDateRange(Sun)
}
