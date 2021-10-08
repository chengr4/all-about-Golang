package main

// fmt: for formatting and printing
import "fmt"

// main(): the entry point of the application
func main() {

	age := 35
	name := "Tifa"

	// print
	fmt.Print("Hello, 世界")

	// println
	fmt.Println("age: ", age, ",name: ", name)

	// format string
	// %v = variable
	// %q = add quote
	// %T = type of variable
	fmt.Printf("age: %v, name: %v\n", age, name)
	fmt.Printf("age: %q, name: %q\n", age, name) // age: '#', name: "Tifa"
	fmt.Printf("age type: %T\n", age)            // age type: int
	fmt.Printf("float: %0.1f\n", 222.4)

	// sprintf (save formatted strings)
	var str = fmt.Sprintf("age: %v, name: %v\n", age, name)
	fmt.Println(str)

}
