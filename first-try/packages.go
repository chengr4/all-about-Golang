package main

// fmt: for formatting and printing
import (
	"fmt"
	"sort"
	"strings"
)

// main(): the entry point of the application
func main() {
	greeting := "Hi Hi"

	// strings
	fmt.Println(strings.Contains(greeting, "H"))           // true if pattern is in the string
	fmt.Println(strings.ReplaceAll(greeting, "Hi", "bad")) // it copied and replaced
	fmt.Println(strings.ToUpper(greeting))
	fmt.Println(strings.Split(greeting, " "))

	//
	ages := []int{45, 20, 35, 30, 75, 60}

	sort.Ints(ages)
	fmt.Println(ages)
}
