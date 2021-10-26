package main

/*
this file includes concepts: input, switch
*/

// fmt: for formatting and printing
import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// make new bills
func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}
	return b
}

func getInput(prompt string, reader *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := reader.ReadString('\n')

	return strings.TrimSpace(input), err
}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)
	// fmt.Print("Create a new bill name: ")
	// name, _ := reader.ReadString('\n')
	// name = strings.TrimSpace(name)

	// instead of above
	name, _ := getInput("Create a new bill name: ", reader)

	b := newBill(name)
	fmt.Println("Created the bill - ", b.name)

	return b
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)
	option, _ := getInput("Choose an option (a: add a item, s: save the bill, t: add the tip): ", reader)

	// switch
	switch option {
	case "a":
		fmt.Println("you chose a")
		name, _ := getInput("Item name: ", reader)
		price, _ := getInput("Item price: ", reader)

		fmt.Println("name: ", name, ", price: ", price)
	case "t":
		fmt.Println("you chose t")
		tip, _ := getInput("Enter tip amount ($)", reader)

		fmt.Println("tip: ", tip)
	case "s":
		fmt.Println("you chose s")
	default:
		fmt.Println("that was not a vaild option...")
		promptOptions(b)
	}
}

// main(): the entry point of the application
func main() {
	myBill := createBill()
	promptOptions(myBill)
}
