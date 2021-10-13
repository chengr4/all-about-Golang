package main

// fmt: for formatting and printing
import (
	"fmt"
	"math"
)

func sayYa(n string) {
	fmt.Println(n)
}

func cycleNames(n []string, f func(string)) {
	for _, v := range n {
		f(v)
	}
}

func circleArea(r float64) float64 {
	return math.Pi * r * r
}

// main(): the entry point of the application
func main() {
	sayYa("123")
	cycleNames([]string{"cloud", "tifa", "barret"}, sayYa)

	a1 := circleArea(14.5)
	a2 := circleArea(133.5)
	fmt.Println(a1, a2)
	fmt.Printf("%0.2f, %0.3f\n", a1, a2)
}
