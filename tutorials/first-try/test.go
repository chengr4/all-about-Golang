package main

// fmt: for formatting and printing
import (
	"fmt"
)

// return multiple values
func testPrimitive(n string) {
	s := n
	s = "fefe"
	fmt.Println(s, n)
}

// main(): the entry point of the application
func main() {
	testPrimitive("3")
}
