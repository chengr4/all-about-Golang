package main

// fmt: for formatting and printing
import "fmt"

// main(): the entry point of the application
func main() {
	age := 45

	if age < 30 {
		fmt.Println(" age is less than 30")
	}

	fmt.Println(" age is not less than 30")

	names := []string{"apple", "banana", "cat", "dog"}

	for i, v := range names {
		if i == 1 {
			fmt.Println("in if: ", v)
			continue
		}
		fmt.Println(v)
		if i > 3 {
			fmt.Println("breaking")
			break
		}
	}
}
