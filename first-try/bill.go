package main

/*
this file includes concepts: input, switch, parsing floats, saving files
*/

// fmt: for formatting and printing
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

// add item to bill
func (b *bill) addItem(name string, price float64) {
	b.items[name] = price
}

// update tip
func (b *bill) updateTip(tip float64) {
	// here will dereference itself automatically
	b.tip = tip
}

// format the bill
func (b *bill) format() string {
	fs := "Bill brealdown: \n"
	var total float64 = 0

	// list items
	for k, v := range b.items {
		fs += fmt.Sprintf("%v ... $%v \n", k+":", v)
		total += v
	}

	fs += fmt.Sprintf("%v ... $%v\n", "tip:", b.tip)

	// total
	fs += fmt.Sprintf("%v ... $%0.2f\n", "total:", total+b.tip)

	return fs
}

// save bill
func (b *bill) saveBill() {
	data := []byte(b.format())

	// 0644: permission
	err := os.WriteFile("bills/"+b.name+".txt", data, 0644)
	if err != nil {
		// stop the flow
		panic(err)
	}
	fmt.Println("Bill was saved")
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

		p, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("THe price must be a number")
			promptOptions(b)
		}
		b.addItem(name, p)

		fmt.Println("name: ", name, ", price: ", price)
		promptOptions(b)

	case "t":
		fmt.Println("you chose t")
		tip, _ := getInput("Enter tip amount ($): ", reader)
		t, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("THe price must be a number")
			promptOptions(b)
		}
		b.updateTip(t)

		fmt.Println("tip added - ", tip)
		promptOptions(b)

	case "s":
		fmt.Println("you chose s to save the bill")
		fmt.Println(b.format())
		b.saveBill()

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
