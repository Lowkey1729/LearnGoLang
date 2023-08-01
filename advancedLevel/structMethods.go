package main

import "fmt"

type Student struct {
	name   string
	grades []int
	age    int
}

func main() {
	s1 := Student{
		"Grace",
		[]int{12, 90, 9, 12, 13},
		90,
	}

	//s1.setAge(9)
	average := s1.averageGrade()
	fmt.Println(average)
}

//func (student *Student) setAge(age int) {
//	student.age = age
//}

func (student Student) averageGrade() float32 {
	sum := 0
	for _, v := range student.grades {
		sum += v
	}

	return float32(sum) / float32(len(student.grades))
}
