// You can edit this code!
// Click here and start typing.
// package "main" tells the go compiler to make the code to be executable program
package main

// fmt: for formatting and printing
import "fmt"

// main(): the entry point of the application
func main() {

	// string
	// string must use double quote
	var name1 string = "mario"
	var name2 = "fe"
	var name3 string

	fmt.Println(name1, name2, name3)

	// can only be existed in the function
	name4 := "yoshi"
	fmt.Println(name4)

	// int
	var age1 int = 20
	var age2 = 20
	age3 := 20

	fmt.Println(age1, age2, age3)

	// bit & memory
	// var numError int8 = 215 // error: int range only from -128 to 127
	var num1 int8 = 15  // int range only from -128 to 127
	var num2 uint = 25  // no negative number
	var num3 uint8 = 25 // no negative number

	// floatf
	var score1 float32 = 25.98
	var score2 float64 = 25324324234.98
	score3 := 3213.4343

	fmt.Println(num1, num2, num3, score1, score2, score3)
}
