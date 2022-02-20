package main

import "fmt"

type Student struct {
	name   string
	grades []int
	age    int
}

// struct method
// important to put pointer there When needing to change object's value
func (s *Student) setAge(age int) {
	s.age = age
}

func (s Student) getAverageGrade() float64 {
	sum := 0
	for _, v := range s.grades {
		sum += v
	}
	return float64(sum) / float64(len(s.grades))
}

func main() {
	s1 := Student{"Tim", []int{80, 65, 75, 55}, 18}
	fmt.Println(s1.age) // 18
	s1.setAge(7)
	fmt.Println(s1.age) // 7 (without pointer in setAge it will get 18)

}
