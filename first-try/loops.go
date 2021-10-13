package main

// fmt: for formatting and printing
import "fmt"

// main(): the entry point of the application
func main() {
	x := 0

	// while loop
	for x < 5 {
		fmt.Println("value of x is ", x)
		x++
	}

	// traidtional for loop
	for i := 0; i < 5; i++ {
		fmt.Println("value of i is ", i)
	}

	names := []string{"apple", "banana", "cat", "dog"}

	// range (for..in)
	for index, value := range names {
		fmt.Printf("the position at index %v is %v \n", index, value)
	}

	// range (for..of)
	for _, value := range names {
		fmt.Printf("the position at is %v \n", value)
		// update the value and it does not affect original array
		value = "new string"
	}

	fmt.Println(names)

}
