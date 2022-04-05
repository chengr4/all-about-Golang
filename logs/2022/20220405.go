package main

// pointer research

import "fmt"

type point struct {
	x int
	y int
	z *string
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

type Vertex struct {
	X, Y float64
}

func PrintPointer(v *Vertex) {
	fmt.Println(v)
}

func PrintValue(v *Vertex) {
	fmt.Println(*v)
}

func main() {
	a := Vertex{3, 4} // not allocated
	PrintValue(&a)    // {3 4}

	b := &Vertex{3, 4} // not allocated
	PrintValue(b)      // {3 4}

	c := Vertex{3, 4} // allocated
	PrintPointer(&c)  // &{3 4}

	d := &Vertex{3, 4} // allocated
	PrintPointer(d)    // &{3 4}
}
