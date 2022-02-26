package main

// fmt: for formatting and printing
import "fmt"

// main(): the entry point of the application
func main() {
	meun := map[string]float64{
		"soup":  123.2,
		"apple": 444,
		"pie":   4322.2,
	}

	fmt.Println(meun)
	fmt.Println(meun["pie"])

	for k, v := range meun {
		fmt.Println(k, v)
	}

	// int as key type
	phonebook := map[int]string{
		34143242:     "feaffe",
		34143222242:  "fefdfdaaffe",
		344343143242: "feawwwwwffe",
	}

	fmt.Println(phonebook)
	fmt.Println(phonebook[34143242])

	phonebook[34143242] = "apple"
	fmt.Println(phonebook[34143242])

}
