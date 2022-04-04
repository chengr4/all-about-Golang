package main

import "fmt"

type point struct {
	x int
	y int
}

type Person struct {
	Name string
	Age  int
}

func passPointer(a *point) {
	fmt.Println(a)
	fmt.Println(*a)
	fmt.Println(&a)
}

func test1() {
	p := Person{"Steve", 28}

	printPerson1(p) // String: {Steve 28} | Name: Steve | Age: 28
	updatePerson1(p)
	// no update
	printPerson1(p) // String: {Steve 28} | Name: Steve | Age: 28
}

func updatePerson1(p Person) {
	p.Age = 32
	printPerson1(p) // String: {Steve 32} | Name: Steve | Age: 32
}

func printPerson1(p Person) {
	fmt.Printf("String: %v | Name: %v | Age: %d\n",
		p,
		p.Name,
		p.Age)
}

// *****************************************************************************
// Test 2 - Pass by Reference
// *****************************************************************************

func test2() {
	p := &Person{"Steve", 28}
	printPerson2(p) // String: &{Steve 28} | Name: Steve | Age: 28
	updatePerson2(p)
	// updated
	printPerson2(p) // String: &{Steve 32} | Name: Steve | Age: 32
}

func updatePerson2(p *Person) {
	p.Age = 32
	printPerson2(p) // String: &{Steve 32} | Name: Steve | Age: 32
}

func printPerson2(p *Person) {
	fmt.Printf("String: %v | Name: %v | Age: %d\n",
		p,
		p.Name,
		p.Age)
}

// *****************************************************************************
// Test 3 - Pass by Reference (requires more typing)
// *****************************************************************************

func test3() {
	p := Person{"Steve", 28}
	printPerson3(&p) // String: &{Steve 28} | Name: Steve | Age: 28
	updatePerson3(&p)
	printPerson3(&p) // String: &{Steve 32} | Name: Steve | Age: 32
}

func updatePerson3(p *Person) {
	p.Age = 32
	printPerson3(p) // String: &{Steve 32} | Name: Steve | Age: 32
}

func printPerson3(p *Person) {
	fmt.Printf("String: %v | Name: %v | Age: %d\n",
		p,
		p.Name,
		p.Age)
}

// *****************************************************************************
// Test 4 - Pass by Value (requires more typing)
// *****************************************************************************

func test4() {
	p := &Person{"Steve", 28}
	printPerson4(*p) // {Steve 28} | Name: Steve | Age: 28
	updatePerson4(*p)
	printPerson4(*p) // {Steve 28} | Name: Steve | Age: 28
}

func updatePerson4(p Person) {
	p.Age = 32
	printPerson4(p) // {Steve 32} | Name: Steve | Age: 28
}

func printPerson4(p Person) {
	fmt.Printf("String: %v | Name: %v | Age: %d\n",
		p,
		p.Name,
		p.Age)
}

func main() {
	str1 := "eee"
	str2 := &str1
	fmt.Println(str2)
	p1 := &point{x: 3}
	fmt.Println("p1", &p1)
	passPointer(p1)
	test1()
	test2()
	test3()
	test4()
}
