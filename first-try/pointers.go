package main

import "fmt"

// func updateName(x string) {
// 	x = "yoshi"
// }

func updateName(x *string) {
	*x = "wedge"
}

func updateMenu(y map[string]float64) {
	y["coffee"] = 2.99
}

func main() {
	// group A types -> strings, ints, bools, floats, arrays, structs
	// non-pointer wrapper values
	name := "tifa"

	fmt.Println("memory address of name is: ", &name) // 0xc000096210
	mem := &name                                      // mem (pointer) stores memory location of name
	fmt.Println("memory address of name is: ", mem)   // 0xc000096210
	fmt.Println("value at memory address: ", *mem)

	fmt.Println(name) // tifa

	updateName(mem)
	fmt.Println(name) // widge

}
