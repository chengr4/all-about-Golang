// package "main" tells the go compiler to make the code to be executable program
package main

// fmt: for formatting and printing
import "fmt"

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// format the bill
func (b bill) format() string {
	fs := "Bill brealdown: \n"
	var total float64 = 0

	// list items
	for k, v := range b.items {
		fs += fmt.Sprintf("%v ... $%v \n", k+":", v)
		total += v
	}

	// total
	fs += fmt.Sprintf("%v ... $%0.2f", "total:", total)

	return fs
}

// make a new bills
func createBill(name string) bill {
	initBill := bill{
		name:  name,
		items: map[string]float64{"apple": 5.99, "cake": 3.99},
		tip:   0,
	}

	return initBill
}

// main(): the entry point of the application
func main() {
	myBill := createBill("Tifa's bill")

	fmt.Println(myBill)
	fmt.Println(myBill.format())

}
