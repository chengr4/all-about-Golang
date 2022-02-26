package main

import "fmt"

// func updateName(x string) {
// 	x = "yoshi"
// }

// x is a copy of the argument not argument itself
func updateName(x string) string {
	x = "wedge"
	return x
}

func updateMenu(y map[string]float64) {
	y["coffee"] = 2.99
}

func main() {
	// group A types -> strings, ints, bools, floats, arrays, structs
	// non-pointer wrapper values
	name := "tifa"

	updateName(name)
	fmt.Println("1: ", name) // tifa

	name = updateName(name)
	fmt.Println("2: ", name) // wedge

	// group B types -> slices, maps, functions
	// pointer wrapper values
	menu := map[string]float64{
		"pie":       5.95,
		"ice cream": 3.99,
	}

	updateMenu(menu)
	fmt.Println(menu) //map[coffee:2.99 ice cream:3.99 pie:5.95]
}
