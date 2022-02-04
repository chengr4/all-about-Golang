package main

// fmt: for formatting and printing
import "fmt"

type score struct {
	x int32
	y int64
}

type cirlce struct {
	redius float32
	len    score
}

func changeXWithPointer(pt *score) {

}

// main(): the entry point of the application
func main() {
	p1 := &score{x: 3}
	fmt.Println(p1)
	changeXWithPointer(p1)

}
