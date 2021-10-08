package main

// fmt: for formatting and printing
import "fmt"

// main(): the entry point of the application
func main() {
	var ages [3]int = [3]int{20, 25, 30}
	// shorter way
	var ages1 = [3]int{20, 25, 30}
	fmt.Println(ages, ages1, len(ages1))

	// slices
	var scores = []int{100, 50, 60}
	scores[2] = 333

	// return a new slice for us than assaign to scores
	scores = append(scores, 3232)

	// slice ranges
	range1 := ages1[1:3] // 1 to 3 (exclude index 3)
	fmt.Println(range1)
}
